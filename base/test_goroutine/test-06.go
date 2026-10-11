package main

import (
	"fmt"
	"time"
)

// goroutine 里 使用 defer 和 recover 捕获 panic

func test6_1() {
	for range 10 {
		time.Sleep(time.Millisecond * 50)

		fmt.Println("hello world")
	}
}

func test6_2() {
	defer func() {
		// 捕获 test6_2 内部抛出的 panic, 就不会影响其它 goroutine 的执行

		if err := recover(); err != nil {
			fmt.Println("test6_2 panic err ==>", err) // test6_2 panic err ==> test6_2 panic error
		}
	}()

	// 抛出一个错误
	panic("test6_2 panic error")

	// var myMap map[string]string
	// myMap[0] = "golang"
}

func main() {
	go test6_1()

	go test6_2()

	// 防止主线程退出程序，这里使用 time.Sleep 演示。
	time.Sleep(time.Second * 3)

	fmt.Println("main goroutine exit...")
}
