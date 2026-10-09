package main

import (
	"encoding/json"
	"fmt"
)

// 结构体 —— 序列化，反序列化。

type person7_1 struct {
	ID     int    `json:"id"` // 通过指定 tag 实现 json 序列化时的 key 名称
	Gender string // 必须是 大写开头，才能被序列化
	name   string `json:"name"` // 小写开头，不能被序列化 (私有属性)
	Sno    string // 默认序列化后是 "Sno"，如果想要序列化后是 "sno"，可以使用 tag `json:"sno"`
}

func main() {
	// 1. JSON 序列化 结构体

	p1 := person7_1{
		ID:     10,
		Gender: "男",
		name:   "张三",
		Sno:    "sno123",
	}

	fmt.Printf("p1: %#v \n", p1) // p1: main.person7_1{ID:10, Gender:"男", name:"张三", Sno:"sno123"}

	p1JsonByte, err := json.Marshal(p1)

	if err != nil {
		fmt.Println("json.Marshal(p1) err ==>", err)
	}

	p1JsonStr := string(p1JsonByte)

	fmt.Printf("p1JsonStr: %#v \n", p1JsonStr) // p1JsonStr: "{\"id\":10,\"Gender\":\"男\",\"Sno\":\"sno123\"}"
	fmt.Printf("p1JsonStr: %#s \n", p1JsonStr) // p1JsonStr: {"id":10,"Gender":"男","Sno":"sno123"}
	fmt.Println("p1JsonStr ==> ", p1JsonStr)   // p1JsonStr ==>  {"id":10,"Gender":"男","Sno":"sno123"}

	fmt.Println("----------------分割线---111-----------------")

	// 2. JSON 反序列化 结构体

	jsonStr := `{"id":10,"Gender":"男","Sno":"sno123"}`
	var p2 person7_1

	err2 := json.Unmarshal([]byte(jsonStr), &p2)

	if err2 != nil {
		fmt.Println("json.Unmarshal([]byte(jsonStr), &p2) err ==>", err2)
	}

	fmt.Printf("p2: %#v \n", p2) // p2: main.person7_1{ID:10, Gender:"男", name:"", Sno:"sno123"}

	// 注意：name 就算有值在 jsonStr 中，也不会被反序列化到 p3 中，因为 name 是小写开头的私有属性，不能被序列化和反序列化。

	var p3 person7_1
	jsonStr2 := `{"id":20,"Gender":"女","name":"李四","Sno":"sno456"}`
	json.Unmarshal([]byte(jsonStr2), &p3)

	fmt.Printf("p3: %#v \n", p3)               // p3: main.person7_1{ID:20, Gender:"女", name:"", Sno:"sno456"}
	fmt.Printf("p3.name: %#v \n", p3.name)     // p3.name: ""
	fmt.Printf("p3.Gender: %#v \n", p3.Gender) // p3.Gender: "女"
}
