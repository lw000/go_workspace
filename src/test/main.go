// demo0001 project main.go

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//	"net/url"

	log "github.com/thinkboy/log4go"
)

func httpGet() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func main() {
	var val1 = 2
	var val2 = 1
	if val1 > val2 {
		httpGet()
	}

	log.Debug("OK")

	//	httpPost()
	//	httpPostForm()
	//	httpDo()
}
