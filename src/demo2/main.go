// demo0002 project main.go
package main

import (
	"encoding/json"
	"fmt"

	//	"os"

	//	"runtime"
	"time"

	//	"github.com/henrylee2cn/pholcus/exec"
	"github.com/json-iterator/go"
	log "github.com/thinkboy/log4go"
)

func for_test() {
	for {
		fmt.Println("die for")
		break
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	var i = 0
	for i < 5 {
		fmt.Println(i)
		i = i + 1
	}
}

func if_test() {
	if a := 0; a < 5 {
		fmt.Println("a < 100")
	} else {
		fmt.Println("a >= 100")
	}
}

func Factorial(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * Factorial(x-1)
	}

	return
}

func switch_test() {
	x := 1
	switch x {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
	}
}

func sclie_test() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
}

func array_test() {
	var numbers []int
	numbers = make([]int, 5, 5)
	numbers = append(numbers, 1, 2, 3, 4, 5)

	var numbers1 []int
	numbers1 = make([]int, 5, 5)
	numbers1 = []int{10, 20, 30, 40, 50}
	copy(numbers, numbers1)
	fmt.Println(numbers, len(numbers), cap(numbers), numbers1)
}

func json_test() {
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	log.Debug(string(slcB))
}

func interface_test() {

}

func chan_test() {
	//		chm := make(chan int)
	//		chn := make(chan int)
	//		quit := make(chan bool)
	var chm = make(chan int)
	var chn = make(chan int)
	var quit = make(chan bool)
	chm = make(chan int)
	chn = make(chan int)
	quit = make(chan bool)
	go func(m chan int, n chan int) {
		i := 0
		j := 50
		for i < 50 {
			m <- i
			n <- j
			i++
			j--
		}
		quit <- true
	}(chm, chn)

ForEnd:
	for {
		select {
		case m := <-chm:
			{
				log.Debug("m : %d", m)
			}
		case n := <-chn:
			{
				log.Debug("n : %d", n)
			}
		case <-quit:
			{
				break ForEnd
			}
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
	fmt.Println("over")
}

func map_test() {
	defer func() {
		log.Debug("map_test func end.")
	}()

	{
		var m map[string]string
		m = make(map[string]string)
		m["1"] = "1111111111111"
		m["2"] = "2222222222222"
		m["3"] = "3333333333333"

		for k, v := range m {
			fmt.Println(k, v)
		}

		c, ok := m["1"]
		if ok {
			log.Debug("c: %v", c)
		} else {
			log.Error("error")
		}

		if c1, ok1 := m["1"]; ok1 {
			log.Debug("c1: %v", c1)
		} else {
			log.Error("error")
		}
	}

	{
		m := make(map[string]interface{})
		m["1"] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		n := make(map[string]string)
		n["1"] = "1111111111111111"
		n["2"] = "2222222222222222"
		n["3"] = "3333333333333333"
		n["4"] = "4444444444444444"
		m["2"] = n
		m["3"] = "11111111111111111"

		log.Debug("m: %v", m)
		str, err := json.Marshal(m)
		if err == nil {
			log.Debug("json: %s", str)
		}

		o := make(map[string]interface{})
		if err := json.Unmarshal(str, o); err == nil {
			log.Debug("o: %v", o)
		}
	}
}

func getName() (firstName, middleName, lastName, nickName string) {
	firstName = "May"
	nickName = "M"
	lastName = "Chen"
	nickName = "Babe"
	return firstName, nickName, lastName, nickName
}

func calc_sum(values []int, chanResult chan int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	sum := 0

	for _, value := range values {
		sum += value
	}
	chanResult <- sum
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var MakecoreData *int = nil
	//	MakecoreData = new(int)
	*MakecoreData = 10000
	fmt.Println(*MakecoreData)
	fmt.Println("hello world")

	return

	//	runtime.GOMAXPROCS(runtime.NumCPU())

	//	var args = os.Args
	//	log.Debug(len(args))
	//	log.Debug(args)

	{
		type ColorGroup struct {
			ID     int
			Name   string
			Colors []string
		}

		group := ColorGroup{
			ID:     1,
			Name:   "Reds",
			Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		}

		b, err := json.Marshal(group)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println(string(b))

		var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
		b, err = json_iterator.Marshal(group)
		if err != nil {

		}
		fmt.Println(string(b))
	}

	return

	//	for_test()

	//	if_test()

	//	switch_test()

	sclie_test()

	array_test()

	interface_test()

	map_test()

	//	_, _, lastName, nickName := getName()
	//	log.Debug("lastName: %s nickName: %s", lastName, nickName)

	//	{
	//		x, y := func(i, j int) (m, n int) { // x y 为函数返回值
	//			return j, i
	//		}(1, 9) // 直接创建匿名函数并执行
	//		fmt.Println(x, y)
	//	}

	{
		values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		chanResult := make(chan int, 2)
		go calc_sum(values[:len(values)/2], chanResult)
		go calc_sum(values[len(values)/2:], chanResult)
		sum1 := <-chanResult
		sum2 := <-chanResult
		fmt.Printf("%d + %d = %d\n", sum1, sum2, sum1+sum2)
	}

	//	{
	//		var ccc = 0
	//		go func(c int) {
	//			for {
	//				fmt.Printf("%d\n", c)
	//				c = c + 1
	//				time.Sleep(time.Microsecond * 1000)
	//			}
	//		}(ccc)
	//	}

	chan_test()
}
