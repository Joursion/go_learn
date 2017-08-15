package main

import (
	"context"
	"fmt"
	"time"
)

//reference: http://www.jianshu.com/p/6032f2db6be5

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("parent Done: ", ctx.Err())
					return
				case dst <- n:
					n++
					go childFunc(ctx, &n)
				}
			}
		}()
		return dst
	}

	/*  WithCancel returns a copy of parent with a
	new Done channel. The returned context's Done channel
	is closed when the returned cancel function is called
	or when the parent context's Done channel is closed,
	whichever happens first.
	*/

	ctx, cancel := context.WithCancel(context.Background())
	for n := range gen(ctx) {
		fmt.Println(n)
		if n >= 5 {
			break
		}
	}
	cancel()
	time.Sleep(5 * time.Second)
}

func childFunc(cont context.Context, num *int) {
	ctx, _ := context.WithCancel(cont) // init a child context
	for {
		select {
		case <-ctx.Done():
			fmt.Println("child one: ", ctx.Err())
			return
		}
	}
}
