package main

import (
	"Gate/packet"
	"Gate/protocol"
	"Gate/sock"
	"fmt"
	"github.com/golang/protobuf/proto"
	log "github.com/thinkboy/log4go"
	"os"
	"os/signal"
	"time"
)

func PacketTest() {
	req := example.ReqRegisterService{SrvId: 1000, SrvType: 0}
	log.Info("req:[%d] %v", time.Now().Unix(), req)
	reqBuf, err := proto.Marshal(&req)
	if err != nil {

	}

	pkReq := lwpacket.NewPacket(1111, 2222)
	err = pkReq.Encode(reqBuf)
	if err != nil {

	}

	pkAck := lwpacket.NewPacketWithData(pkReq.Bytes())
	data, err := pkAck.Decode()
	if data != nil {

	}

	var req1 = example.ReqRegisterService{}
	err = proto.Unmarshal(pkAck.Bytes(), &req1)
	if err != nil {

	}

	log.Info("ack:[%d] %v", time.Now().Unix(), req1)

	//ack := example.AckRegisterService{Code:100, Data:"1111111111"}
	//log4go.Debug(ack)
	//ack_buf, err := proto.Marshal(&ack)
	//if err != nil {
	//
	//}
}

func installSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill, os.Interrupt)
	go func() {
		select {
		case data := <-c:
			fmt.Println(data)
			lwsockt.Stop()
			os.Exit(1)
		}
	}()
}

func main() {
	installSignal()
	PacketTest()
	go lwsockt.RunClient()
	lwsockt.RunServer()
}
