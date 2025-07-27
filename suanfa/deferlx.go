package main

import "fmt"

func main() {
	var wahtever [5]struct{}
	for i := range wahtever {
		defer fmt.Println(i)
	}
	for i := range wahtever {
		defer func() { fmt.Println(i) }()
	}
	return
}
