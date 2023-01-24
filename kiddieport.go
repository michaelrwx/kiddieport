package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func scanPort(protocol, hostname string, port int) bool { //function that actually scans the ports
	address := hostname + ":" + strconv.Itoa((port))
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil { //error handler
		return false
	}
	defer conn.Close()
	return true
}

func main() {
	fmt.Println("Which port would you like to scan?")

	var portnum int
	fmt.Scanln(&portnum) //input

	open := scanPort("tcp", "localhost", portnum)
	fmt.Printf("Port Open: %t\n", open)
}
