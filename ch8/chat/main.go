package main

import (
	"bufio"
	"fmt"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcast() {
	clients := make(map[client]bool)

	for {

		select {

		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)

		}
	}
}

func writeToClient(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)

	go writeToClient(conn, ch)

	who := conn.RemoteAddr().String()

	ch <- "You are " + who
	messages <- who + " has arrived!"
	entering <- ch

	input := bufio.NewScanner(conn)

	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	if err := input.Err(); err != nil {
		fmt.Println(err)
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func main() {
	const addr = "localhost:8000"
	lestener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	go broadcast()

	fmt.Printf("Listening on %s\n", addr)

	for {
		conn, err := lestener.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go handleConn(conn)
	}
}
