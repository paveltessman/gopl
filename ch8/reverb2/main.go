package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
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

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	for input.Scan() {
		wg.Add(1)
		go echo(c, input.Text(), 1*time.Second, &wg)
	}
	if err := input.Err(); err != nil {
		panic(err)
	}

	go func() {
		wg.Wait()
		closeWrite(c)
	}()
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
		go handleConn(conn)
	}
}
