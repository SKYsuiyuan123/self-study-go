package main

import "fmt"

// 使用 make 创建切片

// make: 可用于 初始化 slice, map 和 channel

func main() {
	// 1. make 初始化

	var slice1 = make([]int, 2) // 只声明长度，那么容量跟长度一样大。

	fmt.Printf("slice1: %#v -- %T \n", slice1, slice1)                  // slice1: []int{0, 0} -- []int
	fmt.Printf("slice1: len: %v, cap: %v \n", len(slice1), cap(slice1)) // slice1: len: 2, cap: 2

	slice2 := make([]string, 5)

	fmt.Printf("slice2: %#v -- %T \n", slice2, slice2)                  // slice2: []string{"", "", "", "", ""} -- []string
	fmt.Printf("slice2: len: %v, cap: %v \n", len(slice2), cap(slice2)) // slice2: len: 5, cap: 5

	// 2. 空(零值)切片 扩容

	slice2 = append(slice2, "a")         // 会在零值的后边 追加
	fmt.Printf("slice2: %#v \n", slice2) // slice2: []string{"", "", "", "", "", "a"}

	fmt.Println("----------------分割线---111-----------------")

	// 3. make 初始化的同时 声明长度和容量 (长度不能比容量小)

	slice3 := make([]string, 3, 5)

	fmt.Printf("slice3: %#v \n", slice3)                                // slice3: []string{"", "", ""}
	fmt.Printf("slice3: len: %v, cap: %v \n", len(slice3), cap(slice3)) // slice3: len: 3, cap: 5
}
