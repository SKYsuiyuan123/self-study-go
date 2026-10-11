package main

import (
	"fmt"
	"sync"
	"time"
)

// go 中 默认定义的管道都是双向管道 (可读可写)

// 只写管道
// var chan1 chan<- int

// 只读管道
// var chan2 <-chan int

var wg11_1 sync.WaitGroup

func fn11_1(ch chan<- int) { // 只写管道
	defer wg11_1.Done()
	defer close(ch) // 关闭管道

	for i := 1; i <= 10; i++ {
		fmt.Printf("fn11_1 写入数据: %d \n", i)

		ch <- i

		time.Sleep(time.Second) // 每秒写入一次
	}
}

func fn11_2(chn <-chan int) { // 只读管道
	defer wg11_1.Done()

	for val := range chn {
		fmt.Println("fn11_2 读取数据 ==> ", val)
	}
}

func main() {
	// 1. 声明只写管道

	ch1 := make(chan<- int, 2) // 只写管道，缓冲区大小为2

	ch1 <- 10
	ch1 <- 12

	// <-ch1 // 只写管道 无法读取

	// 2. 声明只读管道

	ch2 := make(<-chan int, 2) // 只读管道，缓冲区大小为2
	// ch2 <- 10 // 只读管道 无法写入

	fmt.Printf("ch2 的值为: %#v, type: %T \n", ch2, ch2) // ch2 的值为: (<-chan int)(0x1a05806a0080), type: <-chan int

	fmt.Println("----------------分割线---111-----------------")

	var chan2 = make(chan int, 10)

	wg11_1.Add(1)
	go fn11_1(chan2) // 双向管道可以 传递给 只读 和 只写管道

	wg11_1.Add(1)
	go fn11_2(chan2)

	wg11_1.Wait()

	fmt.Println("main 结束。。。")
}
