package main

import (
	"errors"
	"fmt"
)

// 使用 panic / recover 来处理异常

// (panic: 抛出异常, recover: 捕获异常)。
// panic 可以在任何地方引发， 但是 recover 只有在 defer 函数中才能捕获异常。

// defer 函数
//  1. defer 语句会将函数推迟到外层函数返回之后执行。
//  2. 关键字 defer 用于注册一个函数调用，该调用将在函数返回时执行。
//  3. 这些调用直到 return 前才被执行。因此，可以用来做一些清理工作，比如关闭文件、解锁资源等。
//  4. 多个 defer 语句会按照后进先出（LIFO）的顺序执行。谁先注册的 defer 语句会最后执行，谁后注册的 defer 语句会最先执行。
//     谁离 return 近谁先执行。
//  5. defer 语句可以用来处理异常情况，确保资源的释放和清理工作能够被执行。
//  6. defer 语句可以在函数中使用匿名函数来实现更复杂的逻辑，比如在函数返回前执行一些特定的操作。
//  7. defer 语句可以在函数中使用闭包来捕获变量的值，从而实现一些特定的功能，比如延迟计算、缓存等。
//  8. defer 语句可以在函数中使用 recover 函数来捕获异常，从而实现一些特定的功能，比如错误处理、日志记录等。
//  9. defer 语句可以在函数中使用 panic 函数来触发异常，从而实现一些特定的功能，比如错误处理、日志记录等。
//  10. defer 语句中的变量，在 defer 声明时就已经确定了其值，而不是在 defer 执行时才确定。

func fn3_1() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Printf("fn3_1 recover err: %v, type: %T \n", err, err) // fn3_1 recover err: fn3_1 panic, type: string
		}
	}()

	panic("fn3_1 panic")
}

func fn3_2(a, b int) int {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Printf("fn3_2 recover err: %v, type: %T \n", err, err) // fn3_2 recover err: runtime error: integer divide by zero, type: runtime.errorString
		}
	}()

	return a / b
}

// 模拟读取文件失败
func fn3_3(fileName string) error {
	if fileName == "main.go" {
		return nil
	}

	return errors.New("fn3_3 read file failed")
}

func fn3_4() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("fn3_4 recover err ==>", err) // fn3_4 recover err ==> fn3_3 read file failed
		}
	}()

	err1 := fn3_3("test.go")

	if err1 == nil {
		fmt.Println("fn3_4 read file success")
	} else {
		fmt.Println("fn3_4 read file failed, err1 ==>", err1) // fn3_4 read file failed, err1 ==> fn3_3 read file failed
		panic(err1)                                           // 把异常信息抛出，交给 defer 中的 recover 来捕获。
	}

	// 因为上边抛出了 panic 所以后边的内容不再执行。

	err2 := fn3_3("main.go")

	if err2 == nil {
		fmt.Println("fn3_4 read file success")
	} else {
		fmt.Println("fn3_4 read file failed, err2 ==>", err2)
		panic(err2)
	}
}

// 打印：九九乘法表
func fn3_5(a int) (bool, error) {
	if a != 9 {
		// 业务级别抛错 (不会阻止程序继续执行) (推荐)
		return false, errors.New("参数只能是 9")

		// 程序级别的异常，抛出错误后，程序不再执行。
		// panic(errors.New("参数只能是 9"))
	}

	for i := 1; i <= a; i++ {
		for k := 1; k <= i; k++ {
			fmt.Printf("%d * %d = %d \t", k, i, k*i)
		}

		fmt.Println()
	}

	return true, nil
}

func main() {
	fn3_1() // recover 捕获之后不会崩溃，否则程序执行到这里 会崩溃。

	fmt.Println("10 / 0 ==>", fn3_2(10, 0)) // 0, 失败后 默认的返回值是 0。

	fn3_4() // 读取文件失败，抛出 panic，recover 捕获之后不会崩溃。

	fmt.Println("----------------分割线---111-----------------")

	status1, err1 := fn3_5(9)

	if err1 != nil {
		fmt.Println("main err1 ==>", err1)
	}

	fmt.Println("status1 ==>", status1) // true

	// 捕获程序级别的错误
	defer func() {
		err := recover()
		fmt.Println("main recover err ==>", err) // main recover str ==> 参数只能是 9
	}()

	status2, err2 := fn3_5(2)

	// 如果是 panic 的错误，则 后边的程序不再执行。
	if err2 != nil {
		fmt.Println("main err2 ==>", err2) // 参数只能是 9
	}

	fmt.Println("status2 ==>", status2) // false

	fmt.Println("----------------分割线---222-----------------")

	fmt.Println("main end...")
}
