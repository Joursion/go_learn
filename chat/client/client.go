//reference http://blog.csdn.net/wangshubo1989/article/details/70668916
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var writeStr, readStr = make([]byte, 1024), make([]byte, 1024)

var (
	host   = "localhost:8080"
	data   = make([]byte, 1024)
	reader = bufio.NewReader(os.Stdin)
)

func main() {
	conn, err := net.Dial("tcp", host)
	defer conn.Close()
	if err != nil {
		fmt.Println("server not found.")
		os.Exit(-1)
	}
	fmt.Println("Connection success ...")
	fmt.Printf("Enter your name:  ")
	fmt.Scanf("%s", &writeStr)
	in, err := conn.Write([]byte(writeStr))
	if err != nil {
		fmt.Printf("Error when send to server: %d\n", in)
		os.Exit(0)
	}
	go read(conn)

	for {
		writeStr, _, _ := reader.ReadLine()
		if string(writeStr) == "quit" {
			fmt.Println("Communication rerminated.")
			os.Exit(1)
		}

		in, err := conn.Write([]byte(writeStr))
		if err != nil {
			fmt.Printf("Error when send to server : %d \n", in)
			os.Exit(0)
		}
	}
}

func read(conn net.Conn) {
	for {
		length, err := conn.Read(readStr)
		if err != nil {
			fmt.Printf("Error when read from server. Error: %s\n", err)
			os.Exit(0)
		}
		fmt.Println(string(readStr[:length]))
	}
}
