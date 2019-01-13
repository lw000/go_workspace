// SocketServerTest project main.go
package main

import (
	"fmt"
	log "github.com/thinkboy/log4go"
	"net"
	"os"
	"time"
)

func handleClient(conn net.Conn) {
	defer func() {
		err := conn.Close()
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
			log.Debug(("connected closed"))
			break
		}
		if n > 0 {

		}

		log.Debug("read:%s", string(buf[0:n]))

		n, err = conn.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
		if err != nil {
			log.Debug(("connected closed"))
			break
		}

		if n > 0 {

		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	s, err := net.Listen("tcp", ":7777")
	checkError(err)

	defer func() {
		err := s.Close()
		if err != nil {

		}
	}()

	log.Debug("server start... port:[12563]")

	for {
		conn, err := s.Accept()
		if err != nil {
			continue
		}

		log.Debug("[%s]", conn.RemoteAddr().String())

		go handleClient(conn)
	}
}
