package main

import (
	"fmt"
)

// 		场景														推荐
// 只重复 N 次，不用索引						for range N
// 简单递增计数，用索引							for i := range N 或传统写法都行
// 需要步长、倒序、复杂条件					 传统 for
// 遍历数组/切片/map/字符串					for range collection

func test4_1() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello ==>", i)
	}

	// 优化 (推荐使用)
	for i := range 10 { // range 10 会返回 0~9 的索引值, i 每次循环都会自增 1。
		fmt.Println("world ==>", i)
	}

	// 或 range over int 更简洁的场景 (只想要“重复 N 次”)
	for range 10 { // range 10 会返回 0~9 的索引值, 但是不需要使用索引值，所以使用 _ 来忽略索引值。
		fmt.Println("hello world ==>")
	}

	// range over int 的更多用法
	// for i := range 10 { ... }      // i 从 0 到 9
	// for range 10 { ... }           // 不要 i，只重复
	// for i := range n { ... }       // n 是变量也行

	n := -5 // 负数也可以, 但是不会循环。
	for range n {
		fmt.Println("hello ---")
	}
}

func main() {
	test4_1()

	fmt.Println("主线程执行完毕，程序退出...")
}
