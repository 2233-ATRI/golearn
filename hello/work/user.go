package main

import "net"

type User struct {
	Name string
	Age  string
	C    chan string
	conn net.Conn
}

func NewUser(conn net.Conn) *User {

	useraddr := conn.RemoteAddr().String()
	user := &User{
		Name: useraddr,
		conn: conn,
		C:    make(chan string),
		Age:  useraddr,
	}
	go user.lensten()
	return user
}

func (u *User) lensten() {
	for {
		mag := <-u.C
		u.conn.Write([]byte(mag + "\n"))
	}
}
