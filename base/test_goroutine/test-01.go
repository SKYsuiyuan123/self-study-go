package main

import (
	"fmt"
	"time"
)

// goroutine

// 在一个 go 程序的主线程上可以起多个 goroutine (协程)。
// 多个 goroutine 可以并发执行，goroutine 是轻量级线程，goroutine 的创建和销毁的开销很小。

// 协程可以理解为一种用户态的线程，协程的调度完全由用户控制，协程的切换不需要内核参与，协程的切换开销很小。

// 协程开启: go test()

// 在主线程中(也可以理解为进程)，开启一个 goroutine，主线程和 goroutine 是并发执行的，主线程和 goroutine 之间是异步的。

// 开启一个 goroutine, 该协程每隔 100毫秒 输出 "hello world"。
// 在主线程中 也每隔 200毫秒 输出 "你好 go"，输出 10次后，退出程序。
// 要求 主线程 和 goroutine 同时执行。

func test1_1() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("hello world ==>", i)
	// 	time.Sleep(time.Millisecond * 100)
	// }

	// 优化 (推荐使用)
	for i := range 10 { // range 10 会返回 0~9 的索引值, i 每次循环都会自增 1。
		fmt.Println("hello world ==>", i)
		time.Sleep(time.Millisecond * 100)
	}

	// 或 range over int 更简洁的场景 (只想要“重复 N 次”)
	// for range 10 { // range 10 会返回 0~9 的索引值, 但是不需要使用索引值，所以使用 _ 来忽略索引值。
	// 	fmt.Println("hello world ==>")
	// 	time.Sleep(time.Millisecond * 100)
	// }
}

func main() {
	// 开启一个协程

	go test1_1()

	for i := range 10 {
		fmt.Println("你好 go ==>", i)
		time.Sleep(time.Millisecond * 200) // 这里时间一定要大于协程的时间，否则主线程执行完毕，程序退出，协程还没有执行完毕。
	}

	fmt.Println("主线程执行完毕，程序退出...") // 主线程 执行结束后，整个进程就终止了，协程也就终止了。
}
