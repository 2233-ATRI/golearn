package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Price int    `json:"rmb"`
	Actor string `json:"actor"`
}

func main() {
	movie := Movie{"aca", 1991, 134, "wqh"}
	fmt.Println(movie)
	jsonstr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("error is ", err)
	} else {
		fmt.Println(string(jsonstr))
		//fmt.Printf("%s\n", jsonstr)//等价
	}
	my_movie := Movie{}
	err = json.Unmarshal(jsonstr, &my_movie)
	if err != nil {
		fmt.Println("error is ", err)

	} else {
		fmt.Printf("%v\n", my_movie)
	}
}
