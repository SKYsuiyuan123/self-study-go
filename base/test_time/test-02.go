package main

import (
	"fmt"
	"time"
)

// time

// 延迟求和函数

// 方式一：先延时，然后返回一个 延时后的函数，后边计算就不会再延时了。
func awaitAdd(second int) func(...int) int {
	time.Sleep(time.Second * time.Duration(second)) // 注意：这里的 second 是 int 类型，需要转换为 time.Duration 类型。

	return func(numList ...int) (sum int) {
		for _, item := range numList {
			sum += item
		}

		return sum
	}
}

// 方式二：直接返回一个函数，计算的时候再延时。
func awaitAdd2(second int) func(...int) int {
	return func(numList ...int) (sum int) {
		time.Sleep(time.Second * time.Duration(second))

		for _, item := range numList {
			sum += item
		}

		return
	}
}

func main() {
	// 方式 1

	addFn1 := awaitAdd(2)                // 延时两秒后返回一个函数
	fmt.Printf("addFn1: %#v \n", addFn1) // addFn1: (func(...int) int)(0x102669b80)

	t1 := time.Now()

	res1 := addFn1(1, 2, 3)
	fmt.Println("res1 ==>", res1) // res1 ==> 6 (延时 2 秒后才会输出)

	subTime1 := time.Since(t1)            // 计算从 t1 到现在的时间间隔，返回 Duration 类型的值.
	fmt.Println("subTime1 ==>", subTime1) // subTime1 ==> 21.542µs

	fmt.Println("----------------分割线---111-----------------")

	// 方式 2

	addFn2 := awaitAdd2(2)               // 立刻返回一个函数，计算的时候再延时。
	fmt.Printf("addFn2: %#v \n", addFn2) // addFn2: (func(...int) int)(0x10449dce0)

	t2 := time.Now()

	res2 := addFn2(1, 2, 3)       // 计算的时候再延时。
	fmt.Println("res2 ==>", res2) // 6

	subTime2 := time.Since(t2)
	fmt.Println("subTime2 ==>", subTime2) // subTime2 ==> 2.000123456s (延时 2 秒后才会输出)
}
