package main

import "fmt"

// 结构体方法 和 接受者

// 在 go 语言中，没有类的概念，只有结构体和方法。方法是作用于特定类型的函数。
// 方法与函数的区别在于，方法有一个特殊的接收者参数，它表示方法所属的类型。
// 接受者的概念就类似于其他语言中的 this 或 self，表示当前对象的实例。

// 注意：非本地类型不能定义方法，也就是说，我们不能给别的包的类型定义方法。

// 结构体是 值类型。
// 结构体实例是互相独立的，不会互相影响。

type person3_1 struct {
	name string
	age  int
	sex  string
}

// 结构体添加方法

func (p person3_1) printInfo() {
	fmt.Printf("p: %#v, type: %T \n", p, p)
}

func (p person3_1) setAge() { // p 是 person3_1 的副本，修改 p.age 不会影响原来的结构体实例。
	p.age = 30
}

// 指针类型的接收者方法 (可以修改结构体实例的属性值)

func (p *person3_1) setAge2() { // p 是 person3_1 的指针，修改 p.age 会影响原来的结构体实例。
	p.age = 100
}

// 自定义类型 添加方法

type myInt3_1 int

func (m myInt3_1) printMyInt() {
	fmt.Printf("m: %d, type: %T \n", m, m)
}

func main() {
	// 1. 定义结构体

	p1 := person3_1{
		name: "sky",
		age:  20,
		sex:  "男",
	}

	p2 := p1 // 结构体是值类型，p2 是 p1 的副本，互不影响。

	p2.name = "blue"

	fmt.Printf("p1: %#v \n", p1)               // p1: main.person3_1{name:"sky", age:20, sex:"男"}
	fmt.Printf("p2: %#v, type: %T \n", p2, p2) // p2: main.person3_1{name:"blue", age:20, sex:"男"}, type: main.person3_1

	fmt.Println("----------------分割线---111-----------------")

	// 2. 调用结构体方法

	p2.printInfo() // p: main.person3_1{name:"blue", age:20, sex:"男"}, type: main.person3_1

	p2.setAge()
	fmt.Printf("p2: %#v \n", p2) // p2: main.person3_1{name:"blue", age:20, sex:"男"}

	p2.setAge2() // 不会影响 p1

	fmt.Printf("p2: %#v \n", p2) // p2: main.person3_1{name:"blue", age:100, sex:"男"}
	fmt.Printf("p1: %#v \n", p1) // p1: main.person3_1{name:"sky", age:20, sex:"男"}

	fmt.Println("----------------分割线---222-----------------")

	// 3. 自定义类型添加方法

	var aInt myInt3_1 = 10

	aInt.printMyInt() // m: 10, type: main.myInt3_1
}
