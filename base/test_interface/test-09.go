package main

import "fmt"

// 空接口 和 类型断言 使用细节

type address9_1 struct {
	name  string
	phone int
}

func main() {
	// 1. 断言切片类型

	userInfo1 := make(map[string]interface{})

	userInfo1["username"] = "张三"
	userInfo1["age"] = 18
	userInfo1["hobby"] = []string{"篮球", "足球", "乒乓球"}

	fmt.Printf("userInfo1: %#v \n", userInfo1) // userInfo1: map[string]interface {}{"age":18, "hobby":[]string{"篮球", "足球", "乒乓球"}, "username":"张三"}
	// fmt.Println("userInfo1[hobby][0] ==>", userInfo1["hobby"][0]) // 不能直接访问，因为 userInfo1["hobby"] 的类型是 interface{}，需要进行类型断言。

	hobby1, ok1 := userInfo1["hobby"].([]string) // 类型断言 获取切片的类型

	if ok1 {
		fmt.Printf("hobby1: %#v \n", hobby1) // hobby1: []string{"篮球", "足球", "乒乓球"}
	}

	fmt.Println("----------------分割线---111-----------------")

	address1 := address9_1{
		name:  "李四",
		phone: 123456,
	}

	userInfo1["address"] = address1

	fmt.Printf("userInfo1: %#v \n", userInfo1)                     // userInfo1: map[string]interface {}{"address":main.address9_1{name:"李四", phone:123456}, "age":18, "hobby":[]string{"篮球", "足球", "乒乓球"}, "username":"张三"}
	fmt.Printf("userInfo1[address]: %#v \n", userInfo1["address"]) // userInfo1[address]: main.address9_1{name:"李四", phone:123456}

	// 可以通过断言来访问

	address2, address2Ok := userInfo1["address"].(address9_1)

	if address2Ok {
		fmt.Printf("address2: %#v \n", address2) // address2: main.address9_1{name:"李四", phone:123456}
	}
}
