package main

import (
	"fmt"
	"sync"
)

func main() {

}

// zhan
type ARRSTACK struct {
	Array []string
	size  int
	lock  sync.Mutex
}

func (stack *ARRSTACK) pash(value string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.Array = append(stack.Array, value)
	stack.size += 1
}

func (stack *ARRSTACK) unp() {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.size < 1 {
		return
	}
	value := stack.Array[stack.size-1]
	fmt.Println(value)
	stack.size -= 1
}

func (stack *ARRSTACK) peek() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.size == 0 {
		return "error"
	}
	v := stack.Array[stack.size-1]
	return v
}

func (stack *ARRSTACK) sizeof() int {
	return stack.size
}
