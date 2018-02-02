package channel

import (
	"fmt"
)

func Recv(ch <- chan int, lock chan <- bool) {
	for value := range ch {
		fmt.Println(value)
	}
	lock <- true
}

func Send(ch chan <- int) {
	for i:=0; i < 5; i++{
		ch <- i
	}
	close(ch)
}

func main () {
	ch := make(chan int)
	lock := make(chan bool)
	go Recv(ch, lock)
	go Send(ch)
	<- lock
}

func t() {
	ch := make(chan int)
	chRead := <- chan int(ch)
	chWrite := chan <- int(ch)
}