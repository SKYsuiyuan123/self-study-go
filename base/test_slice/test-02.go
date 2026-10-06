package main

import "fmt"

// 切片 —— 循环

func main() {

	slice1 := []string{"a", "b", "c", "d"}

	// 1. 使用 for 遍历 切片

	for i := 0; i < len(slice1); i++ {
		fmt.Println("slice1[i] val ==>", slice1[i])
		// slice1[i] val ==> a
		// slice1[i] val ==> b
		// slice1[i] val ==> c
		// slice1[i] val ==> d
	}

	fmt.Println("----------------分割线---111-----------------")

	// 2. 使用 range 遍历 切片

	// 2.1 遍历 string 切片
	for index, item := range slice1 {
		fmt.Printf("slice1 index: %v, item: %v \n", index, item)
		// slice1 index: 0, item: a
		// slice1 index: 1, item: b
		// slice1 index: 2, item: c
		// slice1 index: 3, item: d
	}

	// 2.2 遍历 byte 切片
	slice2 := []byte{'a', 'b', 'c'}
	for _, item := range slice2 {
		fmt.Printf("slice1 item: %v \n", item)
		// slice1 item: 97
		// slice1 item: 98
		// slice1 item: 99
	}

	// 3.3 遍历 nil 切片
	var slice3 []string

	for index, item := range slice3 { // 可以遍历，不会报错。但是遍历没有内容。✅ range nil slice 一次都不循环，不 panic
		fmt.Println("slice3 index, item ==>", index, item)
	}
}
