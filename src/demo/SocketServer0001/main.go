// SocketServer0001 project main.go
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
	//	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleConnect(conn net.Conn) {
	go func() {
		defer func() {
			err := conn.Close() // we're finished with this client
			if err != nil {

			}
		}()

		var (
			n   int
			err error
			buf []byte
		)

		buf = make([]byte, 1024)

		for {
			n, err = conn.Read(buf)
			if err != nil {
				log.Println("connected closed")
				break
			}

			if n > 0 {
				log.Printf("[%s] read:%s\n", conn.RemoteAddr().String(), buf[0:n])
			}

			daytime := time.Now().Format("2006-01-02 15:04:05")
			n, err = conn.Write([]byte(daytime)) // don't care about return value
			if err != nil {

			}

			if n > 0 {

			}
		}
	}()
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":7777")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		log.Println(conn.RemoteAddr().String())

		go handleConnect(conn)
	}
}
