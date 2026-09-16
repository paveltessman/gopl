package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := range 101 {
			naturals <- x
			time.Sleep(100 * time.Millisecond)
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for d := range squares {
		fmt.Printf("\r%d", d)
	}
	fmt.Printf("\ndone\n")
}
