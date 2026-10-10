package main

import "fmt"

// 结构体 值接收者 和 指针接收者 实现接口的区别

// 值接收者：
// 如果结构体中的方法是值接收者，那么该结构体的实例可以直接调用该方法，而不需要使用指针。值接收者会创建一个结构体的副本，因此在方法中对结构体字段的修改不会影响原始结构体。
// 如果结构体中的方法是值接收者，那么实例化后的 结构体值类型 和 结构体指针类型 都可以赋值给接口变量。

// 指针接收者：
// 如果结构体中的方法是指针接收者，那么该结构体的实例必须使用指针来调用该方法。指针接收者允许在方法中修改结构体字段，因为它操作的是原始结构体的地址。
// 如果结构体中的方法是指针接收者，那么实例化后 结构体指针类型 可以赋值给接口变量。结构体值类型 不能赋值给接口变量。

type usb3_1er interface {
	start()
	stop()
}

type phone3_1 struct {
	name string
}

func (p phone3_1) start() { // 值接收者
	fmt.Println(p.name, "手机开始工作")
}

func (p phone3_1) stop() { // 值接收者
	fmt.Println(p.name, "手机停止工作")
}

type phone3_2 struct {
	name string
}

func (p *phone3_2) start() { // 指针接收者
	fmt.Println(p.name, "iphone running...")
}

func (p *phone3_2) stop() { // 指针接收者
	fmt.Println(p.name, "iphone stop...")
}

func main() {
	// 1. 值接收者实现接口

	p1 := phone3_1{
		name: "小米",
	}

	var u1 usb3_1er = p1 // 结构体值类型 可以赋值给接口变量
	u1.start()           // 小米 手机开始工作

	p2 := &phone3_1{
		name: "华为",
	}

	var u2 usb3_1er = p2 // 结构体指针类型 可以赋值给接口变量
	u2.stop()            // 华为 手机停止工作

	fmt.Println("----------------分割线---111-----------------")

	// 2. 指针接收者实现接口

	p3 := phone3_2{
		name: "iphone",
	}

	var u3 usb3_1er = &p3 // 结构体指针类型 可以赋值给接口变量
	u3.start()            // iphone running...

	p4 := &phone3_2{
		name: "步步高",
	}

	var u4 usb3_1er = p4 // 结构体指针类型 可以赋值给接口变量
	u4.start()           // 步步高 running...

	fmt.Println("----------------分割线---222-----------------")

	// 3. 指针接收者实现接口，结构体值类型不能赋值给接口变量

	// p6 := phone3_2{
	// 	name: "Oppo 手机",
	// }

	// var u5 usb3_1er = p6 // 结构体值类型 不能赋值给接口变量
	// u5.start()           // 编译报错：cannot use p6 (variable of type phone3_2) as usb3_1er value in variable declaration: phone3_2 does not implement usb3_1er (start method has pointer receiver)
}
