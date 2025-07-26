package main

import (
	"encoding/json"
	"fmt"
)

type preacc struct {
	name string
	age  int
}

type student struct {
	ID     int `json:"id"`
	Gender string
	Age    int
}

func newpre(name string, age int) *preacc {
	return &preacc{name: name, age: age}
}

func (p *preacc) setage(age *int) {
	p.age = *age
}

func main() {
	p1 := newpre("sss", 23)
	a := 30
	p1.setage(&a)
	println(p1.age)

	s1 := student{ID: 1, Gender: "male", Age: 23}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Printf("json str:%s\n", data)
}
