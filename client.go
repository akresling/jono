package jono

import (
	"fmt"
	"net"
)

// StartTCPListener is a basic tcp listener
func StartTCPListener() {
	l, err := net.Listen("tcp", ":3333")
	if err != nil {
		panic(err)
	}

	defer l.Close()
	fmt.Println("Listening on localhost:3333")

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buf := make([]byte, 2048)

	rL, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("SIZE: %d \t RAW : %v\n", rL, buf[:rL])
	/*
		var vReq pb.JonoRequest
		err = proto.Unmarshal(buf[:rL], &vReq)
		if err != nil {
			panic(err)
		}
	*/
	conn.Write([]byte("SUCCESS"))
}
