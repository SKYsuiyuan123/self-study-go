package main

import (
	"fmt"
	"time"
)

// 一个 go 协程写入数据，另一个 go 协程读取数据.

// 读取数据的 go 协程先启动，会阻塞等待数据写入.
// 写入成功后 读取数据的 go 协程会被唤醒，读取数据成功后继续执行.

// 1: 当 channel 缓冲区大小为 1 时，写入数据的 go 协程 不会阻塞等待数据被读取后 再执行.而是直接写入数据后继续执行.
// 0: 当 channel 缓冲区大小为 0 时，写入数据的 go 协程 会阻塞等待数据被读取后才能继续执行.

// ch := make(chan int, 1) // 缓冲区大小为 1
// 打印：
// 1791877820
// 1791877820 写入 4
// 1791877822 读出 4

// ch := make(chan int, 0) // 缓冲区大小为 0
// 打印：
// 1791877871
// 1791877873 写入 4
// 1791877873 读出 4

func main() {
	ch := make(chan int, 0)
	fmt.Println(time.Now().Unix())

	go func() {
		ch <- 4
		fmt.Printf("%d 写入 %d \n", time.Now().Unix(), 4)
	}()

	time.Sleep(time.Second * 2)

	go func() {
		a := <-ch
		fmt.Printf("%d 读出 %d \n", time.Now().Unix(), a)
	}()

	time.Sleep(time.Second * 1)
}
