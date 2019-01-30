// demo0001 project main.go

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	//	"pb/game"
	//	"pb/platform"
	"demo_goprotobuf/pb/game"
	"demo_goprotobuf/pb/platform"
	"time"

	proto "github.com/golang/protobuf/proto"
	log "github.com/thinkboy/log4go"
)

type Person struct {
	Name        string `json:"username"`
	Age         int
	Gender      bool `json:",omitempty"`
	Profile     string
	OmitContent string `json:"-"`
	Count       int    `json:",string"`
}

func HttpGet(url string) (body string) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	body = string(res)
	return
}

func main() {

	log.Debug("start")

	{
		married := flag.Bool("married", false, "Are you married?")
		age := flag.Int("age", 22, "How old are you?")
		name := flag.String("name", "", "What your name?")

		var address string
		//flag.StringVar这样的函数第一个参数换成了变量地址，后面的参数和flag.String是一样的。
		flag.StringVar(&address, "address", "GuangZhou", "Where is your address?")

		flag.Parse() //解析输入的参数

		fmt.Println("married :", *married) //不加*号的话,输出的是内存地址
		fmt.Println("age :", *age)
		fmt.Println("name :", *name)
		fmt.Println("address:", address)

	}

	{
		var p *Person = &Person{
			Name:        "brainwu",
			Age:         21,
			Gender:      true,
			Profile:     "I am Wujunbin",
			OmitContent: "OmitConent",
		}

		if bs, err := json.Marshal(&p); err != nil {
			panic(err)
		} else {
			//result --> {"username":"brainwu","Age":21,"Gender":true,"Profile":"I am Wujunbin","Count":"0"}
			fmt.Println(string(bs))
		}
	}

	// {
	// 	var body = HttpGet("http://www.baidu.com")
	// 	log.Debug(body)
	// }

	{
		defer time.Sleep(time.Second)

		login := &platform.CsMsgLogin{
			Device:   1,
			Username: "liwei",
			Userpsd:  "123456",
		}

		msg := &game.NetMsg{}
		log.Debug("%v", msg)

		data, err := proto.Marshal(login)
		if err != nil {
			log.Error("marshaling error: ", err)
			return
		}

		newlogin := &platform.CsMsgLogin{}
		err = proto.Unmarshal(data, newlogin)
		if err != nil {
			log.Error("unmarshaling error: ", err)
		}

		log.Debug("[%v]", login)
		log.Debug("[%v]", newlogin)

	}

}
