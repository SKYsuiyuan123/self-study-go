package main

import (
	"fmt"
	"reflect"
)

// 反射

// 反射可以实现的功能
//  1. 获取变量的类型、变量的值
//  2. 修改变量的值 (前提是变量是可设置的)（注意：反射只能修改可导出的字段，不能修改不可导出的字段）(需要使用 reflect.Value.Elem() 方法获取指针指向的值，然后再使用 reflect.Value.Set() 方法修改值)
//  3. 获取结构体的 字段信息、方法信息。
//  4. 动态调用结构体的方法
//  5. 动态创建结构体对象
//  6. 动态创建切片、数组、map等数据结构
//  7. 动态调用函数
//  8. 动态调用接口方法
//  9. 动态调用函数的参数和返回值、参数和返回值类型、参数和返回值数量、参数和返回值顺序、参数和返回值类型转换、参数和返回值类型断言、参数和返回值类型判断

// go 反射可以实现的功能
// 	1. 反射可以在程序运行期间动态的获取变量的各种信息，比如变量的类型、类别。
//  2. 如果是结果体，通过反射还可以获取结构体本身的信息，比如结构体的字段、结构体的方法。
//  3. 通过反射，可以修改变量的值，可以调用关联的方法。

// 反射的典型应用场景
// 			场景										例子
// 序列化/反序列化				encoding/json、yaml、xml、protobuf
// ORM									gorm 根据结构体字段生成 SQL
// 配置解析								viper 把配置映射到结构体
// 参数校验								validator 读 tag 校验字段
// 依赖注入								wire、dig 按类型注入
// RPC									动态调用方法
// 模板									 text/template、html/template
// 测试断言								testify 比较结构体
// 深拷贝								 copier、deepcopy

// Go 1.18 后，能用泛型解决的优先用泛型，反射留给“编译期类型未知”的场景。

type myInt1_1 int
type person1_1 struct {
	Name string
	Age  int
}

func reflectFn1_1(x any) {
	// 使用反射获取任意变量的类型
	typeVal := reflect.TypeOf(x)

	name := typeVal.Name() // 获取类型的名称 (数组、切片、Map、指针、函数等类型没有名称)
	kind := typeVal.Kind() // 获取类型的种类 (底层类型)

	fmt.Printf("类型: %v, 类型名称: %#v, 类型种类: %v \n", typeVal, name, kind)
}

func reflectFn1_2(x any) {
	res := 0
	b, ok := x.(int) // 类型断言，判断 x 是否是 int 类型

	if ok {
		res += b
	}

	fmt.Println("reflectFn1_2 res ==>", res)
}

func reflectFn1_3(x any) {
	v := reflect.ValueOf(x)

	kind := v.Kind() // 获取类型的种类 (底层类型)

	switch kind {
	case reflect.Int:
		// 反射获取变量的原始值
		m := 10 + v.Int() // 获取 int 类型的值
		fmt.Println("reflectFn1_3 m ==>", m)

	case reflect.Float64:
		m := v.Float() + 10.2
		fmt.Println("reflectFn1_3 m ==>", m)

	case reflect.String:
		m := v.String() + " world"
		fmt.Println("reflectFn1_3 m ==>", m)

	default:
		fmt.Println("reflectFn1_3 m ==>", "x 不是 int 或 float64 类型")
	}
}

func main() {
	a := 10
	b := 23.4
	c := true
	d := "hello"

	reflectFn1_1(a) // 类型: int, 类型名称: "int", 类型种类: int
	reflectFn1_1(b) // 类型: float64, 类型名称: "float64", 类型种类: float64
	reflectFn1_1(c) // 类型: bool, 类型名称: "bool", 类型种类: bool
	reflectFn1_1(d) // 类型: string, 类型名称: "string", 类型种类: string

	var a2 myInt1_1 = 34
	var f1 = person1_1{
		Name: "张三",
		Age:  23,
	}

	reflectFn1_1(a2) // 类型: main.myInt1_1, 类型名称: "myInt1_1", 类型种类: int
	reflectFn1_1(f1) // 类型: main.person1_1, 类型名称: "person1_1", 类型种类: struct

	a3 := new(12)
	reflectFn1_1(a3) // 类型: *int, 类型名称: "", 类型种类: ptr

	arr1 := [3]int{1, 2, 3}
	slice1 := []int{1, 2, 3}
	m1 := map[string]string{"a": "aaa"}

	reflectFn1_1(arr1)   // 类型: [3]int, 类型名称: "", 类型种类: array
	reflectFn1_1(slice1) // 类型: []int, 类型名称: "", 类型种类: slice
	reflectFn1_1(m1)     // 类型: map[string]string, 类型名称: "", 类型种类: map

	fmt.Println("----------------分割线---111-----------------")

	reflectFn1_2(10) // reflectFn1_2 res ==> 20
	reflectFn1_2(m1) // reflectFn1_2 res ==> 0, m1 不是 int 类型。

	fmt.Println("----------------分割线---222-----------------")

	reflectFn1_3(10)  // reflectFn1_3 m ==> 20
	reflectFn1_3(2.2) // reflectFn1_3 m ==> 12.399999999999999
	reflectFn1_3(m1)  // reflectFn1_3 m ==> x 不是 int 或 float64 类型
}
