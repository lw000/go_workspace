// demo0002 project main.go
package main

import (
	"fmt"
	"os"
	//	"time"

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

	var i = 20
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

func calc_sum(values []int, resultChannel chan int) {
	sum := 0

	for _, value := range values {
		sum += value
	}
	resultChannel <- sum
}

func main() {
	var args = os.Args
	log.Debug(len(args))
	log.Debug(args)

	for_test()

	if_test()

	_, _, lastName, nickName := getName()
	fmt.Println(fmt.Sprintf("lastName: %s nickName: %s", lastName, nickName))

	//	fmt.Println(time.Now())

	//	slice1 := []int{1, 2, 3, 4, 5}
	//	slice2 := []int{5, 4, 3}
	//	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中

	array := [5]int{1, 2, 3, 4, 5} // 定义并初始化一个数组
	modify(array)
	fmt.Println(array)

	xxx := 1
	switch xxx {
	case 0:
		fmt.Println("0")
		//break
	case 1:
		fmt.Println("1")
		//break
	case 2:
		fmt.Println("2")
		//break
	default:

	}
	//	{
	//		x, y := func(i, j int) (m, n int) { // x y 为函数返回值
	//			return j, i
	//		}(1, 9) // 直接创建匿名函数并执行
	//		fmt.Println(x, y)
	//	}

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
		resultChannel := make(chan int, 2)
		go calc_sum(values[:len(values)/2], resultChannel)
		go calc_sum(values[len(values)/2:], resultChannel)
		sum1, sum2 := <-resultChannel, <-resultChannel
		fmt.Println("Result: ", sum1, sum2, sum1+sum2)
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

	//	for {
	//		select {
	//			case
	//		}
	//	}
}
