package main

import (
	"fmt"
	"reflect"
)

func reflectSetVal2_1(x any) {
	if v := reflect.ValueOf(x); v.Kind() == reflect.Int64 {
		v.SetInt(100) // 基本类型不能这样修改 panic: reflect: reflect.Value.SetInt using unaddressable value
	}
}

func reflectSetVal2_2(x any) {
	// *x = 100 // 基本类型不能这样修改 panic: reflect: reflect.Value.SetInt using unaddressable value

	v1, ok1 := x.(*int64) // 通过类型断言获取指针类型
	if ok1 {
		*v1 = 100 // 通过指针修改值
	}

	v2 := reflect.ValueOf(x)

	fmt.Println("v.Elem().Kind() ==>", v2.Elem().Kind()) // v.Elem().Kind() ==> int64 / string

	if v2.Elem().Kind() == reflect.Int64 { // 传递进来的是指针的话，需要通过 Elem() 来判断。
		v2.Elem().SetInt(100) // 通过反射修改值
	} else if v2.Elem().Kind() == reflect.String {
		v2.Elem().SetString("hello") // 通过反射修改值
	}
}

func main() {
	var a int64 = 10

	// reflectSetVal2_1(a)
	reflectSetVal2_2(&a) // 传入指针类型

	fmt.Println("a ==>", a) // a ==> 100

	fmt.Println("----------------分割线---111-----------------")

	b := "你好"

	reflectSetVal2_2(&b) // 传入指针类型

	fmt.Println("b ==>", b) // b ==> "hello"
}
