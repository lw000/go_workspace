// demo0001 project main.go

package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	//"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
	"time"
	//	"github.com/voxelbrain/goptions"
)

type Person struct {
	Name        string `json:"username"`
	Age         int
	Gender      bool `json:",omitempty"`
	Profile     string
	OmitContent string `json:"-"`
	Count       int    `json:",string"`
}

func (p *Person) Marshal() error {
	if str, err := json.Marshal(p); err == nil {
		log.Println(string(str))
	} else {
		log.Println(err)
	}
	return nil
}

func HttpGet(url string) (data []byte, err error) {
	var resp *http.Response
	resp, err = http.Get("http://www.baidu.com")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res []byte
	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data = res
	log.Println(string(data))

	return
}

func md5Test() {
	m := md5.New()
	n, err := m.Write([]byte("111111111"))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(n, hex.EncodeToString(m.Sum(nil)))
}

func main() {
	defer func() {
		time.Sleep(time.Second)
	}()

	//	{
	//		options := struct {
	//			Server   string        `goptions:"-s, --server, obligatory, description='Server to connect to'"`
	//			Password string        `goptions:"-p, --password, description='Don\\'t prompt for password'"`
	//			Timeout  time.Duration `goptions:"-t, --timeout, description='Connection timeout in seconds'"`
	//			Help     goptions.Help `goptions:"-h, --help, description='Show this help'"`

	//			goptions.Verbs
	//			Execute struct {
	//				Command string   `goptions:"--command, mutexgroup='input', description='Command to exectute', obligatory"`
	//				Script  *os.File `goptions:"--script, mutexgroup='input', description='Script to exectute', rdonly"`
	//			} `goptions:"execute"`
	//			Delete struct {
	//				Path  string `goptions:"-n, --name, obligatory, description='Name of the entity to be deleted'"`
	//				Force bool   `goptions:"-f, --force, description='Force removal'"`
	//			} `goptions:"delete"`
	//		}{ // Default values goes here
	//			Timeout: 10 * time.Second,
	//		}
	//		goptions.ParseAndFail(&options)
	//	}

	md5Test()

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
		p := &Person{
			Name:        "brainwu",
			Age:         21,
			Gender:      true,
			Profile:     "I am Wujunbin",
			OmitContent: "OmitConent",
		}
		p.Marshal()
	}

	{
		var (
			data []byte
			err  error
		)
		data, err = HttpGet("http://www.baidu.com")
		if err != nil {

		}
		log.Println("%s", string(data))
	}
}
