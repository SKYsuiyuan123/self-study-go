package main

import "fmt"

// 函数

// 1. 具名函数 (具名函数里边不可以再嵌套 具名函数)
func calc() {
	fmt.Println("calc 执行了。")

	// 报错
	// func a() {}
	// a()
}

// 2. 函数 —— 形参 (a int, b int) 可以简写为 (a, b int)
func calcFn1(a int, b int) int {
	return a + b
}

// 3. 函数 —— 多返回值 (Go 的习惯：最后一个返回值通常是 error)
func calcFn2(a, b int) (int, int) {
	return a + b, a - b
}

// 4. 参数 —— 多类型 (泛型)
func calcFn3[T int | uint | float64](a, b T) (T, T) { // 参数: a ==> T, b ==> T
	return a + b, a - b
}

// 5. 函数 —— 可变参数 (可变参数必须放在最后一个参数)
func calcFn4(a int, b ...int) int { // 参数 b 是 int 类型的可变参数, b 是一个切片 []int
	fmt.Printf("b 的类型是: %T \n", b) // b 的类型是: []int

	sum := a
	for _, v := range b {
		sum += v
	}
	return sum
}

func calcFn4_2(b ...int) int { // 参数 b 是 int 类型的可变参数, b 是一个切片 []int
	fmt.Printf("b 的类型是: %T \n", b) // b 的类型是: []int

	sum := 0
	for _, v := range b {
		sum += v
	}
	return sum
}

// 6. 函数 —— 具名返回值
func calcFn5() (ok bool) {
	// return true // true

	// return // 返回的是 false, 因为 bool 的默认零值是 false。

	// var ok1 = false
	// return ok1 // 返回的是 ok1 的值

	// ok = true
	// return ok

	// 常用
	ok = false
	return
}

// 7. 返回值命名: 函数定义时，可以给返回值命名，并在函数体中直接使用这些变量，最后通过 return 关键字返回。
func calc1(x, y int) (sum int, sub int) { // 相当于 sum 和 sub 提前声明好了
	sum = x + y
	sub = x - y

	return // 返回的是 sum 和 sub 的值
}

func calc2(x, y int) (sum, sub int) { // 相当于 sum 和 sub 提前声明好了
	sum = x + y
	sub = x - y

	return // 返回的是 sum 和 sub 的值
}

// 8. 函数参数是 map 类型
func mapSort(m1 map[string]string) string {
	return ""
}

// 9. 递归 —— 求和 (n - 100之间)
func add1(n int) int {
	if n >= 100 {
		return n
	}

	return n + add1(n+1)
}

// 10. 递归 —— 打印从 1 - n 的所有整数
func printNum(n int) {
	fmt.Println("printNum n ==>", n)

	if 1 >= n {
		return
	}

	printNum(n - 1) // 递归调用
}

// 11. 递归 —— 阶乘
func factorial1(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial1(n-1) // 递归调用
}

// 12. 闭包 —— 实现计数器函数
func counterFn() func() int { // 返回的是一个函数
	// 普通变量会进入 栈里边，但是变成闭包里的变量后，就会进入 堆里边，直到闭包不再使用，才会被垃圾回收机制回收。

	// 优点：可以做到数据的私有化，外部无法直接访问 count 变量，只能通过闭包函数来访问和修改它。
	// 缺点：容易造成内存泄漏，因为闭包函数会持有对外部变量的引用，导致这些变量无法被垃圾回收机制回收。
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {
	// 1. 函数调用

	val1 := calcFn4_2(1, 2, 3)
	fmt.Println("val1 ==>", val1) // 6

	// 2. 数组传递给 可变参数

	arr1 := [3]int{3, 4, 5}
	fmt.Println("arr1 sum ==>", calcFn4_2(arr1[:]...)) // 12

	// 3. 切片传递给 可变参数
	slice1 := []int{3, 4, 5}
	fmt.Println("slice1 sum ==>", calcFn4_2(slice1...)) // 12

	val2 := calcFn4(1, 2, 3)
	fmt.Println("val2 ==>", val2) // 6

	val3, val4 := calc1(5, 6)
	fmt.Println("val3, val4 ==>", val3, val4) // 11, -1

	// calcFn2
	calcFn2Res1, calcFn2Res2 := calcFn2(10, 20)
	fmt.Println("calcFn2Res1, calcFn2Res2 ==>", calcFn2Res1, calcFn2Res2) //30, -10

	fmt.Println("----------------分割线---111-----------------")

	// 4. 函数返回值推断

	// 4.1 隐式推断
	sum3, sub3 := calcFn3(30.5, 10.2)
	fmt.Println("sum3, sub3 ==>", sum3, sub3) // 40.7, 20.3

	sum4, sub4 := calcFn3(-20, 10)
	fmt.Println("sum4, sub4 ==>", sum4, sub4) // -10, -30

	// 4.2 显示推断 (可以省略 泛型类型参数)
	sum5, sub5 := calcFn3[int](-20, 10)
	fmt.Println("sum5, sub5 ==>", sum5, sub5) // -10, -30

	calcFn5Res1 := calcFn5()
	fmt.Println("calcFn5Res1 ==>", calcFn5Res1) // false

	fmt.Println("----------------分割线---222-----------------")

	// 5. 匿名函数

	a1 := 2

	// 5.1 匿名 自执行函数 (立即执行函数)
	func() {
		a1 := 5
		fmt.Println("a1 ==>", a1) // 5
	}()

	a1 = 3
	fmt.Println("a1 ==>", a1) // 3

	// 5.2 匿名函数赋值给一个变量
	fun1 := func() {
		fmt.Println("fun1 running...")
	}

	fun1() // fun1 running...

	fmt.Println("----------------分割线---333-----------------")

	// 6. 递归函数

	add1Res := add1(1)
	fmt.Println("add1Res ==>", add1Res) // 5050

	printNum(10) // 打印从 1 - 10 的所有整数

	f1Res := factorial1(5)
	fmt.Println("f1Res ==>", f1Res) // 120

	fmt.Println("----------------分割线---444-----------------")

	// 7. 闭包函数

	counterRes1 := counterFn() // 返回的 counterRes1 是一个函数，每调用一次 count 就会 +1。

	fmt.Println("counterRes1() ==>", counterRes1()) // 1
	fmt.Println("counterRes1() ==>", counterRes1()) // 2

	fmt.Println("----------------分割线---555-----------------")
}
