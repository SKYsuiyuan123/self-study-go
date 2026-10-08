package main

import "fmt"

// 类型别名 与 类型定义

type myInt1 = int // 类型别名 (type alias)
type myInt2 int   // 类型定义 (type definition) (new type)

// 自定义方法类型
type calcType1 func(int, int) int // 表示定义一个 calc 的函数类型

func add1_1(x, y int) int {
	return x + y
}

func sub1_1(x, y int) int {
	return x - y
}

func test1_1() {}

// 1. 函数作为另一个函数的参数

func calc1_1(x, y int, cb calcType1) int {
	return cb(x, y)
}

// 2. 函数作为返回值

func calc1_2(x string) calcType1 {
	switch x {
	case "+":
		return add1_1
	case "-":
		return sub1_1
	case "*":
		return func(a, b int) int { // 返回一个 匿名函数
			return a * b
		}
	}

	return func(a, b int) int { // 返回一个 匿名函数
		return 0
	}
}

func main() {
	// 1. 类型匹配才能赋值

	var c1 calcType1

	c1 = add1_1

	// c1 = test1_1 // test1_1 不满足 calcType1 的类型约束。cannot use test1_1 (value of type func()) as calcType1 value in assignment

	c1 = sub1_1 // 类型匹配才能赋值

	fmt.Printf("c1: %#v \n", c1)     // c1: (main.calcType1)(0x100d1cc10)
	fmt.Printf("c1 type: %T \n", c1) // c1 type: main.calcType1

	fmt.Println("c1(10, 5) ==> ", c1(10, 5)) // 5

	fmt.Println("----------------分割线---111-----------------")

	// 2. 自动类型推导

	var f1 = sub1_1
	var f2 calcType1 = sub1_1

	fmt.Printf("f1: %#v, type: %T \n", f1, f1) // f1: (func(int, int) int)(0x102b64c10), type: func(int, int) int

	f1 = c1

	fmt.Println("f1(10, 2) ==>", f1(10, 2)) // 8
	fmt.Println("f2(10, 2) ==>", f2(10, 2)) // 8

	fmt.Println("----------------分割线---222-----------------")

	// 3. 类型别名 与 类型定义

	var a int = 10
	var b myInt1 = 20 // myInt1 是 int 的类型别名，所以可以直接赋值。
	var d myInt2 = 30 // myInt2 是 int 的新类型，后续使用时需要进行类型转换。

	fmt.Printf("b val: %v type: %T \n", b, b) // b val: 20 type: int
	fmt.Println("a + b ==>", a+b)             // 30
	fmt.Println("a + d ==>", myInt2(a)+d)     // 40, 必须将 a 转换为 myInt2 类型才能进行运算
	fmt.Println("a + d ==>", a+int(d))        // 40, 必须将 d 转换为 int 类型才能进行运算
	fmt.Println("b + d ==>", b+int(d))        // 50

	fmt.Println("----------------分割线---333-----------------")

	// 4. 函数作为另一个函数的参数传递

	// 4.1 命名函数作为参数传递
	sum1 := calc1_1(1, 2, add1_1)

	fmt.Println("sum1 ==>", sum1) // 3

	// 4.2 匿名函数作为参数传递
	sum2 := calc1_1(1, 2, func(a, b int) int {
		return a * b
	})

	fmt.Println("sum2 ==>", sum2) // 2

	fmt.Println("----------------分割线---444-----------------")

	// 5. 函数作为返回值

	sum3 := calc1_2("-")                        // 返回一个函数类型
	fmt.Println("sum3(10, 5) ==>", sum3(10, 5)) // 5
}
