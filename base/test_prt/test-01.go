package main

import "fmt"

// 指针

// 变量的本质就是给存储数据的内存地址起了一个好记的别名。
// 指针也是一个变量，但它是一种特殊的变量，它存储的数据不是一个普通的值，而是另一个变量的内存地址。

func main() {
	// 1. Go 里边变量都有一个对应的内存地址(包括指针变量)

	var a1 = 10
	var b1 = &a1 // &: 取地址符，表示取变量 a1 的内存地址。指针变量 b1 的类型是 *int，表示它是一个指向 int 类型变量的指针。

	fmt.Printf("a1: %v, type: %T, address: %p \n", a1, a1, &a1) // a1: 10, type: int, address: 0x4c4729b16020
	fmt.Printf("b1: %v, type: %T, address: %p \n", b1, b1, &b1) // b1: 0x4c4729b16020, type: *int, address: 0x4c4729b02028

	// 2. 取指针变量的值

	var c1 = *b1                                                // *: 取值符，表示取指针变量 b1 指向的内存地址的值。c1 的类型是 int，表示它是一个 int 类型变量。
	fmt.Printf("c1: %v, type: %T, address: %p \n", c1, c1, &c1) // c1: 10, type: int, address: 0x147fcaa16040

	// 3. 指针变量的值是可以改变的

	*b1 = 20                  // 改变指针变量 b1 指向的内存地址的值，也就是改变了变量 a1 的值。
	fmt.Println("a1 ==>", a1) // a1 ==> 20

	fmt.Println("----------------分割线---111-----------------")
}
