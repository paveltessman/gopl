package main

import (
	"fmt"
	"time"
)

func count(out chan<- int) {
	for x := range 101 {
		out <- x
		time.Sleep(100 * time.Millisecond)
	}
	close(out)
}

func square(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func print(in <-chan int) {
	for d := range in {
		fmt.Printf("\r%d", d)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go count(naturals)
	go square(naturals, squares)
	print(squares)
	fmt.Printf("\ndone\n")
}
