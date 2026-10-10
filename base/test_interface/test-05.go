package main

import "fmt"

// 一个结构体 实现 多个接口

// 两个接口

type animal5_1 interface {
	setName(string)
}

type animal5_2 interface {
	getName() string
}

// dog 结构体

type dog5_1 struct {
	name string
}

func (d *dog5_1) setName(name string) {
	d.name = name
}

func (d dog5_1) getName() string {
	return d.name
}

func main() {
	// 1. dog 实现 两个接口

	d1 := dog5_1{
		name: "小黑",
	}

	var d2 animal5_1 = &d1 // dog 实现了 animal5_1 接口
	var d3 animal5_2 = &d1 // dog 实现了 animal5_2 接口

	var d4 animal5_2 = d1 // 注意：这里是值类型，d4 是 dog5_1 的副本，因此 d2.setName("小白") 不会影响 d4 的 name。

	fmt.Println("d3.getName() ==>", d3.getName()) // d3.getName() ==> 小黑

	d2.setName("小白")

	fmt.Println("d3.getName() ==>", d3.getName()) // d3.getName() ==> 小白
	fmt.Println("d1.getName() ==>", d1.getName()) // d1.getName() ==> 小白

	fmt.Println("d4.getName() ==>", d4.getName()) // d4.getName() ==> 小黑

	fmt.Println("----------------分割线---111-----------------")
}
