package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Printf("\r%d ", countdown)
		select {
		case <-tick:
		case <-abort:
			fmt.Printf("\rabort!\n")
			return
		}
	}
	fmt.Printf("\rlaunch!\n")
}
