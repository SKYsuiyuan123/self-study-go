package main

import (
	"encoding/json"
	"fmt"
)

// 结构体 —— tag

// 字段全部小写，私有，JSON 无法读取。
type person8_1 struct {
	name     string
	age      int
	password string
}

// 首字母大写导出，JSON tag 指定序列化后的 key。

type person8_2 struct {
	Name       string `json:"name"`
	Age        int    // json tag 可以忽略，默认就是字段名 "Age"
	password   string
	Password1  string // 空值也会被序列化
	Password2  string `json:"-"`                    // tag 可以强制某个字段 不参与序列化。
	Password3  string `json:"_"`                    // 序列化出来后 key 是 "_"，但是一般不这么用。
	DefaultStr string `json:"defaultStr,omitempty"` // omitempty 可以让字段在序列化时，如果 值是默认零值，就不参与序列化。
	Age2       *int   `json:"age2,omitempty"`       // nil 时不输出，指向 0 时仍输出 (omitempty 只会忽略 nil 指针)
}

// 扩展：反序列化同样遵守这个规则

// `json.Unmarshal` 解析 JSON 字符串到结构体，也要求结构体字段首字母大写，才能被 JSON 解析器访问到。否则解析不进去，值全是零值。
// `json.Unmarshal` 反序列化时，JSON 的 key 会和结构体的字段名进行匹配。

func main() {
	// 1. 创建结构体变量

	p1 := person8_1{
		name:     "峰峰",
		age:      18,
		password: "123456",
	}

	// 结构体字段全部小写，在json序列化时，字段名会被忽略，导致序列化结果为空对象{}
	byteData, _ := json.Marshal(p1)

	p1Str := string(byteData)

	fmt.Printf("p1Str: %#v \n", p1Str) // p1Str: "{}"

	fmt.Println("----------------分割线---111-----------------")

	// 2.

	Age2 := 0

	p2 := person8_2{
		Name:       "sky",
		Age:        18,
		password:   "123456",
		Password2:  "345", // json 忽略
		Password3:  "567",
		DefaultStr: "", // 设置了 omitempty，默认零值 "" 不会被序列化，就算赋值为 "" 也不会序列化。
		Age2:       &Age2,
	}

	byteDat2, _ := json.Marshal(p2)

	p2Str := string(byteDat2)

	fmt.Printf("p2Str: %#v \n", p2Str) // p2Str: "{\"name\":\"sky\",\"Age\":18,\"Password1\":\"\",\"_\":\"567\",\"age2\":0}"
	fmt.Printf("p2Str: %#s \n", p2Str) // p2Str: {"name":"sky","Age":18,"Password1":"","_":"567","age2":0}
}
