// package main
//
// import (
//
//	"fmt"
//	"time"
//
// )
//
// //	func HU() {
// //		time.Sleep(2 * time.Second)
// //		fmt.Println("HU")
// //	}
// //
// //	func main() {
// //		go HU()
// //		fmt.Println("main")
// //		for {
// //			time.Sleep(1000)
// //		}
// //	}
// //
// //	func Hu(ch chan int) {
// //		time.Sleep(1 * time.Second)
// //		fmt.Println("huu")
// //		ch <- 1000
// //	}
// //
// //	func main() {
// //		ch := make(chan int)
// //		go Hu(ch)
// //		fmt.Println("start")
// //		v := <-ch
// //		fmt.Println("end", v)
// //
// // }
//
//	func Receive(ch chan int) {
//		time.Sleep(time.Second)
//		for {
//			select {
//			case v, ok := <-ch:
//				if !ok {
//					fmt.Println("close", v)
//					return
//				}
//				fmt.Println(v)
//			}
//
//		}
//	}
//
//	func Send(ch chan int) {
//		for i := 0; i < 10; i++ {
//			ch <- i
//			fmt.Println("send", i)
//		}
//		close(ch)
//	}
//
//	func main() {
//		ch := make(chan int, 10)
//		go Receive(ch)
//		go Send(ch)
//		for {
//			time.Sleep(time.Second * 100)
//
//		}
//	}
package main

import (
	"fmt"
	"sync"
)

type Money struct {
	lock   sync.Mutex
	amount int
}

func (m *Money) add(amount int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.amount = m.amount + amount
}

func (m *Money) sub(amount int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if amount < m.amount {
		m.amount = m.amount - amount
	} else {
		fmt.Println("meiqian")
		return
	}

}
