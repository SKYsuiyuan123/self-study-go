package main

import (
	"encoding/json"
	jsonv2 "encoding/json/v2"
	"fmt"
)

// JSON 的 v1 与 v2 版本对比

// encoding/json（v1）和 encoding/json/v2 的主要区别在于序列化行为、API 设计和错误处理哲学。
// v2 的默认行为更严格，更符合 JSON 规范，但也因此与 v1 不完全兼容。

// 核心差异对比：

// 行为									v1 (encoding/json)				v2 (encoding/json/v2)
// nil slice/map				 序列化为 null							 序列化为 [] / {}
// 字段名匹配							 大小写不敏感									大小写严格匹配
// 重复 JSON key					允许，后者覆盖前者						直接报错
// 无效 UTF-8							静默替换为 �									直接报错
// omitempty 语义					false/0/空集合都忽略				仅 JSON 层面为空时忽略（推荐改用 omitzero）

type person9_2 struct {
	ID     int
	gender string // 未导出，v1 和 v2 都会忽略。
	Name   string
}

func main() {
	// 1. v1 和 v2 比对

	p1 := person9_2{
		ID:   1,
		Name: "Alice",
	}

	// v1
	p1Byte, err1 := json.Marshal(p1)

	if err1 != nil {
		fmt.Println("json.Marshal err1 ==> ", err1)
	}

	p1StrJson := string(p1Byte)

	fmt.Printf("v1 p1Byte: %s \n", p1StrJson) // v1 p1Byte: {"ID":1,"Name":"Alice"}

	// v2
	p2Byte, err2 := jsonv2.Marshal(p1)

	if err2 != nil {
		fmt.Println("json.Marshal err2 ==> ", err2)
	}

	p2StrJson := string(p2Byte)

	fmt.Printf("v2 p2StrJson: %s \n", p2StrJson) // v2 p2Byte: {"ID":1,"Name":"Alice"}

	fmt.Println("----------------分割线---111-----------------")

	// 2. 核心差异

	// 2.1 nil slice
	type container struct {
		Items []string
	}

	var c1 container // c1.Items == nil

	b1, _ := json.Marshal(c1)
	b2, _ := jsonv2.Marshal(c1)

	fmt.Printf("b1: %s \n", string(b1)) // b1: {"Items":null}
	fmt.Printf("b2: %s \n", string(b2)) // b2: {"Items":[]}
}
