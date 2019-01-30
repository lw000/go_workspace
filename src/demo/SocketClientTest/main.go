// SocketClientTest project main.go
package main

import (
	"log"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Println("fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	defer conn.Close()

	checkError(err)

	buf := make([]byte, 1024)
	for {
		var n int
		n, err = conn.Write([]byte("Hello!"))
		if err != nil {
			log.Println("connected closed")
			break
		}

		if n > 0 {

		}

		n, err = conn.Read(buf)
		if err != nil {
			log.Println("connected closed")
			break
		}

		if n > 0 {

		}

		log.Printf("read:%s\n", string(buf[0:n]))

		time.Sleep(time.Millisecond * time.Duration(1000))
	}
}
