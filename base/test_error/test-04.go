package main

import (
	"fmt"
	"runtime/debug"
)

// 异常处理 —— 捕获

// debug.Stack() 可以打印堆栈信息

func fn4_1() {
	defer func() {
		fmt.Println("111") // "111" 会被打印出来

		if err := recover(); err != nil {
			fmt.Println("捕获到异常:", err) // runtime error: index out of range [2] with length 2
			fmt.Println("堆栈信息:")

			// 打印堆栈信息
			debug.PrintStack()

			// fmt.Println(string(debug.Stack()))

			// goroutine 1 [running]:
			// runtime/debug.Stack()
			//         /usr/local/go/src/runtime/debug/stack.go:26 +0x64
			// runtime/debug.PrintStack()
			//         /usr/local/go/src/runtime/debug/stack.go:18 +0x1c
			// main.fn4_1.func1()
			//         /Users/sky/Desktop/study_github/self-study-go/base/test_error/test-04.go:21 +0xc4
			// panic({0x102bd4d90?, 0xc4461c32138?})
			//         /usr/local/go/src/runtime/panic.go:859 +0x120
			// main.fn4_1()
			//         /Users/sky/Desktop/study_github/self-study-go/base/test_error/test-04.go:29 +0x38
			// main.main()
			//         /Users/sky/Desktop/study_github/self-study-go/base/test_error/test-04.go:33 +0x1c
		}
	}()

	var list []int = []int{1, 2}
	fmt.Println(list[2]) // 越界访问，这里会发生 panic.
}

func main() {
	fn4_1()

	fmt.Println("main 函数执行结束") // 会被打印出来，因为程序在 fn4_1 中发生了 panic 并被捕获，之后程序继续执行。
}
