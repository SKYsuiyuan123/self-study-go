package main

import "fmt"

// 接口 —— 方法接收参数 和 返回值

type animal4_1 interface {
	setName(string)
	getName() string
}

type dog4_1 struct {
	name string
}

func (d *dog4_1) setName(name string) { // 结构体是 值类型，想要修改 d.name，必须使用 指针接收者。
	d.name = name
}

func (d dog4_1) getName() string { // 值接收者，不能修改 结构体字段。
	return d.name
}

type cat4_1 struct {
	name string
}

func (c *cat4_1) setName(name string) { // 指针接收者，可以修改 结构体字段。
	c.name = name
}

func (c *cat4_1) getName() string { // 不修改 结构体字段，使用 指针类型 或者 值类型都可以。
	return c.name
}

func main() {
	// 1. 指针接收者实现接口

	d1 := &dog4_1{
		name: "旺财",
	}

	var a1 animal4_1 = d1 // 结构体指针类型 可以赋值给接口变量

	fmt.Println("a1.getName() ==>", a1.getName()) // 旺财

	a1.setName("小黑") // 修改 结构体字段

	fmt.Println("a1.getName() ==>", a1.getName()) // 小黑

	fmt.Println("----------------分割线---111-----------------")

	// 2.

	c2 := &cat4_1{
		name: "咪咪",
	}

	var a2 animal4_1 = c2

	fmt.Println("a2.getName() ==>", a2.getName()) // 咪咪

	a2.setName("小花") // 修改 结构体字段

	fmt.Println("a2.getName() ==>", a2.getName()) // 小花
}
