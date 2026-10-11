package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// goroutine

// 使用 sync.WaitGroup 来等待所有的 goroutine 执行完毕。

// 一个 goroutine 最小的栈空间是 2KB，goroutine 的栈空间是可以动态扩展的，goroutine 的栈空间最大可以达到 1GB。

var wg2_1 sync.WaitGroup

func test2_1() {
	for i := range 10 {
		fmt.Println("test2_1 ==>", i)
		time.Sleep(time.Millisecond * 200)
	}

	wg2_1.Done() // 计数器减 1
}

func test2_2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2_2 ==>", i)
		time.Sleep(time.Millisecond * 100)
	}

	wg2_1.Done() // 计数器减 1
}

func main() {
	timeStart := time.Now()

	// 后去当前计算机上的 CPU 个数
	cpuNum1 := runtime.NumCPU()

	fmt.Println("cpuNum1 ==>", cpuNum1) // 12

	// 设置当前使用 多少个 cpu 运行程序
	runtime.GOMAXPROCS(cpuNum1)

	wg2_1.Add(1) // 计数器加 1
	go test2_1()

	wg2_1.Add(1) // 计数器加 1
	go test2_2()

	for i := range 10 {
		fmt.Println("main ==>", i)
		time.Sleep(time.Millisecond * 50)
	}

	// 使用 sync.WaitGroup 来等待所有的 goroutine 执行完毕。

	wg2_1.Wait() // 等待所有的 goroutine 执行完毕

	fmt.Println("主线程执行结束 ==>", time.Since(timeStart))
}
