package main

import (
	"fmt"
	"time"
)

func test5_1() {
	fmt.Println("f1 finish...")

	go test5_2() // 开启协程 (test5_2 和 test5_1 并行执行)

	// test5_1 执行完毕之后，由于 test5_2 是在 test5_1 内部开启的协程，所以 test5_2 会在 test5_1 执行完毕后继续执行。
}

func test5_2() {
	time.Sleep(time.Second * 2)
	fmt.Println("f2 finish...") // f2 finish...
}

func main() {
	go test5_1() // 开启协程 (test5_1 和 main 并行执行)

	time.Sleep(time.Second * 3) // 等待 3 秒，确保 test5_2 执行完毕

	fmt.Println("main finish...") // main finish...
}
