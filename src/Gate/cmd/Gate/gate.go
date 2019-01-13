package main

import (
	"Gate/auth"
	"Gate/packet"
	"Gate/protocol"
	"Gate/sock"
	"Gate/utilty"
	"Gate/ws"
	"fmt"
	"github.com/golang/protobuf/proto"
	log "github.com/thinkboy/log4go"
	"github.com/vmihailenco/msgpack"
	"os"
	"os/signal"
)

var (
	quit chan os.Signal
)

func PacketTest() {
	reqRegisterService := example.ReqRegisterService{SrvId: 1000, SrvType: 0}
	log.Info("req: %v", reqRegisterService)
	reqBuf, err := proto.Marshal(&reqRegisterService)
	if err != nil {
		log.Error(err)
		return
	}

	if reqBuf == nil {
		return
	}

	ackRegisterService := example.AckRegisterService{Code: 100, Data: "1111111111"}
	log.Info("ack: %v", ackRegisterService)
	ack, err := proto.Marshal(&ackRegisterService)
	if err != nil {
		log.Error(err)
		return
	}

	if ack == nil {
		return
	}

	pkReq := lwpacket.NewPacket(1111, 2222)
	err = pkReq.Encode(reqBuf)
	if err != nil {
		log.Error(err)
		return
	}

	pkAck := lwpacket.NewPacketWithData(pkReq.Bytes())
	data, err := pkAck.Decode()
	if data == nil {
		return
	}

	var req1 = example.ReqRegisterService{}
	err = proto.Unmarshal(pkAck.Bytes(), &req1)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("ack: %v", req1)
}

func installSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill, os.Interrupt)
	go func() {
		select {
		case data := <-c:
			fmt.Println(data)
			lwsockt.StopSrv()
			lwsockt.StopCli()
			lwws.StopSrv()
			lwws.StopCli()
			os.Exit(1)
		}
	}()
}

func Test() {
	log.Info(lwutilty.UUID())
	log.Info(string(lwauth.Md5([]byte("111111111111111111111"))))

	in := map[string]string{
		"foo":   "mwerwerewrwr",
		"liwei": "sdffffdsfsf",
	}

	data, err := msgpack.Marshal(in)
	if err != nil {
		return
	}

	if len(data) != 0 {
		log.Info(data)
	}
}

func main() {
	Test()
	installSignal()
	PacketTest()
	go lwsockt.StartCli()
	go lwws.StartCli()
	go lwws.StartSrv()
	go lwsockt.StartSrv()
	quit = make(chan os.Signal, 1)
	select {
	case <-quit:
		break
	}
}
