// SocketClientTest project main.go
package main

import (
	"fmt"
	//	"io/ioutil"

	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleRecv(conn net.Conn) {
	buf := make([]byte, 100)

	for {
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("server conn closed")
			return
		}

		fmt.Println("recv msg:", string(buf[0:n]))

		time.Sleep(time.Second * 1)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12563")
	defer conn.Close()

	checkError(err)

	buf := make([]byte, 100)

	for {
		conn.Write([]byte("Hello!"))

		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("conn closed")
			return
		}

		fmt.Println("recv server msg:", string(buf[0:n]))

		time.Sleep(time.Second * time.Duration(1))
	}
}
