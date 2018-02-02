package main

import (
	"fmt"
)

func Test(ch chan int, index int) {
	ch <- index
	fmt.Println("Go...")
}

func main () {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i ++ {
		chs[i] = make(chan int)
		go Test(chs[i], i)
	}
	for _, ch := range chs {
		value := <- ch
		fmt.Println(value)
	}
}

