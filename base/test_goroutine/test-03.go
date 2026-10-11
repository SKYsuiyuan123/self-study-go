package main

import (
	"fmt"
	"sync"
	"time"
)

// 批量开启多个协程

var wg3_1 sync.WaitGroup

func test3_1(num int) {
	defer wg3_1.Done() // 在函数结束时调用 Done()，表示该协程完成。

	for i := range 5 {
		fmt.Println("协程 ==>", num, i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	for i := 0; i < 6; i++ {
		wg3_1.Add(1) // 在开启协程前调用 Add(1)，表示有一个协程需要等待。

		go test3_1(i) // 开启协程
	}

	wg3_1.Wait() // 等待所有协程完成

	fmt.Println("所有协程执行完毕")
}
