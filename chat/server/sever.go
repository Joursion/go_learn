package main

import (
	"fmt"
	"net"
	"os"
)

var (
	host = "localhost:8080"
	data = make([]byte, 1024)
)

// const clients []net.Conn
var clients []net.Conn

func main() {
	fmt.Println("server started", host)
	lt, err := net.Listen("tcp", host)
	defer lt.Close()
	if err != nil {
		fmt.Printf("Error when listen %s, Err : %s\n", host, err)
	}
	for {
		var res string
		conn, err := lt.Accept()
		if err != nil {
			fmt.Println("Error accepting client: ", err.Error())
			os.Exit(0)
		}
		clients = append(clients, conn)

		go func(conn net.Conn) {
			fmt.Println("new connection: ", conn.RemoteAddr())
			length, err := conn.Read(data)
			if err != nil {
				fmt.Printf("Client %v quit. \n", conn.RemoteAddr())
				conn.Close()
				disconnect(conn, conn.RemoteAddr().String())
				return
			}
			name := string(data[:length])
			comeStr := name + "entered the room. "
			notify(conn, comeStr)

			//recieve message from client
			for {
				length, err := conn.Read(data)
				if err != nil {
					fmt.Printf("Client %s quit.\n", name)
					conn.Close()
					disconnect(conn, name)
					return
				}
				res = string(data[:length])
				sprdMsg := name + " said: " + res
				fmt.Println(sprdMsg)
				res = "You said:" + res
				conn.Write([]byte(res))
				notify(conn, sprdMsg)
			}
		}(conn)
	}
}

func disconnect(conn net.Conn, name string) {
	for index, con := range clients {
		if con.RemoteAddr() == conn.RemoteAddr() {
			disMsg := name + " has left the room ."
			fmt.Println(disMsg)
			clients = append(clients[:index], clients[index+1:]...)
			notify(conn, disMsg)
		}
	}
}

func notify(conn net.Conn, msg string) {
	for _, con := range clients {
		if con.RemoteAddr() != conn.RemoteAddr() {
			con.Write([]byte(msg))
		}
	}
}
