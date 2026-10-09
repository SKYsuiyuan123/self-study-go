package main

import "fmt"

// 结构体 —— 实例化

// 使用 type 和 struct 关键字来自定义结构体。

type person1_1 struct {
	name string
	age  int
	sex  string
}

func main() {
	// 1. 方法一, 只声明不初始化，属性都是默认零值。

	var p1 person1_1 // 属性没有初始化，默认值都是 默认零值。

	fmt.Printf("p1: %#v \n", p1) // p1: main.person1_1{name:"", age:0, sex:""}

	p1.name = "sky1"
	p1.age = 30
	p1.sex = "男"

	fmt.Printf("p1: %#v, type: %T \n", p1, p1) // p1: main.person1_1{name:"sky1", age:30, sex:"男"}, type: main.person1_1

	fmt.Println("----------------分割线---111-----------------")

	// 2. 方法二， 直接初始化结构体的属性值。(空属性值会被赋予默认零值)

	var p2 person1_1 = person1_1{} // 提示：先声明结构体、再单独给嵌入字段赋值，这种写法可以合并到结构体字面量里。(embedded field assignment can be moved to struct literal)

	p2.name = "sky2"

	fmt.Printf("p2: %#v, type: %T \n", p2, p2) // p2: main.person1_1{name:"sky2", age:0, sex:""}, type: main.person1_1

	fmt.Println("----------------分割线---222-----------------")

	// 3. 方法三，直接初始化结构体的属性值，并且指定属性名。(未初始化的属性会被赋予默认零值)

	var p3 = person1_1{
		name: "sky3",
	}

	fmt.Printf("p3: %#v type: %T \n", p3, p3) // p3: main.person1_1{name:"sky3", age:0, sex:""} type: main.person1_1

	fmt.Println("----------------分割线---333-----------------")

	// 4. 使用 new 实例化结构体

	var p4 = new(person1_1) // 返回的是结构体的指针类型 *person1_1

	p4.name = "blue1" // 通过指针访问结构体的字段时，Go 会自动解引用，所以可以直接使用 p4.name 来访问。
	p4.age = 25

	fmt.Printf("p4: %#v, type: %T \n", p4, p4) // p4: &main.person1_1{name:"blue1", age:25, sex:""}, type: *main.person1_1

	var p4_2 = new(person1_1{
		name: "blue2",
		age:  28,
	})

	fmt.Printf("p4_2: %#v, type: %T \n", p4_2, p4_2) // p4_2: &main.person1_1{name:"blue2", age:28, sex:""}, type: *main.person1_1

	fmt.Println("----------------分割线---444-----------------")

	// 5. 方法五，使用 & 取地址符号来实例化结构体。

	var p5 = &person1_1{} // 返回的是结构体的指针类型 *person1_1

	p5.name = "green1"

	fmt.Printf("p5: %#v, type: %T \n", p5, p5) // p5: &main.person1_1{name:"green1", age:0, sex:""}, type: *main.person1_1

	var p5_2 = &person1_1{
		name: "green2",
		age:  35,
	}

	fmt.Printf("p5_2: %#v, type: %T \n", p5_2, p5_2) // p5_2: &main.person1_1{name:"green2", age:35, sex:""}, type: *main.person1_1

	fmt.Println("----------------分割线---555-----------------")

	// 6. 方法六，初始化时省略属性名。

	var p6 = person1_1{
		"sky6", // 省略属性名，必须按照结构体定义的顺序来赋值。
		40,
		"男", // 不能缺少任何一个 value，否则报错。
	}

	fmt.Printf("p6: %#v, type: %T \n", p6, p6) // p6: main.person1_1{name:"sky6", age:40, sex:"男"}, type: main.person1_1
}
