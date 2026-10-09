package main

import (
	"fmt"
	"strings"
)

// 数据类型
// 	基础类型
// 		bool (布尔)
// 		int8 int16 int64... (整型)
// 		float32 float64 (浮点型)
// 		complex64 complex128 (复数型)
// 		rune (字符型)
// 		string (字符串)
// 		error (错误型)
// 	复合类型
// 		array (数组)
// 		struct (结构体)
// 		slice (切片)
// 		map (字典)
// 		chan (通道)
// 		pointer	(指针)
// 		interface (接口)

// 值类型：数组、基本数据类型、结构体
// 引用类型：切片、map, func, interface

// 字符串的本质：不可变的字节序列。默认使用 UTF-8 编码存储文本。
// 字节序列是一串连续排列的字节，并且每个字节都是 byte 类型。

// 默认零值
// 零值：如果给一个基本数据类型只声明不赋值，那么这个变量的值就是对应类型的零值。

// bool -> false
// number -> 0
// string -> ""
// pointers -> nil (指针)
// array -> nil
// slices -> nil (切片) (var s2 []int 是 nil, 但是 s1 := []int{} 不是)
// maps -> nil
// channel -> nil
// functions -> nil
// interface -> nil

func defaultZero() {
	var a1 int
	var a2 float32
	var a3 string
	var a4 bool
	var a5 []int
	var a6 map[string]string

	fmt.Printf("a1 ==> %#v \n", a1) // a1 ==> 0
	fmt.Printf("a2 ==> %#v \n", a2) // a2 ==> 0
	fmt.Printf("a3 ==> %#v \n", a3) // a3 ==> ""
	fmt.Printf("a4 ==> %#v \n", a4) // a4 ==> false
	fmt.Printf("a5 ==> %#v \n", a5) // a5 ==> []int(nil)
	fmt.Printf("a6 ==> %#v \n", a6) // a6 ==> map[string]string(nil)

	fmt.Println("a5 == nil ==>", a5 == nil) // a5 == nil ==> true
	fmt.Println("a6 == nil ==>", a6 == nil) // a6 == nil ==> true

	fmt.Println("----------------分割线---defaultZero-----------------")
}

func main() {
	var a = 1
	var b = 2

	// 1. 交换两个值

	a, b = b, a

	fmt.Println("a, b ==>", a, b) // 2, 1

	fmt.Println("----------------分割线---111-----------------")

	// 2. 打印默认零值

	defaultZero()

	// 3. 除法
	// 当两个数字都是整数，相除后会得到 一个整数，去掉了小数部分

	res1 := 2 / 3

	fmt.Println("2 / 3 ==>", res1) // 0

	// 4. a++ 只能单独使用,不能赋值给其它使用。

	res1++
	// res2 := res1++ // 不符合语法

	fmt.Println("res1 ==>", res1) // 1

	// 5. 浮点数相除

	var f1 float64 = 100
	f2 := (f1 - 32) / 1.8 // 浮点数 减 整数 等于 浮点数

	fmt.Printf("f2 ==> %.2f \n", f2) // 37.78

	// 6. 字符串重复

	for i := 1; i < 4; i++ {
		fmt.Println(strings.Repeat("*", i))
		// *
		// **
		// ***
	}
}
