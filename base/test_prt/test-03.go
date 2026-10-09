package main

import "fmt"

func main() {
	var age int
	var sex string

	fmt.Print("请输入年龄：")
	fmt.Scan(&age) // fmt.Scan: 从标准输入(stdin) 读取数据，按空白字符分隔，依次填入传入的变量。

	fmt.Print("请输入性别：")
	fmt.Scan(&sex) // 读取用户输入的值，并将其存储在变量 sex 中

	fmt.Println("age ==>", age)
	fmt.Println("sex ==>", sex)

	fmt.Printf("age: %d, type: %T \n", age, age) // age: 33, type: int
	fmt.Printf("sex: %s, type: %T \n", sex, sex) // sex: 男, type: string
}
