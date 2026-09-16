package main

import (
	"io"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	for  {
	_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
	if err != nil {
		return
	}
	time.Sleep(1 * time.Second)
}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}
		handleConn(conn)
	}
}
