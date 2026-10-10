package main

import "fmt"

// 接口嵌套

// 接口 与 接口 之间可以通过嵌套创造出新的接口。

type animal6_1 interface {
	setName(string)
	run()
}

type animal6_2 interface {
	getName() string
}

// 接口嵌套 创建新的接口

type animal6_3 interface {
	animal6_1
	animal6_2
}

type dog6_1 struct {
	name string
}

func (d *dog6_1) setName(name string) {
	d.name = name
}

func (d dog6_1) getName() string {
	return d.name
}

func (d dog6_1) run() {
	println(d.name + " is running")
}

func main() {
	// 1.

	d1 := dog6_1{
		name: "小黑",
	}

	var a1 animal6_3 = &d1

	fmt.Println("a1.getName() ==>", a1.getName()) // a1.getName() ==> 小黑

	a1.setName("小白")

	fmt.Println("a1.getName() ==>", a1.getName()) // a1.getName() ==> 小白

	a1.run() // 小白 is running

	fmt.Println("----------------分割线---111-----------------")

	// 2.

	d2 := &dog6_1{ // 推荐：直接使用指针类型
		name: "花花",
	}

	var a2 animal6_3 = d2

	a2.run() // 花花 is running

}
