package main

import "fmt"

// for range 键值遍历
// 可以用来遍历 数组、切片、字符串、map 以及 通道(channel)

// 通过 for range 遍历的返回值有以下规律：
// 	1. 数组、切片、字符串返回索引和值。
//  2. map 返回键和值。
//  3. 通过 (channel) 只返回通道内的值。

func main() {
	// 1. 遍历 字符串

	str1 := "你好 golang"

	for key, val := range str1 {
		fmt.Printf("key = %v, val = %c \n", key, val)
		// key = 0, val = 你
		// key = 3, val = 好
		// key = 6, val =
		// key = 7, val = g
		// key = 8, val = o
		// key = 9, val = l
		// key = 10, val = a
		// key = 11, val = n
		// key = 12, val = g
	}

	// 1.2 遍历字符串

	str2 := "味真足"
	for byIndex, uniCode := range str2 {
		fmt.Println("byIndex, uniCode ==>", byIndex, uniCode, string(uniCode))
		// byIndex, uniCode ==> 0 21619 味
		// byIndex, uniCode ==> 3 30495 真
		// byIndex, uniCode ==> 6 36275 足
	}

	// 1.3 字符串转成切片 然后遍历 []rune(str)

	for index, uniCode3 := range []rune(str2) {
		fmt.Println("index, uniCode3 ==>", index, uniCode3, string(uniCode3))
		// index, uniCode3 ==> 0 21619 味
		// index, uniCode3 ==> 1 30495 真
		// index, uniCode3 ==> 2 36275 足
	}

	fmt.Println("----------------分割线---111-----------------")

	// 2. 遍历 数组

	var arr1 = [3]string{"php", "java", "node"}

	for index, item := range arr1 {
		fmt.Println("index, item ==>", index, item)
		// index, item ==> 0 php
		// index, item ==> 1 java
		// index, item ==> 2 node
	}

	fmt.Println("----------------分割线---222-----------------")

	// 3. 遍历 map

	scores := map[string]int{
		"数学": 80,
		"语文": 90,
	}

	for key, val := range scores {
		fmt.Println("key, val ==>", key, val)
		// key, val ==> 数学 80
		// key, val ==> 语文 90
	}
}
