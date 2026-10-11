package main

import (
	"fmt"
	"sync"
	"time"
)

// 在 goroutine 中 操作管道

// 需求：定义两个方法，一个方法给管道里面写入数据，一个读取管道里面的数据，要求写入和读取 同步进行。

// 1. 开启一个 fn1 的协程，向管道 inChan 中写入 10条 数据。
// 2. 开启一个 fn2 的协程，从管道 inChan 中读取数据。
// 3. 注意：fn1 和 fn2 同时操作 同一个管道。
// 4. 主线程必须等待操作完成后才可以退出。

var wg9_1 = sync.WaitGroup{}

// 写入数据
func fn9_1(ch chan int) {
	defer wg9_1.Done()

	// 方式一: 通过 defer 关闭管道，保证写入完成后关闭管道
	defer close(ch) // 写入完成后关闭管道

	for i := 1; i <= 10; i++ {
		fmt.Printf("fn9_1 写入数据：%d \n", i)

		ch <- i

		time.Sleep(time.Second * 1) // 模拟写入数据的耗时
	}

	// 方式二：直接关闭管道
	// close(ch) // 写入完成后关闭管道
}

// 读取数据
func fn9_2(ch chan int) {
	defer wg9_1.Done()

	// for range 循环遍历管道，读取数据。好处：不用处理 管道里到底有多少条数据。(管道是安全的)
	for val := range ch {
		fmt.Printf("fn9_2 读取数据：%d \n", val)
	}
}

func main() {
	ch1 := make(chan int, 10)

	wg9_1.Add(1) // Add 必须放在 goroutine 启动之前
	go fn9_1(ch1)

	wg9_1.Add(1)
	go fn9_2(ch1)

	// 等待协程执行完成
	wg9_1.Wait()

	fmt.Println("main 函数执行完成，程序退出！")
}
