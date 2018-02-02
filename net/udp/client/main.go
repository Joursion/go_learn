package main

import(
	"net"
	"os"
	"fmt"
)


func main() {
	if len(os.Args) != 2{
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
	}
	servic := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp", servic)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	_, err = conn.Write([]byte("Hello Server!"))
	checkError(err)
	var buf [512]byte
	n, addr, err := conn.ReadFromUDP(buf[0:])
	checkError(err)
	fmt.Println("Reply from server", addr.String(), string(buf[0:n]))
	conn.Close()
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}