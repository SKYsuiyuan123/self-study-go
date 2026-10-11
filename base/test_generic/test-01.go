package main

import "fmt"

// 泛型 —— 函数

func fn1[T int | float64](n1, n2 T) T {
	return n1 + n2
}

func fn2[T int | uint, K string](n1 T, n2 K) {
	fmt.Println("n1 ==>", n1, "n2 ==>", n2)
}

type Number1 interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func fn3[T Number1, K string](n1 T, n2 K) {
	fmt.Println("n1 ==>", n1, "n2 ==>", n2)
}

func main() {
	res1 := fn1(1, 2)
	fmt.Println("res1 ==>", res1) // res1 ==> 3

	num1, num2 := 1.1, 2.2

	res2 := fn1(num1, num2)
	fmt.Println("res2 ==>", res2) // res2 ==> 3.3000000000000003

	fmt.Println("----------------分割线---111-----------------")

	fn2(1, "hello") // n1 ==> 1 n2 ==> hello

	fn3(1, "hello") // n1 ==> 1 n2 ==> hello

	fn3(1.1, "world") // n1 ==> 1.1 n2 ==> world
}
