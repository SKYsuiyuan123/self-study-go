package main

import "fmt"

// Go 指针 *T

// go 函数的变量传递都是 值传递。

// 参数是数组：直接传递的就是 数组的值
// 参数是切片：传递的是切片的值，切片的值包含了指向底层数组的指针，所以函数内部修改切片的元素时，会影响外部变量。
// 参数是 map：map 变量本身就是一个指针(指向底层 hmap 结构)，传参时拷贝的是这个指针，但两个指针指向同一份底层数据。
// 参数是 结构体：直接传递的就是结构体的值，函数内部修改结构体的字段时，不会影响外部变量。(除非传递的是结构体的指针)

// 参数是：普通变量

func fn4_1(x int) { // 值传递，函数内部修改时，不会影响外部变量。
	x = 10
}

func fn4_2(x *int) { // 指针传递，函数内部修改时，会影响外部变量。
	*x = 100 // *x: 解引用，获取指针指向的值。
}

// 参数是：切片

func fn4_3(s []string) { // go 切片传递，切片是引用类型，函数内部修改底层数组的元素时，会影响外部变量。
	s[0] = "hello" // 修改底层数组的元素，外部变量也会受到影响。

	// 不会影响外部变量，因为切片的底层数组可能会重新分配内存。
	s = append(s, "world") // 追加元素，s 的底层数组可能会重新分配内存
}

func fn4_4(s *[]string) { // go 切片引用传递，函数内部修改时，会影响外部变量。

	// 会影响外部变量, *s = ... 直接改了外部的切片头。
	*s = append(*s, "world") // 追加元素，s 的底层数组可能会重新分配内存

	// 常见做法是 直接返回新切片。
	// return append(s, "world")
}

// 参数是：数组

func fn4_5(s [3]string) { // go 数组传递，函数内部修改时，不会影响外部变量。
	s[0] = "hello" // 修改底层数组的元素，外部变量不会受到影响。
}

// 参数是：map

func fn4_6(m map[string]int) { // go map 传递，map 是引用类型，函数内部修改时，会影响外部变量。
	if m == nil { // 对 nil map 进行写操作会 panic，所以要先判断是否为 nil。
		return
	}

	m["语文"] = 0

	delete(m, "数学") // 删除 map 中的元素，外部变量也会受到影响。
}

// 参数是：结构体

type person4_1 struct {
	name string
	age  int
	sex  bool
}

func fn4_7(p person4_1) { // go 结构体传递，函数内部修改时，不会影响外部变量。
	p.age = 500
}

func fn4_8(p *person4_1) { // go 结构体指针传递，函数内部修改时，会影响外部变量。
	p.age = 300 // 修改结构体的字段，外部变量也会受到影响。

	// 全称
	(*p).sex = false
}

func main() {
	x1 := 1
	fn4_1(x1) // 值传递

	fmt.Println("x ==>", x1)    // 1
	fmt.Println("&x1 ==>", &x1) // 0x66a0e2b8e020

	fn4_2(&x1)               // 指针传递，传入 x1 的地址
	fmt.Println("x ==>", x1) // 100

	fmt.Println("----------------分割线---111-----------------")

	s1 := []string{"a", "b", "c"}

	fn4_3(s1)                    // 切片传递
	fmt.Printf("s1: %#v \n", s1) // s1: []string{"hello", "b", "c"}

	s2 := make([]string, 0, 3)
	s2 = append(s2, "A")

	fn4_3(s2)
	fmt.Printf("s2: %#v, cap(s2): %v \n", s2, cap(s2)) // s2: []string{"hello"}, cap(s2): 3

	fn4_4(&s2)
	fmt.Printf("s2: %#v, cap(s2): %v \n", s2, cap(s2)) // s2: []string{"hello", "world"}, cap(s2): 3

	fn4_4(&s1)
	fmt.Printf("s1: %#v, cap(s1): %v \n", s1, cap(s1)) // s1: []string{"hello", "b", "c", "world"}, cap(s1): 6

	fmt.Println("----------------分割线---222-----------------")

	arr1 := [3]string{"a", "b", "c"}

	fn4_5(arr1)
	fmt.Printf("arr1: %#v \n", arr1) // arr1: [3]string{"a", "b", "c"}

	fmt.Println("----------------分割线---333-----------------")

	m1 := map[string]int{"语文": 100, "数学": 90}

	fn4_6(m1)
	fmt.Printf("m1: %#v \n", m1) // m1: map[string]int{"语文":0}

	fmt.Println("----------------分割线---444-----------------")

	p1 := person4_1{
		name: "sky",
		age:  18,
		sex:  true,
	}

	fn4_7(p1)                    // 结构体传递
	fmt.Printf("p1: %#v \n", p1) // p1: main.person4_1{name:"sky", age:18, sex:true}

	fn4_8(&p1)                   // 结构体指针传递
	fmt.Printf("p1: %#v \n", p1) // p1: main.person4_1{name:"sky", age:300, sex:false}
}
