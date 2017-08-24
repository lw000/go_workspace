// SocketServer0001 project main.go
package main

import (
	"fmt"
	"net"
	"os"
	//	"time"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("ip4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//		daytime := time.LocalTime().String()
		daytime := "111111111111111111111"
		conn.Write([]byte(daytime)) // don't care about return value
		conn.Close()                // we're finished with this client
	}
}
func checkError(err os.Error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.String())
		os.Exit(1)
	}
}
