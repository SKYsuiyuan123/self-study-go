package main

import "fmt"

// 管道 —— 遍历

// 使用 for range 循环 遍历管道
// 使用 for range 遍历管道，当管道被关闭的时候就会退出 for range, 如果没有关闭管道，for range 会一直阻塞等待数据写入管道 (死锁)。

func main() {
	// 1. for range 循环 遍历管道

	var ch1 = make(chan int, 10)

	for i := 1; i <= 10; i++ {
		ch1 <- i
	}

	// 写入完成后，关闭管道 (for range 遍历 必须关闭管道，for 循环可以不关闭管道。)

	close(ch1)

	// for range 循环遍历输出管道的值 (注意：管道没有 key)
	for val := range ch1 {
		fmt.Println("val ==>", val)

		// 如果没有关闭管道，则遍历完之后会报错。(fatal error: all goroutines are asleep - deadlock!)
	}

	fmt.Println("----------------分割线---111-----------------")

	ch2 := make(chan int, 5)

	for i := 20; i < 25; i++ {
		ch2 <- i
	}

	// 使用 for 循环 遍历管道 (for 循环遍历管道不需要关闭管道)

	for j := 0; j < 5; j++ {
		fmt.Println("ch2 val ==>", <-ch2)
	}
}
