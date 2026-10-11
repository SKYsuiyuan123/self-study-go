package main

import (
	"errors"
	"fmt"
)

// 函数异常处理 —— 错误向上抛

func test1_1(a, b int) (res int, err error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}

	return a / b, nil
}

func test1_2() (res int, err error) {
	res, err = test1_1(10, 0)

	if err != nil {
		return // 这里没有处理错误，而是直接返回错误，向上抛。
	}

	res += 1

	return
}

func main() {
	res1, err1 := test1_2()

	if err1 != nil {
		fmt.Println("test1_2() 出现错误 ==>", err1) // test1_2() 出现错误 ==> 除数不能为0
		return
	}

	fmt.Println("test1_2() 结果:", res1)
}
