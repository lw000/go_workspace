// demo0002 project main.go
package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

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

func getName() (firstName, middleName, lastName, nickName string) {
	firstName = "May"
	nickName = "M"
	lastName = "Chen"
	nickName = "Babe"
	return firstName, nickName, lastName, nickName
}

func modify(array [5]int) {
	array[0] = 10 // 试图修改数组的第一个元素
	fmt.Println("In modify(), array values:", array)
}

func calc_sum(values []int, chanResult chan int) {
	sum := 0

	for _, value := range values {
		sum += value
	}
	chanResult <- sum
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var args = os.Args
	log.Debug(len(args))
	log.Debug(args)

	for_test()

	if_test()

	switch_test()

	sclie_test()

	_, _, lastName, nickName := getName()
	fmt.Println(fmt.Sprintf("lastName: %s nickName: %s", lastName, nickName))

	array := [5]int{1, 2, 3, 4, 5} // 定义并初始化一个数组
	modify(array)
	fmt.Println(array)

	{
		x, y := func(i, j int) (m, n int) { // x y 为函数返回值
			return j, i
		}(1, 9) // 直接创建匿名函数并执行
		fmt.Println(x, y)
	}

	//	{
	//		var numbers []int
	//		numbers = make([]int, 5, 5)
	//		numbers = append(numbers, 1, 2, 3, 4, 5)

	//		var numbers1 []int
	//		numbers1 = make([]int, 5, 5)
	//		numbers1 = []int{10, 20, 30, 40, 50}
	//		copy(numbers, numbers1)
	//		fmt.Println(numbers, len(numbers), cap(numbers), numbers1)
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

	//	var c1, c2, c3 chan int
	//	var i1, i2 int
	//	select {
	//	case i1 = <-c1:
	//		fmt.Printf("received ", i1, " from c1\n")
	//	case c2 <- i2:
	//		fmt.Printf("sent ", i2, " to c2\n")
	//	case i3, ok := (<-c3): // same as: i3, ok := <-c3
	//		if ok {
	//			fmt.Printf("received ", i3, " from c3\n")
	//		} else {
	//			fmt.Printf("c3 is closed\n")
	//		}
	//	default:
	//		fmt.Printf("no communication\n")
	//	}

	//	var ccc = 0

	//	for {
	//		fmt.Printf("%d\n", ccc)
	//		ccc = ccc + 1
	//		time.Sleep(time.Microsecond * 50)
	//
	//	}

	var m map[string]string
	m = make(map[string]string)
	m["1"] = "1111111111111"
	m["2"] = "1111111111111"
	m["3"] = "1111111111111"

	for k, v := range m {
		fmt.Println(k, v)
	}

	{
		chm := make(chan int)
		chn := make(chan int)
		quit := make(chan bool)
		go func(m chan int, n chan int) {
			i := 0
			j := 100
			for i < 100 {
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
					fmt.Print("m : %d", m)
				}
			case n := <-chn:
				{
					fmt.Print("n : %d", n)
				}
			case <-quit:
				{
					break ForEnd
				}
			default:
				time.Sleep(time.Millisecond * 10)
			}
		}

	}
	fmt.Println("over")
}
