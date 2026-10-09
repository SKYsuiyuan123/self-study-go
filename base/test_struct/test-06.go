package main

import "fmt"

// 结构体 —— 重写 String 方法

type person6_1 struct {
	name string
	age  int
}

func (p person6_1) String() string {
	return fmt.Sprintf("name: %s, age: %d \n", p.name, p.age)
}

func main() {
	// 1.

	p1 := person6_1{
		name: "sky",
		age:  20,
	}

	fmt.Println("p1 ==>", p1) // 调用了 person6_1.String 方法，打印：p1 ==> name: sky, age: 20
}
