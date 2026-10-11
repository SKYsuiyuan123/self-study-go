package main

import (
	"encoding/json"
	"fmt"
)

// 泛型 —— 结构体

type response2_1[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type user2_1 struct {
	Name string `json:"name"`
}

type userInfo2_1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 1. 使用泛型结构体

	res1 := response2_1[user2_1]{
		Code:    0,
		Message: "success",
		Data: user2_1{
			Name: "sky",
		},
	}

	fmt.Printf("res1: %#v \n", res1)                  // res1: main.response2_1[main.user2_1]{Code:0, Message:"success", Data:main.user2_1{Name:"sky"}}
	fmt.Println("res1.Data.Name ==>", res1.Data.Name) // res1.Data.Name ==> sky

	fmt.Println("----------------分割线---111-----------------")

	// 2. 使用泛型结构体 + json 序列化

	user1 := userInfo2_1{
		Name: "sky2",
		Age:  31,
	}

	res2 := response2_1[userInfo2_1]{
		Code:    0,
		Message: "success",
		Data:    user1,
	}

	fmt.Printf("res2: %#v \n", res2) // res2: main.response2_1[main.userInfo2_1]{Code:0, Message:"success", Data:main.userInfo2_1{Name:"sky2", Age:31}}

	res2Byte, _ := json.Marshal(res2)
	res2JsonStr := string(res2Byte)

	fmt.Println("res2JsonStr ==>", res2JsonStr) // {"code":0,"message":"success","data":{"name":"sky2","age":31}}

	fmt.Println("----------------分割线---222-----------------")

	// 3. 反序列化 json 字符串为 泛型结构体

	// 反序列时 使用 泛型

	var res3 response2_1[userInfo2_1]

	json.Unmarshal([]byte(res2JsonStr), &res3)

	fmt.Printf("res3: %#v \n", res3)                  // res3: main.response2_1[main.userInfo2_1]{Code:0, Message:"success", Data:main.userInfo2_1{Name:"sky2", Age:31}}
	fmt.Println("res3.Data.Name ==>", res3.Data.Name) // res3.Data.Name ==> sky2
}
