package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func closeWrite(conn net.Conn) error {

	type writeCloser interface {
		CloseWrite() error
	}

	if cw, ok := conn.(writeCloser); ok {
		return cw.CloseWrite()
	}
	fmt.Println("nope")
	return conn.Close()
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	closeWrite(conn)
	<-done
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
