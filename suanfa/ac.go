package main

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"net"
	"sync"
	"time"
)

const (
	headerSize = 4 + 8 // keyLen (4 bytes) + dataLen (8 bytes)
)

// Conn 表示一个网络连接
type Conn struct {
	rw      *bufio.ReadWriter
	netConn net.Conn
	lock    sync.Mutex
	recvCh  chan recvItem
	once    sync.Once
	closed  bool
}

// recvItem 表示接收到的数据项
type recvItem struct {
	key string
	buf *bufferedReader
}

// bufferedReader 是一个缓冲读取器
type bufferedReader struct {
	data   []byte
	cursor int
}

func (b *bufferedReader) Read(p []byte) (int, error) {
	if b.cursor >= len(b.data) {
		return 0, io.EOF
	}
	n := copy(p, b.data[b.cursor:])
	b.cursor += n
	return n, nil
}

// NewConn 创建一个新的 Conn 实例
func NewConn(conn net.Conn) *Conn {
	rwc := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	c := &Conn{
		rw:      rwc,
		netConn: conn,
		recvCh:  make(chan recvItem, 1024),
	}
	go c.readLoop()
	return c
}

// Send 发送数据并返回一个写入器
func (c *Conn) Send(key string) (io.WriteCloser, error) {
	pr, pw := io.Pipe()
	errCh := make(chan error, 1)

	go func() {
		defer pr.Close()
		data, err := io.ReadAll(pr)
		if err != nil {
			errCh <- fmt.Errorf("read pipe: %w", err)
			return
		}

		c.lock.Lock()
		if c.closed {
			c.lock.Unlock()
			errCh <- fmt.Errorf("connection closed")
			return
		}

		keyLen := uint32(len(key))
		if err := binary.Write(c.rw, binary.BigEndian, keyLen); err != nil {
			c.lock.Unlock()
			errCh <- fmt.Errorf("write key length: %w", err)
			return
		}
		if _, err := c.rw.Write([]byte(key)); err != nil {
			c.lock.Unlock()
			errCh <- fmt.Errorf("write key: %w", err)
			return
		}
		dataLen := uint64(len(data))
		if err := binary.Write(c.rw, binary.BigEndian, dataLen); err != nil {
			c.lock.Unlock()
			errCh <- fmt.Errorf("write data length: %w", err)
			return
		}
		c.lock.Unlock()

		// 分块写入数据
		const chunkSize = 1 << 20 // 1MB
		for i := 0; i < len(data); i += chunkSize {
			end := i + chunkSize
			if end > len(data) {
				end = len(data)
			}
			if _, err := c.rw.Write(data[i:end]); err != nil {
				errCh <- fmt.Errorf("write data: %w", err)
				return
			}
			if err := c.rw.Flush(); err != nil {
				errCh <- fmt.Errorf("flush: %w", err)
				return
			}
		}
		if err := c.rw.Flush(); err != nil {
			errCh <- fmt.Errorf("final flush: %w", err)
			return
		}
		errCh <- nil
		fmt.Printf("send completed, key: %s, dataLen: %d, time: %v\n", key, len(data), time.Now())
	}()

	return &errorPipeWriter{pw: pw, errCh: errCh}, nil
}

// errorPipeWriter 是一个带有错误通道的管道写入器
type errorPipeWriter struct {
	pw    *io.PipeWriter
	errCh chan error
}

func (w *errorPipeWriter) Write(p []byte) (int, error) {
	return w.pw.Write(p)
}

func (w *errorPipeWriter) Close() error {
	if err := w.pw.Close(); err != nil {
		return err
	}
	return <-w.errCh
}

// Receive 接收数据
func (c *Conn) Receive() (string, io.Reader, error) {
	item, ok := <-c.recvCh
	if !ok {
		fmt.Printf("Receive EOF, time: %v\n", time.Now())
		return "", nil, io.EOF
	}
	fmt.Printf("received key: %s, time: %v\n", item.key, time.Now())
	return item.key, item.buf, nil
}

// Close 关闭连接
func (c *Conn) Close() {
	c.once.Do(func() {
		c.lock.Lock()
		c.closed = true
		fmt.Printf("initiating close, time: %v\n", time.Now())
		go func() {
			timeout := time.NewTimer(60 * time.Second) // 延长到 60 秒
			defer timeout.Stop()
			ticker := time.NewTicker(100 * time.Millisecond)
			defer ticker.Stop()
			for {
				select {
				case <-timeout.C:
					fmt.Printf("close timeout, recvCh len: %d, time: %v\n", len(c.recvCh), time.Now())
					close(c.recvCh)
					return
				case <-ticker.C:
					if len(c.recvCh) == 0 {
						close(c.recvCh)
						fmt.Printf("recvCh closed, len: %d, time: %v\n", len(c.recvCh), time.Now())
						return
					}
					fmt.Printf("waiting to close, recvCh len: %d, time: %v\n", len(c.recvCh), time.Now())
				}
			}
		}()
		c.netConn.Close()
		c.lock.Unlock()
		fmt.Printf("conn closed, time: %v\n", time.Now())
	})
}

// readLoop 读取循环
func (c *Conn) readLoop() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("readLoop recovered from panic: %v, time: %v\n", r, time.Now())
		}
	}()

	for {
		c.lock.Lock()
		if c.closed {
			c.lock.Unlock()
			return
		}
		c.lock.Unlock()

		var keyLen uint32
		if err := binary.Read(c.rw, binary.BigEndian, &keyLen); err != nil {
			if err == io.EOF || err == io.ErrClosedPipe {
				fmt.Printf("readLoop EOF or closed pipe, recvCh len: %d, time: %v\n", len(c.recvCh), time.Now())
				time.Sleep(10 * time.Second) // 延迟 10 秒等待 recvCh 清空
				return
			}
			if netErr, ok := err.(net.Error); ok && !netErr.Temporary() {
				fmt.Printf("readLoop non-temporary net error: %v, recvCh len: %d, time: %v\n", err, len(c.recvCh), time.Now())
				time.Sleep(10 * time.Second) // 延迟 10 秒等待 recvCh 清空
				return
			}
			fmt.Printf("temporary read key length error, retrying: %v, time: %v\n", err, time.Now())
			time.Sleep(10 * time.Millisecond)
			continue
		}

		keyBytes := make([]byte, keyLen)
		if _, err := io.ReadFull(c.rw, keyBytes); err != nil {
			if err == io.EOF || err == io.ErrClosedPipe {
				fmt.Printf("readLoop EOF or closed pipe, recvCh len: %d, time: %v\n", len(c.recvCh), time.Now())
				time.Sleep(10 * time.Second) // 延迟 10 秒等待 recvCh 清空
				return
			}
			if netErr, ok := err.(net.Error); ok && !netErr.Temporary() {
				fmt.Printf("readLoop non-temporary net error: %v, recvCh len: %d, time: %v\n", err, len(c.recvCh), time.Now())
				time.Sleep(10 * time.Second) // 延迟 10 秒等待 recvCh 清空
				return
			}
			fmt.Printf("temporary read key error, retrying: %v, time: %v\n", err, time.Now())
			time.Sleep(10 * time.Millisecond)
			continue
		}
		key := string(keyBytes)

		var dataLen uint64
		if err := binary.Read(c.rw, binary.BigEndian, &dataLen); err != nil {
			if err == io.EOF || err == io.ErrClosedPipe {
				fmt.Printf("readLoop EOF or closed pipe, recvCh len: %d, time: %v\n", len(c.recvCh), time.Now())
				time.Sleep(10 * time.Second) // 延迟 10 秒等待 recvCh 清空
				return
			}
			if netErr, ok := err.(net.Error); ok && !netErr.Temporary() {
				fmt.Printf("readLoop non-temporary net error: %v, recvCh len: %d, time: %v\n", err, len(c.recvCh), time.Now())
				time.Sleep(10 * time.Second) // 延迟 10 秒等待 recvCh 清空
				return
			}
			fmt.Printf("temporary read data length error, retrying: %v, time: %v\n", err, time.Now())
			time.Sleep(10 * time.Millisecond)
			continue
		}

		data := make([]byte, dataLen)
		if _, err := io.ReadFull(c.rw, data); err != nil {
			if err == io.EOF || err == io.ErrClosedPipe {
				fmt.Printf("readLoop EOF or closed pipe, recvCh len: %d, time: %v\n", len(c.recvCh), time.Now())
				time.Sleep(10 * time.Second) // 延迟 10 秒等待 recvCh 清空
				return
			}
			if netErr, ok := err.(net.Error); ok && !netErr.Temporary() {
				fmt.Printf("readLoop non-temporary net error: %v, recvCh len: %d, time: %v\n", err, len(c.recvCh), time.Now())
				time.Sleep(10 * time.Second) // 延迟 10 秒等待 recvCh 清空
				return
			}
			fmt.Printf("temporary read data error, retrying: %v, time: %v\n", err, time.Now())
			time.Sleep(10 * time.Millisecond)
			continue
		}

		select {
		case c.recvCh <- recvItem{key: key, buf: &bufferedReader{data: data}}:
			fmt.Printf("sent to recvCh, key: %s, dataLen: %d, time: %v\n", key, len(data), time.Now())
		default:
			fmt.Printf("Warning: recvCh full, dropping data for key %s, time: %v\n", key, time.Now())
		}
	}
}

// 测试代码保持不变...

// 测试代码保持不变...

// 测试代码保持不变...

// 测试代码保持不变...

// 测试代码保持不变...

// 测试代码保持不变，以下省略...

// 除了上面规定的接口，你还可以自行定义新的类型，变量和函数以满足实现需求

//////////////////////////////////////////////
///////// 接下来的代码为测试代码，请勿修改 /////////
//////////////////////////////////////////////

// 连接到测试服务器，获得一个你实现的连接对象
func dial(serverAddr string) *Conn {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		panic(err)
	}
	return NewConn(conn)
}

// 启动测试服务器
func startServer(handle func(*Conn)) net.Listener {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("[WARNING] ln.Accept", err)
				return
			}
			go handle(NewConn(conn))
		}
	}()
	return ln
}

// 简单断言
func assertEqual[T comparable](actual T, expected T) {
	if actual != expected {
		panic(fmt.Sprintf("actual:%v expected:%v\n", actual, expected))
	}
}

// 简单 case：单连接，双向传输少量数据
func testCase0() {
	const (
		key  = "Bible"
		data = `Then I heard the voice of the Lord saying, “Whom shall I send? And who will go for us?”
And I said, “Here am I. Send me!”
Isaiah 6:8`
	)
	ln := startServer(func(conn *Conn) {
		// 服务端等待客户端进行传输
		_key, reader, err := conn.Receive()
		if err != nil {
			panic(err)
		}
		assertEqual(_key, key)
		dataB, err := io.ReadAll(reader)
		if err != nil {
			panic(err)
		}
		assertEqual(string(dataB), data)

		// 服务端向客户端进行传输
		writer, err := conn.Send(key)
		if err != nil {
			panic(err)
		}
		n, err := writer.Write([]byte(data))
		if err != nil {
			panic(err)
		}
		if n != len(data) {
			panic(n)
		}
		conn.Close()
	})
	//goland:noinspection GoUnhandledErrorResult
	defer ln.Close()

	conn := dial(ln.Addr().String())
	// 客户端向服务端传输
	writer, err := conn.Send(key)
	if err != nil {
		panic(err)
	}
	n, err := writer.Write([]byte(data))
	if n != len(data) {
		panic(n)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}
	// 客户端等待服务端传输
	_key, reader, err := conn.Receive()
	if err != nil {
		panic(err)
	}
	assertEqual(_key, key)
	dataB, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	assertEqual(string(dataB), data)
	conn.Close()
}

// 生成一个随机 key
func newRandomKey() string {
	buf := make([]byte, 8)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(buf)
}

// 读取随机数据，并返回随机数据的校验和：用于验证数据是否完整传输
func readRandomData(reader io.Reader, hash hash.Hash) (checksum string) {
	hash.Reset()
	var buf = make([]byte, 23<<20) //调用者读取时的 buf 大小不是固定的，你的实现中不可假定 buf 为固定值
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		_, err = hash.Write(buf[:n])
		if err != nil {
			panic(err)
		}
	}
	checksum = hex.EncodeToString(hash.Sum(nil))
	return checksum
}

// 写入随机数据，并返回随机数据的校验和：用于验证数据是否完整传输
func writeRandomData(writer io.Writer, hash hash.Hash) (checksum string) {
	hash.Reset()
	const (
		dataSize = 500 << 20 //一个 key 对应 500MB 随机二进制数据，dataSize 也可以是其他值，你的实现中不可假定 dataSize 为固定值
		bufSize  = 1 << 20   //调用者写入时的 buf 大小不是固定的，你的实现中不可假定 buf 为固定值
	)
	var (
		buf  = make([]byte, bufSize)
		size = 0
	)
	for i := 0; i < dataSize/bufSize; i++ {
		_, err := rand.Read(buf)
		if err != nil {
			panic(err)
		}
		_, err = hash.Write(buf)
		if err != nil {
			panic(err)
		}
		n, err := writer.Write(buf)
		if err != nil {
			panic(err)
		}
		size += n
	}
	if size != dataSize {
		panic(size)
	}
	checksum = hex.EncodeToString(hash.Sum(nil))
	return checksum
}

// 复杂 case：多连接，双向传输，大量数据，多个不同的 key
func testCase1() {
	var (
		mapKeyToChecksum = map[string]string{}
		lock             sync.Mutex
	)
	ln := startServer(func(conn *Conn) {
		// 服务端等待客户端进行传输
		key, reader, err := conn.Receive()
		if err != nil {
			panic(err)
		}
		var (
			h         = sha256.New()
			_checksum = readRandomData(reader, h)
		)
		lock.Lock()
		checksum, keyExist := mapKeyToChecksum[key]
		lock.Unlock()
		if !keyExist {
			panic(fmt.Sprintln(key, "not exist"))
		}
		assertEqual(_checksum, checksum)

		// 服务端向客户端连续进行 2 次传输
		for _, key := range []string{newRandomKey(), newRandomKey()} {
			writer, err := conn.Send(key)
			if err != nil {
				panic(err)
			}
			checksum := writeRandomData(writer, h)
			lock.Lock()
			mapKeyToChecksum[key] = checksum
			lock.Unlock()
			err = writer.Close() //表明该 key 的所有数据已传输完毕
			if err != nil {
				panic(err)
			}
		}
		conn.Close()
	})
	//goland:noinspection GoUnhandledErrorResult
	defer ln.Close()

	conn := dial(ln.Addr().String())
	// 客户端向服务端传输
	var (
		key = newRandomKey()
		h   = sha256.New()
	)
	writer, err := conn.Send(key)
	if err != nil {
		panic(err)
	}
	checksum := writeRandomData(writer, h)
	lock.Lock()
	mapKeyToChecksum[key] = checksum
	lock.Unlock()
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	// 客户端等待服务端的多次传输
	keyCount := 0
	for {
		key, reader, err := conn.Receive()
		if err == io.EOF {
			// 服务端所有的数据均传输完毕，关闭连接
			break
		}
		if err != nil {
			panic(err)
		}
		_checksum := readRandomData(reader, h)
		lock.Lock()
		checksum, keyExist := mapKeyToChecksum[key]
		lock.Unlock()
		if !keyExist {
			panic(fmt.Sprintln(key, "not exist"))
		}
		assertEqual(_checksum, checksum)
		keyCount++
	}
	assertEqual(keyCount, 2)
	conn.Close()
}

func main() {
	testCase0()
	testCase1()
}
