package main

import "fmt"

// 空接口

// go 中的接口可以不定义任何方法，没有定义任何方法的接口就是空接口。
// 空接口表示没有任何约束，因此任何类型的变量都可以实现空接口。
// 空接口在实际项目中用的是非常多的，用空接口可以表示任意数据类型。

type aInterface interface{}

// 空接口可以作为函数参数

func show8_1(a interface{}, b any) { // any 类型背后就是 interface{}，any 是 interface{} 的别名
	fmt.Printf("show8_1: a: %v, b: %v \n", a, b)
}

// 类型断言

// 如果要判断空接口中的值的类型，可以使用类型断言。
// 语法格式: x.(T), 其中 x 是一个接口类型的变量，T 是一个类型。x: 表示要断言的接口变量，T: 表示要断言的类型。
// 该语法返回两个值，第一个值是接口变量中保存的具体类型的值，第二个值是一个布尔值，表示断言是否成功。
// 该语法 断言成功时，返回接口变量中保存的具体类型的值，并且第二个值为 true；断言失败时，返回该类型的零值，并且第二个值为 false。
// 如果只接收一个返回值是不安全的，一旦断言失败，程序就会 panic。

// 注意: 变量.(type) 只能在类型 switch 中使用，不能在普通的 switch 中 和 其它语句中使用。

func justifyType8_1(x any) {
	switch v := x.(type) {
	case int:
		fmt.Printf("x is int, value: %v \n", v)
	case string:
		fmt.Printf("x is string, value: %s \n", v)
	case bool:
		fmt.Printf("x is bool, value: %v \n", v)
	default:
		fmt.Println("x is unknown type")
	}
}

func main() {
	// 1.

	var a1 aInterface

	var str = "你好 golang"
	a1 = str // 空接口可以保存任意类型的值

	fmt.Printf("a1 值: %v, 类型: %T \n", a1, a1) // a1 值: 你好 golang, 类型: string

	var num1 = 20
	a1 = num1

	fmt.Printf("a1 值: %v, 类型: %T \n", a1, a1) // a1 值: 20, 类型: int

	var flag = true
	a1 = flag

	fmt.Printf("a1 值: %v, 类型: %T \n", a1, a1) // a1 值: true, 类型: bool

	fmt.Println("----------------分割线---111-----------------")

	// 2. 空接口也可以直接当做类型来使用，可以表示任意类型。

	var a2 interface{} // 空接口类型的变量 a2

	a2 = 20
	a2 = "hello golang"
	a2 = true
	a2 = map[string]string{"a": "A"}

	fmt.Printf("a2 值: %#v, 类型: %T \n", a2, a2) // a2 值: map[string]string{"a":"A"}, 类型: map[string]string

	show8_1(123, "hello golang") // show8_1: a: 123, b: hello golang

	fmt.Println("----------------分割线---222-----------------")

	// 3. map 和 slice 的元素类型为 any 类型

	// 3.1 map 类型的值为 any 类型
	var m1 = map[string]any{
		"a": 123,
		"b": "hello golang",
	}
	fmt.Printf("m1 值: %#v, 类型: %T \n", m1, m1) // m1 值: map[string]interface {}{"a":123, "b":"hello golang"}, 类型: map[string]interface {}

	// 3.2 切片的元素类型为 any 类型
	var slice1 = []any{"a", 1, true}
	fmt.Printf("slice1 值: %#v, 类型: %T \n", slice1, slice1) // slice1 值: []interface {}{"a", 1, true}, 类型: []interface {}

	fmt.Println("----------------分割线---333-----------------")

	// 4. 类型断言

	aIsInt, ok1 := a1.(int)                           // 断言 a1 的类型为 int
	fmt.Printf("aIsInt: %v, ok1: %v \n", aIsInt, ok1) // aIsInt: 0, ok1: false

	aIsBool, ok2 := a1.(bool)                           // 断言 a1 的类型为 bool
	fmt.Printf("aIsBool: %v, ok2: %v \n", aIsBool, ok2) // aIsBool: true, ok1: true

	justifyType8_1(123) // x is int, value: 123
	justifyType8_1(a1)  // x is bool, value: true
	justifyType8_1(a2)  // x is unknown type

	fmt.Println("----------------分割线---444-----------------")

	aIsInt2 := a1.(int)                   // 断言 a1 的类型为 int, 断言失败会 panic (panic: interface conversion: main.aInterface is bool, not int)
	fmt.Printf("aIsInt2: %v \n", aIsInt2) // aIsInt2: 0
}
