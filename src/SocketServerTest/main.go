// SocketServerTest project main.go
package main

import (
	"fmt"
	"net"
	"os"
	"time"

	log "github.com/thinkboy/log4go"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 100)

	for {
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("conn closed")
			return
		}

		fmt.Println("recv msg:", string(buf[0:n]))

		conn.Write([]byte("world!"))

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	listen_sock, err := net.Listen("tcp", ":12563")
	defer listen_sock.Close()

	checkError(err)

	log.Debug(time.Now())
	log.Debug("server start... port:[12563]")

	for {
		conn, err := listen_sock.Accept()
		if err != nil {
			continue
		}

		log.Debug(conn.RemoteAddr().String(), " tcp connect success")

		go handleClient(conn)
	}

}
