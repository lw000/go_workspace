// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"runtime"
	"time"

	"fmt"

	"github.com/gorilla/websocket"
)

var host = *flag.String("host", "ws://127.0.0.1:5000/ws", "websocket host")
var rid = *flag.Int("rid", 1, "room id")
var ridnum = *flag.Int("ridnum", 10, "room num")
var begin = *flag.Int("begin", 1, "begin user id")
var end = *flag.Int("end", 1000, "end user id")

type client_args struct {
	rid  int
	uid  int
	data int
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	log.SetFlags(0)

	numPerRoom := (end - begin) / ridnum //每个房间人数

	for i := begin; i < end; i++ {

		extra := "eyJpZCI6ImFhYWEiLCJuYW1lIjoiZ3Vlc3RfYWFhYSIsIm1vdW50Ijp7ImlkIjoxMjM0NTYsIm5hbWUiOiJtb3VudDEyMzQ1NiJ9fQ=="

		u_rid := rid + i/numPerRoom
		u_uid := i

		url := fmt.Sprintf("%s?rid=%d&uid=%d&extra=%s", host, u_rid, u_uid, extra)
		log.Printf("client connect to %s", url)

		args := &client_args{
			rid:  u_rid,
			uid:  u_uid,
			data: u_rid + u_uid,
		}

		con, _, err := websocket.DefaultDialer.Dial(url, nil)
		go func(con *websocket.Conn, c *client_args) {
			for {
				if err != nil {
					log.Fatal("dial:", err)
				}
				defer con.Close()
				for {
					_, msg, err := con.ReadMessage()
					if err != nil {
						log.Println("read:", err)
						return
					}
					str := fmt.Sprintf("[rid:%d uid:%d, data:%s", c.rid, c.uid, string(msg))
					log.Printf(str)
				}
			}
		}(con, args)

		//		go func(con *websocket.Conn, c *client_args) {
		//			for {
		//				for {
		//					str := fmt.Sprintf("[rid:%d uid:%d, data:%d", c.rid, c.uid, c.data)
		//					con.WriteMessage(websocket.TextMessage, []byte(str))
		//					time.Sleep(time.Second * 10)
		//				}
		//			}
		//		}(con, args)

		time.Sleep(time.Millisecond * 10)
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			return
		}
	}
}
