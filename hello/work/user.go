package main

import "net"

type User struct {
	Name   string
	Age    string
	C      chan string
	conn   net.Conn
	Server *Server
}

func NewUser(conn net.Conn, server Server) *User {

	useraddr := conn.RemoteAddr().String()
	user := &User{
		Name:   useraddr,
		conn:   conn,
		C:      make(chan string),
		Age:    useraddr,
		Server: &server,
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

func (this *User) Online() {
	this.Server.maplock.Lock()
	this.Server.Onlinemap[this.Name] = this
	this.Server.maplock.Unlock()
	this.Server.Btoadcast(this, "shangxain")
}

func (this *User) Offline() {
	this.Server.maplock.Lock()
	delete(this.Server.Onlinemap, this.Name)
	this.Server.maplock.Unlock()
	this.Server.Btoadcast(this, "xiaxian")
}

func (this *User) Domassage() {
	//this.Server.Btoadcast(User, "msg")
}
