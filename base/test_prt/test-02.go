package main

import "fmt"

// new 和 make

// 共同点：
//  1. new 和 make 都是内建函数，都是用来分配内存的。
//  2. new 和 make 都是用来创建引用类型的对象的，但是它们创建的对象类型不同。

// 区别：
//  1. new 用于值类型和引用类型，make 只能用于引用类型。
//  2. new 返回的是指针类型，make 返回的是引用类型。
//  3. new 分配的内存是零值初始化的，make 分配的内存是非零值初始化的。
//  4. new 返回的是指向类型的指针，make 返回的是类型本身。
//  5. new 不能用于 slice、map 和 channel 类型，make 可以用于 slice、map 和 channel 类型。
//  6. new 只会分配内存，并不会初始化内存。make 会分配内存，并初始化内存。

func main() {
	// 1. new

	// 1.26+
	var d1 *int = new(10) // new 返回的是指针类型，d1 的类型是 *int，表示它是一个指向 int 类型变量的指针。
	*d1 = 100

	fmt.Printf("d1: %v, type: %T, address: %p \n", d1, d1, &d1) // d1: 0x3a123a688020, type: *int, address: 0x3a123a694020

	// 等价于
	f1 := new(int(10))
	fmt.Printf("f1: %v, type: %T, address: %p \n", f1, f1, &f1) // f1: 0x1fe099b96030, type: *int, address: 0x1fe099b82038

	f2 := new(int64(20))                                        // int64 要使用特殊写法，不能直接用 new(10)
	fmt.Printf("f2: %v, type: %T, address: %p \n", f2, f2, &f2) // f2: 0x669161a1c020, type: *int64, address: 0x669161a00038

	// 1.26 以前
	var e1 *int
	e1 = new(int)

	*e1 = 101
	fmt.Printf("e1: %v, type: %T, address: %p \n", e1, e1, &e1) // e1: 0x3a0b4516030, type: *int, address: 0x3a0b4502038

	fmt.Println("----------------分割线---111-----------------")

	// 2. make

	// 2.1 make 初始化 map

	userInfo1 := make(map[string]string)

	fmt.Printf("userInfo1: %#v \n", userInfo1)            // userInfo1: map[string]string{}
	fmt.Println("userInfo1 == nil ==>", userInfo1 == nil) // userInfo1 == nil ==> false

	userInfo1["username"] = "张三"               // make 初始化时 已经分配了内存空间，因此可以直接 赋值。
	fmt.Printf("userInfo1: %#v \n", userInfo1) // userInfo1: map[string]string{"username":"张三"}

	// 2.2 make 初始化 slice

	slice1 := make([]string, 4)

	fmt.Printf("slice1: %#v \n", slice1)            // slice1: []string{"", "", "", ""}
	fmt.Println("slice1 == nil ==>", slice1 == nil) // slice1 == nil ==> false

	slice1[1] = "b"                      // make 初始化时 已经分配了内存空间
	fmt.Printf("slice1: %#v \n", slice1) // slice1: []string{"", "b", "", ""}

	// 2.3 slice 必须分配内存空间

	var slice2 []string
	fmt.Printf("slice2: %#v \n", slice2)            // slice2: []string(nil)
	fmt.Println("slice2 == nil ==>", slice2 == nil) // slice2 == nil ==> true

	// slice2[1] = "b" // 报错，必须先初始化分配内存空间。

	// 借助 append
	slice2 = append(slice2, "a")

	fmt.Printf("slice2: %#v \n", slice2)            // slice2: []string{"a"}
	fmt.Println("slice2 == nil ==>", slice2 == nil) // slice2 == nil ==> false
}
