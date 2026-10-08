package main

import "fmt"

// defer 注册要执行的函数时，该函数所有的参数都需要确定其值。

func calc1(key string, a, b int) int {
	res := a + b

	fmt.Printf("key: %s, a: %d, b: %d, res: %d \n", key, a, b, res)

	return res
}

func foo() int {
	a, b := 3, 4
	c := a + b

	defer fmt.Println("111 ==>", c) // 7, defer 注册时，c 的值为 7

	fmt.Println("c ==>", c) // 7

	defer fmt.Println("222 ==>", c) // 7, defer 注册时，c 的值为 7

	defer func() {
		fmt.Println("333 ==>", c) // 100, defer 注册时，c 的值为 7，但在 defer 执行时，c 的值已经被修改为 100.
	}()

	c = 100

	return c

	// 打印顺序：
	// 	c ==> 7
	// 	333 ==> 100
	// 	222 ==> 7
	// 	111 ==> 7
}

func main() {
	x := 1
	y := 2

	defer calc1("AA", x, calc1("A", x, y))

	x = 10

	defer calc1("BB", x, calc1("B", x, y))

	y = 20

	fmt.Printf("main x: %d, y: %d \n", x, y)

	// defer 注册顺序
	// 	1. defer calc1("AA", x, calc1("A", x, y))
	//  2. defer calc1("BB", x, calc1("B", x, y))

	// defer 执行顺序
	//  1. defer calc1("BB", x, calc1("B", x, y))
	//  2. defer calc1("AA", x, calc1("A", x, y))

	// 执行步骤分析:
	//  1. 执行非 defer calc1("A", x, y) => x = 1, y = 2 打印 ==> "A", 1, 2, 3
	//  2. 执行非 defer calc1("B", x, y) => x = 10, y = 2 打印 ==> "B", 10, 2, 12
	// 	3. main x: 10, y: 20
	// 	4. 执行 defer calc1("BB", x, 12) => x = 10, y = 12 打印 ==> "BB", 10, 12, 22
	//  5. 执行 defer calc1("AA", x, 3) ==> x = 1, y = 3 打印 ==> "AA", 1, 3, 4

	// 打印结果:
	// key: A, a: 1, b: 2, res: 3
	// key: B, a: 10, b: 2, res: 12
	// main x: 10, y: 20
	// key: BB, a: 10, b: 12, res: 22
	// key: AA, a: 1, b: 3, res: 4

	fmt.Println("----------------分割线---111-----------------")

	foo()
}
