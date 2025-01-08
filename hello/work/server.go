package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func Newserver(ip string, port int) *Server {
	Server := &Server{
		Ip:   ip,
		Port: port,
	}
	return Server
}

func (this *Server) Handler(conn net.Conn) {
	fmt.Fprintf(conn, "hello world")
}

func (this *Server) Start() {
	lister, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer lister.Close()
	}

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
