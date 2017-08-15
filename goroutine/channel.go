package main

import (
	"fmt"
)

// use channel for goroutine

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
}
