package main

import (
	"net"
	"os"
	"fmt"
)

func main () {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s service proto name\n", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	proto := os.Args[2]
	name := os.Args[3]

	fmt.Println("---", service, proto, name)
	cname, addrs, err := net.LookupSRV(service, proto, name)
	if err != nil {
		fmt.Println("LookupSRV error:", err.Error())
		os.Exit(1)
	}

	fmt.Println("The service canonical name is: ", cname)
	for _, addr := range addrs {
		fmt.Println("The service addr list is:", addr)
	}
	os.Exit(0)
}