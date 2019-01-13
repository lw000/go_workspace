package lwsockt

import (
	log "github.com/thinkboy/log4go"
	"net"
	"time"
)

func RunClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	defer func() {
		err := conn.Close()
		if err != nil {

		}
	}()

	if err != nil {
		log.Error("fatal error: %s", err.Error())
		return
	}

	buf := make([]byte, 1024)
	for {
		var n int
		n, err = conn.Write([]byte("Hello!"))
		if err != nil {
			log.Error("connected closed")
			break
		}

		if n > 0 {

		}

		n, err = conn.Read(buf)
		if err != nil {
			log.Error("connected closed")
			break
		}

		if n > 0 {

		}

		log.Info("read: %s", string(buf[0:n]))

		time.Sleep(time.Millisecond * time.Duration(1000))
	}
}
