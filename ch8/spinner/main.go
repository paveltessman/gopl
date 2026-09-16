package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const usage = "Usage: ./spinner <int>"

func fatal() {
	fmt.Println(usage)
	os.Exit(1)
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}

func spin(delay time.Duration) {
	for {
		for _, c := range `-\|/` {
			fmt.Printf("\r%c", c)
			time.Sleep(delay)
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fatal()
	}

	x, err := strconv.Atoi(args[0])
	if err != nil {
		fatal()
	}

	go spin(100 * time.Millisecond)
	fmt.Printf("\r%d\n", fib(x))
}
