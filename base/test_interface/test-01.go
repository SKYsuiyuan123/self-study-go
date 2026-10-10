package main

import "fmt"

// interface 接口

// 接口是一种类型，一种抽象的类型。接口是一组方法的集合，接口中不能包含任何的变量。接口类型的变量可以存储实现了该接口的任意类型的值。
// 接口只能定义方法的集合。

// 多态 —— 多态是指同一个接口可以指向不同的类型的值，从而实现不同的行为。
// 	1. 有一个父类 (有接口)
//  2. 有子类 (实现了父类的全部接口方法)
//  3. 父类的引用指向子类的对象 (接口类型的变量可以存储实现了该接口的任意类型的值) (3. 父类类型的变量(指针) 指向(引用) 子类的具体数据变量)

// 接口是一个规范

type usb1_1er interface { // 名字后边加er 是一个习惯，表示这个接口是一个行为规范
	start()
	stop()
}

// 如果接口里边有方法的话，那么实现这个接口的类型必须实现接口里的所有方法，否则就会报错。
// 必须要通过结构体 或者 通过自定义类型来实现接口里的方法，才能实现接口的规范。

type phone1_1 struct {
	name string
}

// phone 要实现 usb 接口的话，必须得实现 usb 接口里的所有方法

func (p phone1_1) start() {
	fmt.Println(p.name, "phone running ...")
}

func (p phone1_1) stop() {
	fmt.Println(p.name, "phone stop ...")
}

type camera1_1 struct{}

func (c camera1_1) start() {
	fmt.Println("camera start ...")
}

func (c camera1_1) stop() {
	fmt.Println("camera stop ...")
}

func (c camera1_1) run() {
	fmt.Println("camera run ...")
}

func main() {
	// 1. 定义 phone 类型的变量

	p1 := phone1_1{
		name: "iphone",
	}

	p1.start() // iphone phone running ...

	var usb1 usb1_1er // 定义一个接口类型的变量 usb1

	usb1 = p1 // 让 usb1 接口类型的变量指向 phone 类型的变量 p1 (表示手机实现了 usb 接口的规范)

	usb1.stop() // iphone phone stop ...

	fmt.Println("p1.name ==>", p1.name) // p1.name ==> iphone
	fmt.Printf("usb1: %#v \n", usb1)    // usb1: main.phone1_1{name:"iphone"}
	// fmt.Println("usb1.name ==>", usb1.name) // usb1.name 无法访问，因为 usb1 类型上没有 name 这个属性。报错：usb1.name undefined (type usb1_1er has no field or method name)

	fmt.Println("----------------分割线---111-----------------")

	// 2. 定义 camera 类型的变量

	c1 := camera1_1{}

	var c2 usb1_1er = c1

	c2.start() // camera start ...
	// c2.run()   // 报错：c2.run undefined (type usb1_1er has no field or method run)  因为 c2 是 usb1_1er 类型的变量，而 usb1_1er 接口里边没有 run 方法，所以无法调用 run 方法。

	c1.run() // camera run ... 只能通过 c1 来调用 run 方法，因为 c1 是 camera1_1 类型的变量，而 camera1_1 类型里边有 run 方法。

	fmt.Println("----------------分割线---222-----------------")

	// 3. 定义一个接口类型的变量 usb1_1er，并让它指向 phone1_1 类型的变量

	var d1 usb1_1er
	d1 = &phone1_1{
		name: "xiaomi",
	}

	d1.start() // xiaomi phone running ...
}
