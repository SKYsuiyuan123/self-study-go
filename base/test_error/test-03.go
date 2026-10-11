package main

import (
	"fmt"
	"os"
)

// 异常处理 —— 中断程序

func init() {
	_, err := os.ReadFile("test.txt")

	fmt.Println(err) // open test.txt: no such file or directory

	if err != nil {
		// log.Fatalln("Error reading file:", err) // 程序不再继续往下执行。2026/10/12 01:06:42 Error reading file: open test.txt: no such file or directory

		// 或 (只推荐用在 init() 函数中)
		panic(err) // panic: open test.txt: no such file or directory
	}
}

func main() {
	fmt.Println("main function...") // 不会执行
}
