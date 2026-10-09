package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

// 嵌套结构体 —— JSON 序列化、反序列化

type person9_1 struct {
	ID     int
	Gender string
	Name   string
}

type class9_1 struct {
	Title    string
	Students []person9_1 // person 切片
}

func main() {
	// 1. 创建一个班级对象

	c1 := class9_1{
		Title:    "Go 语言",
		Students: []person9_1{},

		// 或
		// Students: make([]person9_1, 0, 5), // 预分配容量为 5 的切片
	}

	for i := 0; i < 9; i++ {
		p1 := person9_1{
			ID:     i + 1,
			Name:   fmt.Sprintf("stu-%02d", i+1),
			Gender: "男",
		}

		c1.Students = append(c1.Students, p1)
	}

	fmt.Println("c1.Students cap ==>", cap(c1.Students))              // 17
	fmt.Println("c1.Students len ==>", len(c1.Students))              // 9
	fmt.Println("c1.Students Sizeof ==>", unsafe.Sizeof(c1.Students)) // 24

	fmt.Printf("c1: %#v \n", c1) // c1: main.class9_1{Title:"Go 语言", Students:[]main.person9_1{main.person9_1{ID:1, Gender:"男", Name:"stu-01"}, main.person9_1{ID:2, Gender:"男", Name:"stu-02"}, main.person9_1{ID:3, Gender:"男", Name:"stu-03"}, main.person9_1{ID:4, Gender:"男", Name:"stu-04"}, main.person9_1{ID:5, Gender:"男", Name:"stu-05"}}}

	fmt.Println("----------------分割线---111-----------------")

	// 2. JSON 序列化 --> JSON 格式的字符串

	// 2.1 json 化 为 JSON 字符串
	c1Byte, err := json.Marshal(c1)

	if err != nil {
		fmt.Println("json.Marshal err:", err)
	}

	c1JsonStr := string(c1Byte)

	fmt.Printf("c1JsonStr: %#v \n", c1JsonStr) // c1JsonStr: "{\"Title\":\"Go 语言\",\"Students\":[{\"ID\":1,\"Gender\":\"男\",\"Name\":\"stu-01\"},{\"ID\":2,\"Gender\":\"男\",\"Name\":\"stu-02\"},{\"ID\":3,\"Gender\":\"男\",\"Name\":\"stu-03\"},{\"ID\":4,\"Gender\":\"男\",\"Name\":\"stu-04\"},{\"ID\":5,\"Gender\":\"男\",\"Name\":\"stu-05\"}]}"
	fmt.Printf("c1JsonStr: %#s \n", c1JsonStr) // c1JsonStr: {"Title":"Go 语言","Students":[{"ID":1,"Gender":"男","Name":"stu-01"},{"ID":2,"Gender":"男","Name":"stu-02"},{"ID":3,"Gender":"男","Name":"stu-03"},{"ID":4,"Gender":"男","Name":"stu-04"},{"ID":5,"Gender":"男","Name":"stu-05"}]}

	// 2.2 反序列化为结构体对象
	var c2 class9_1
	err2 := json.Unmarshal([]byte(c1JsonStr), &c2)

	if err2 != nil {
		fmt.Println("json.Unmarshal err:", err2)
	}

	fmt.Printf("c2: %#v \n", c2) // c2: main.class9_1{Title:"Go 语言", Students:[]main.person9_1{main.person9_1{ID:1, Gender:"男", Name:"stu-01"}, main.person9_1{ID:2, Gender:"男", Name:"stu-02"}, main.person9_1{ID:3, Gender:"男", Name:"stu-03"}, main.person9_1{ID:4, Gender:"男", Name:"stu-04"}, main.person9_1{ID:5, Gender:"男", Name:"stu-05"}}}

	var c3 = &class9_1{}
	json.Unmarshal([]byte(c1JsonStr), c3)

	fmt.Printf("c3: %#v \n", c3) // c3: &main.class9_1{Title:"Go 语言", Students:[]main.person9_1{main.person9_1{ID:1, Gender:"男", Name:"stu-01"}, main.person9_1{ID:2, Gender:"男", Name:"stu-02"}, main.person9_1{ID:3, Gender:"男", Name:"stu-03"}, main.person9_1{ID:4, Gender:"男", Name:"stu-04"}, main.person9_1{ID:5, Gender:"男", Name:"stu-05"}}}
}
