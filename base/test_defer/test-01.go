package main

import "fmt"

// defer

// defer 是栈，先进后出。
// 先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。

// Go 语言中的 defer 语句会将其后面跟随的语句进行延迟处理。在 defer 归属的函数即将返回时，将延迟处理的语句按照 defer 定义的 逆序进行执行。

// Go 语言中的 defer 语句会将函数推迟到外层函数返回之后执行。被推迟的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。

func fn1() {
	fmt.Println("开始 fn1")

	defer fmt.Println("fn1 延迟执行 1")
	defer fmt.Println("fn1 延迟执行 2")
	defer fmt.Println("fn1 延迟执行 3")

	fmt.Println("结束 fn1")

	// 打印顺序：开始 fn1, 结束 fn1, fn1 延迟执行 3, fn1 延迟执行 2, fn1 延迟执行 1.
}

func fn2() {
	fmt.Println("开始 fn2")

	defer func() {
		fmt.Println("111")
	}()

	defer func() {
		fmt.Println("222")
	}()

	defer func() {
		fmt.Println("333")
	}()

	fmt.Println("结束 fn2")

	// 打印顺序：开始 fn2, 结束 fn2, 333, 222, 111.
}

// defer 匿名函数，注意区分 命名返回值 (fn4) 和 匿名返回值 (fn3)

func fn3() int {
	var a int

	defer func() {
		a++
		fmt.Println("fn3 延迟执行, a = ", a) // fn3 延迟执行, a =  1
	}()

	fmt.Println("结束 fn3, a = ", a) // 结束 fn3, a =  0

	return a // 返回值为 0, 因为 a++ 在 return 之后执行

	// 执行顺序：结束 fn3, a =  0, fn3 延迟执行, a =  1, fn3Res ==> 0
}

// 在 Go 的函数中 return 语句的执行顺序是先计算返回值，然后执行 defer 语句，最后才真正返回。
// 在 Go 的函数中 return 语句在底层是分为两步的：第一步是将返回值赋值给返回变量，第二步是执行 defer 语句。
// 因此，如果在函数中使用了命名返回值，那么在 return 语句执行时，命名返回值已经被初始化了，而 defer 语句会在 return 之后执行，这就可能导致返回值被修改。

// 在 Go 的函数中 return 语句在底层并不是原子操作，它分为 给返回值赋值 和 RET 指令两步。
// 而 defer 语句执行时机就在 返回值赋值操作后，RET 指令执行前。

func fn4() (a int) { // 相当于 a 提前初始化了，之后执行 defer 给 a 做了 ++ 的动作。
	defer func() {
		a++
		fmt.Println("fn4 延迟执行, a = ", a) // fn4 延迟执行, a =  1
	}()

	fmt.Println("结束 fn4, a = ", a) // 结束 fn4, a =  0

	return a // 返回值为 1, 因为 a++ 在 return 之前执行
}

func fn5() (y int) {
	x := 5

	defer func() {
		x++
	}()

	return x // 5, 因为 return x 先将 x 的值赋给 y，然后执行 defer，最后返回 y 的值。
}

func fn6() (x int) {
	defer func() {
		x++ // 这里的 x 是返回时的 x，defer 会在 return 之后执行，因此 返回的是 6。
	}()

	return 5 // 6, 因为 return 5 先将 5 的值赋给 x，然后执行 defer，最后返回 x 的值。
}

func fn7() (x int) {
	defer func(x int) {
		x++ // 匿名函数里边的 x 跟外边的 x 无关，里边的 x++ 后 x 是 1.
	}(x) // defer 注册要延迟执行的函数时，该函数所有的参数都需要明确其值。

	return x // 0, 并没有改变返回值，因为 defer 的参数是值传递，x 的值在 defer 执行时已经被传递了。
}

func fn8() (x int) {
	defer func(x int) {
		fmt.Println("fn8 x ==>", x) // 0
		x++                         // 同上
	}(x)

	x += 3 // 不会影响 匿名函数里的变量 x 的值

	return 5 // 5, 这里相当于把 5 赋值给 x, 但是 defer 里对 x 做的 ++ 操作 并不能影响外部的 x。
	// return // 只有 return 会返回 x 的值，x 的值是 3, 因为 defer 里对 x 做的 ++ 操作 并不能影响外部的 x。
}

func main() {
	fn1()

	fn2()

	fmt.Println("----------------分割线---111-----------------")

	fn3Res := fn3()
	fmt.Println("fn3Res ==>", fn3Res) // 0

	fn4Res := fn4()
	fmt.Println("fn4Res ==>", fn4Res) // 1

	fn5Res := fn5()
	fmt.Println("fn5Res ==>", fn5Res) // 5

	fn6Res := fn6()
	fmt.Println("fn6Res ==>", fn6Res) // 6

	fn7Res := fn7()
	fmt.Println("fn7Res ==>", fn7Res) // 0

	fn8Res := fn8()
	fmt.Println("fn8Res ==>", fn8Res) // 5

	fmt.Println("----------------分割线---222-----------------")
}
