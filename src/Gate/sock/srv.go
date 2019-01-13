package lwsockt

import (
	log "github.com/thinkboy/log4go"
	"net"
	"time"
)

var (
	listener *net.TCPListener
)

func init() {

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
				log.Error("connected closed")
				break
			}

			if n > 0 {
				log.Error("[%s] read:%s\n", conn.RemoteAddr().String(), buf[0:n])
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

func RunServer() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":7777")
	if err != nil {
		log.Error("fatal error: %s", err.Error())
		return
	}

	listener, err = net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Error("fatal error: %s", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		log.Info(conn.RemoteAddr().String())

		go handleConnect(conn)
	}
}

func Stop() {

}
