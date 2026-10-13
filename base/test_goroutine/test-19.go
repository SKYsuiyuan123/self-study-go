package main

import (
	"fmt"
	"time"
)

// 一个 go 协程写入数据，另一个 go 协程读取数据.

// 读取数据的 go 协程先启动，会阻塞等待数据写入.
// 写入成功后 读取数据的 go 协程会被唤醒，读取数据成功后继续执行.

func main() {
	ch := make(chan int, 1)

	go func() {
		a := <-ch
		fmt.Printf("%d 读出 %d \n", time.Now().Unix(), a)
	}()

	time.Sleep(time.Second * 2)

	go func() {
		ch <- 4
		fmt.Printf("%d 写入 %d \n", time.Now().Unix(), 4)
	}()

	time.Sleep(time.Second * 1)
}
