package main

import (
	"fmt"
	"sync"
)

// A WaitGroup must not be copied after first use.

func main() {
	var wg sync.WaitGroup
	var arr = []int{1, 2, 5}
	fmt.Println("Start...")
	for _, v := range arr {
		wg.Add(1) //wg.Add(2) error
		go func(num int) {
			defer wg.Done()
			fmt.Println("number : ", num)
		}(v)
	}
	wg.Wait()
	fmt.Println("End...")
}
