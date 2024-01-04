package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination string, port string) string {
	// combines destination and port number to form an address
	address := destination + ":" + port
	// a timeout for the duration of a connection attempt.
	timeout := time.Duration(5 * time.Second)
	//attempt a TCP connection with the specified address and a timeout.
	conn, err := net.DialTimeout("tcp", address, timeout)

	var status string //variable to hold status message.
	if err != nil {
		//if an error is occured during connection attempt print a status message.
		status = fmt.Sprintf("[DOWN] %v is unreachable, \n Error : %v", destination, err)
	} else {
		//if no error is occured print the following status message.
		status = fmt.Sprintf("[UP] %v is reachable. \n From: %v\n To : %v", destination, conn.LocalAddr(), conn.RemoteAddr())
		
	}
	return status
}
