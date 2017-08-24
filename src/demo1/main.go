// demo0001 project main.go

package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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
	resp, err := http.Get("http://www.baidu.com")
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
	defer time.Sleep(time.Second)

	{
		m := md5.New()
		n, err := m.Write([]byte("111111111"))
		fmt.Println(fmt.Sprint("n = %d err = %d", n, err))

		fmt.Println(hex.EncodeToString(m.Sum(nil)))
	}

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
	log.Debug("OK")

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

	var body = HttpGet("http://www.baidu.com")
	log.Debug(body)

}
