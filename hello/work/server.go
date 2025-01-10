package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	Onlinemap map[string]*User
	maplock   sync.RWMutex
	Message   chan string
}

func Newserver(ip string, port int) *Server {
	Server := &Server{
		Ip:        ip,
		Port:      port,
		Onlinemap: make(map[string]*User),
		Message:   make(chan string),
	}
	return Server
}

func (this *Server) listernmassage() {
	for {
		msg := <-this.Message
		this.maplock.Lock()
		for _, client := range this.Onlinemap {
			client.C <- msg
		}
		this.maplock.Unlock()
	}
}

func (this *Server) Btoadcast(user *User, msg string) {
	sendmsg := "[" + user.Age + "]" + user.Name + ":" + msg
	this.Message <- sendmsg
}

func (this *Server) Handler(conn net.Conn) {
	//fmt.Fprintf(conn, "hello world")
	user := NewUser(conn)
	this.maplock.Lock()
	this.Onlinemap[user.Name] = user
	this.maplock.Unlock()
	this.Btoadcast(user, "shangxian")

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("err is ", err)
				return
			}
			if n == 0 {
				fmt.Println("xiaxian")
				return
			}
			msg := string(buf[:n-1])

			this.Btoadcast(user, msg)
		}
	}()

	select {}
}

func (this *Server) Start() {
	lister, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer lister.Close()
	}
	go this.listernmassage()
	for {
		conn, err := lister.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			go this.Handler(conn)
		}

	}
}
