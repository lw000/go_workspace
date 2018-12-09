// demo000 project main.go
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	log "github.com/jeanphorn/log4go"
)

type shared_data struct {
	Data  int
	Value int
	Name  string
	What  string
	Cmd   string
}

type StringResources struct {
	XMLName        xml.Name         `xml:"resources"`
	ResourceString []ResourceString `xml:"string"`
}

type ResourceString struct {
	XMLName    xml.Name `xml:"string"`
	StringName string   `xml:name,atr`
	InnerText  string   `xml:",innerxml"`
}

var msgqueue = make(chan shared_data, 10)

func main() {
	log.Debug("ok")

	{
		context, err := ioutil.ReadFile("./xml.xml")
		if err != nil {
			log.Error(err)
		}
		var restut StringResources
		err = xml.Unmarshal(context, &restut)
		if err != nil {
			log.Error(err)
		}
		log.Debug(restut)
		log.Debug(restut.ResourceString)
		for _, o := range restut.ResourceString {
			log.Debug(o.StringName + "===" + o.InnerText)
		}
	}

	{
		m := map[string][]string{
			"name":    {"levi"},
			"level":   {"debug"},
			"message": {"1", "2", "3", "4"},
		}

		var data []byte
		var err error
		if data, err = json.Marshal(m); err == nil {
			log.Debug(string(data))
		}
		var m1 map[string][]string
		if json.Unmarshal(data, &m1); err == nil {
			log.Debug(m1)
		}
	}

	{
		a := []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

		for i := 0; i < len(a); i++ {
			fmt.Printf("a[%d]:%d\n", i, a[i])
		}

		for i, v := range a {
			fmt.Printf("a[%d]:%d\n", i, v)
		}
	}

	{
		a := [3][3]int32{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a[i]); j++ {
				fmt.Printf("a[%d][%d]:%d\n", i, j, a[i][j])
			}
		}
	}

	{
		//		a := make([]int, 10)
		var a [10]int
		for i := 0; i < len(a); i++ {
			a[i] = 100 + i
			fmt.Printf("a[%d]:%d\n", i, a[i])
		}
	}
	{
		v, err := strconv.ParseInt("1111", 10, 64)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Println(v)
	}

	go func(s string) {
		for {
			d := <-msgqueue
			fmt.Println("[receiver] Received a sync signal and wait a second...")
			fmt.Printf("d:%v\n", d)
		}
	}("1")

	go func(s string) {
		for {
			d := <-msgqueue
			fmt.Println("[receiver] Received a sync signal and wait a second...")
			fmt.Printf("d:%v\n", d)
		}
	}("2")

	go func() {
		a := []shared_data{
			{100, 1, "levi", "msg", "ok"},
			{100, 2, "levi", "msg", "ok"},
			{100, 3, "levi", "msg", "ok"},
			{100, 4, "levi", "msg", "ok"},
			{100, 5, "levi", "msg", "ok"},
		}
		for _, v := range a {
			msgqueue <- v
		}
	}()

	time.Sleep(time.Second * 1)
}
