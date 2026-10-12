package main

import (
	"fmt"
	"reflect"
)

// 反射 —— 结构体

// 不要乱用反射
//  1. 反射会破坏类型安全，增加代码复杂度。
//  2. 反射会降低程序性能，增加运行时开销。
//  3. 反射会破坏封装性，增加代码耦合度。
//  4. 反射会增加代码维护成本，降低代码可读性。
//  5. 反射会增加代码调试难度，降低代码可测试性。
//  6. 反射会增加代码安全风险，降低代码安全性。
//  7. 反射会增加代码可移植性问题，降低代码可移植性。
//  8. 反射会增加代码可扩展性问题，降低代码可扩展性。
//  9. 反射会增加代码可重用性问题，降低代码可重用性。
//  10. 反射会增加代码可维护性问题，降低代码可维护性。

// 原因
//  1. 基于反射的代码是及其脆弱的，反射中的类型错误会在真正运行的时候才会引发 panic，那很可能是你在生产环境中才会发现。
//  2. 大量使用反射的代码通常难以理解。

type student3_1 struct {
	Name  string `json:"name" from:"username"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s student3_1) GetInfo3_1() string { // 这里必须用 大写字母开头的方法名
	str := fmt.Sprintf("getInfo3_1 Name: %s, Age: %d, Score: %d", s.Name, s.Age, s.Score)

	fmt.Println("getInfo3_1() called str ==>", str) // getInfo3_1() called str ==> getInfo3_1 Name: 张三, Age: 18, Score: 90

	return str
}

func (s student3_1) SetInfo3_1(name string, age int, score int) {
	fmt.Println("setInfo3_1() called")

	s.Name = name
	s.Age = age
	s.Score = score
}

func (s student3_1) PrintInfo3_1() {
	fmt.Println("printInfo3_1() called")
	fmt.Printf("printInfo3_1 Name: %s, Age: %d, Score: %d\n", s.Name, s.Age, s.Score)
}

func printStructFn3_1(s any) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的不是结构体类型")
		return
	}

	// 1. 通过类型变量里面的 Method 可以获取到结构体里面的方法信息
	method0 := t.Method(0)                                                          // Method 方法的索引是按照方法名的字母顺序来排序的
	fmt.Printf("method0.Name: %v, method0.Type: %v \n", method0.Name, method0.Type) // method0.Name: GetInfo3_1, method0.Type: func(main.student3_1) string

	// 2. 通过类型变量里面的 NumMethod 可以获取到结构体里面的方法数量
	fmt.Println("NumMethod ==>", v.NumMethod()) // 2

	// 3. 通过 <值变量> 执行结构体里面的方法 (主要要使用值变量，并且要使用方法的索引来调用。
	// 注意方法的索引是按照方法名的字母顺序来排序的，注意参数的传递
	v.Method(0).Call(nil) // 调用 GetInfo3_1 方法

	// 4. 通过类型变量获取这个结构体的具体方法
	method1, ok1 := t.MethodByName("PrintInfo3_1")

	if ok1 {
		fmt.Printf("method1.Name: %v, method1.Type: %v \n", method1.Name, method1.Type) // method1.Name: PrintInfo3_1, method1.Type: func(main.student3_1)
	}

	// 5. 执行结构体里面的方法
	v.MethodByName("PrintInfo3_1").Call(nil) // 调用 PrintInfo3_1 方法

	info := v.MethodByName("GetInfo3_1").Call(nil) // 调用 GetInfo3_1 方法 注意函数的返回值是一个 []reflect.Value 类型的切片
	fmt.Println("info ==>", info)                  // info ==> [getInfo3_1 Name: 张三, Age: 18, Score: 90]
	fmt.Printf("info %#v \n", info)                // info []reflect.Value{reflect.Value{typ_:(*abi.Type)(0x1027da9a0), ptr:(unsafe.Pointer)(0x6eb12e116100), flag:0x98}}

	var params []reflect.Value
	params = append(params, reflect.ValueOf("李四"), reflect.ValueOf(20), reflect.ValueOf(95))
	v.MethodByName("SetInfo3_1").Call(params) // 调用 SetInfo3_1 方法 并且传入参数
}

// 打印字段
func printStructField(s any) {
	// 判断是否是结构体
	t := reflect.TypeOf(s)

	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的不是结构体类型, 或者结构体指针。")
		return
	}

	// 1. 通过类型变量里面的 field 可以获取结构体的字段

	field0 := t.Field(0)                                                  // 0 表示获取第零个参数
	fmt.Printf("filed0: %v \n", field0)                                   // filed0: {Name  string json:"name" from:"username" 0 [0] false}
	fmt.Printf("field0: %#v \n", field0)                                  // filed0: reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x100a2c790), Tag:"json:\"name\" from:\"username\"", Offset:0x0, Index:[]int{0}, Anonymous:false}
	fmt.Printf("field0.Name: %v \n", field0.Name)                         // field0.Name: Name
	fmt.Printf("field0.Type: %v \n", field0.Type)                         // field0.Type: string
	fmt.Printf("field0.Tag: %v \n", field0.Tag)                           // field0.Tag: json:"name" from:"username"
	fmt.Printf("field0.Tag.Get(\"json\"): %v \n", field0.Tag.Get("json")) // field0.Tag.Get("json"): name
	fmt.Printf("field0.Tag.Get(\"from\"): %v \n", field0.Tag.Get("from")) // field0.Tag.Get("from"): username, 可以获取 json 的 tag, 也可以获取其他的 tag。

	// 2. 通过类型变量里面的 fieldByName 可以获取结构体的字段

	field1, ok1 := t.FieldByName("Age") // 通过字段名获取字段信息

	if ok1 {
		fmt.Printf("field1: %v \n", field1) // field1: {Age  int json:"age" 16 [1] false}
	}

	// 3. 通过类型变量里面的 NumField 可以获取结构体的字段数量
	var fieldNum = t.NumField()             // 获取结构体的字段数量
	fmt.Printf("fieldNum: %v \n", fieldNum) // fieldNum: 3

	// 4. 通过值变量里面的 Field 可以获取结构体的字段值
	v := reflect.ValueOf(s)
	fieldNameVal := v.FieldByName("Name") // 通过字段名获取字段值

	fmt.Println("fieldNameVal ==>", fieldNameVal) // fieldNameVal ==> 张三

	for i := 0; i < fieldNum; i++ {
		fmt.Printf("属性名称: %v, 属性值: %v, 属性类型: %v, 属性tag: %v \n", t.Field(i).Name, v.Field(i), t.Field(i).Type, t.Field(i).Tag.Get("json"))
		// 属性名称: Name, 属性值: 张三, 属性类型: string, 属性tag: name
		// 属性名称: Age, 属性值: 18, 属性类型: int, 属性tag: age
		// 属性名称: Score, 属性值: 90, 属性类型: int, 属性tag: score
	}
}

// 反射修改结构体属性
func reflectChangeStruct(s any) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Ptr {
		fmt.Println("传入的不是结构体指针类型")
		return
	} else if t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的不是结构体指针类型")
		return
	}

	name := v.Elem().FieldByName("Name") // 获取结构体的字段值
	name.SetString("李四")                 // 修改结构体的字段值

	age := v.Elem().FieldByName("Age") // 获取结构体的字段值
	age.SetInt(30)                     // 修改结构体的字段值
}

func main() {
	stu1 := student3_1{
		Name:  "张三",
		Age:   18,
		Score: 90,
	}

	// printStructFn3_1([]string{"1", "2", "3"}) // 传入的不是结构体类型

	// printStructFn3_1(stu1)

	// printStructFn3_1(&stu1)

	// printStructField(stu1)

	reflectChangeStruct(&stu1)

	fmt.Printf("stu1: %#v \n", stu1) // stu1: main.student3_1{Name:"李四", Age:30, Score:90}
}
