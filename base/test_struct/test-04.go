package main

import "fmt"

// 父子结构体

// 一个结构体中可以嵌套包含另一个结构体或结构体指针

// 父结构体
type animal4_1 struct {
	name string
}

// 结构体被嵌套后，结构体的方法也会被继承。(子结构体的实例可以直接调用父结构体的方法)
func (a animal4_1) run() {
	fmt.Printf("a.name: %s, run... \n", a.name)
}

// 一个结构体中嵌套另一个结构体

// 子结构体
type dog4_1 struct {
	age       string
	animal4_1 // 结构体嵌套 —— > 继承
}

// 结构体同名方法遮蔽：dog4_1 的实例在调用 run 方法时，优先调用 dog4_1 的 run 方法，而不是 animal4_1 的 run 方法。

func (d dog4_1) run() {
	fmt.Printf("a.name: %s, dog run... \n", d.name)
}

func (d dog4_1) wang() {
	fmt.Printf("a.name: %s, dog wang... \n", d.name)
}

// 一个结构体中嵌套另一个结构体指针

type animal4_2 struct {
	name string
}

type dog4_2 struct {
	age        int
	*animal4_2 // 结构体指针嵌套 —— > 继承
}

func main() {
	// 1. 创建子结构体的实例

	d1 := dog4_1{
		age: "3",
		animal4_1: animal4_1{
			name: "小黑",
		},
	}

	fmt.Printf("d1: %#v \n", d1) // d1: main.dog4_1{age:"3", animal4_1:main.animal4_1{name:"小黑"}}

	d1.run()  // a.name: 小黑, dog run...
	d1.wang() // a.name: 小黑, dog wang...

	fmt.Println("----------------分割线---111-----------------")

	// 2. 嵌套结构体 —— 指针

	d2 := dog4_2{
		age: 5,
		animal4_2: &animal4_2{
			name: "小白",
		},
	}

	fmt.Printf("d2: %#v \n", d2)                               // d2: main.dog4_2{age:5, animal4_2:(*main.animal4_2)(0xebacf49c040)}
	fmt.Printf("d2.animal4_2: %#v \n", d2.animal4_2)           // d2.animal4_2: &main.animal4_2{name:"小白"}
	fmt.Printf("d2.animal4_2.name: %#v \n", d2.animal4_2.name) // d2.animal4_2.name: "小白"
}
