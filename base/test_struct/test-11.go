package main

import "fmt"

// 泛型结构体

type teacher11_1[P string | int, A string | bool] struct {
	phone   P
	address A
}

func main() {
	teacher1 := teacher11_1[string, bool]{ // 指定泛型类型为 string 和 bool, 必须指定泛型类型，否则会报错。
		phone:   "123456789",
		address: false,
	}

	fmt.Printf("teacher1: %#v, type: %T \n", teacher1, teacher1) // teacher1: main.teacher11_1[string,bool]{phone:"123456789", address:false}, type: main.teacher11_1[string,bool]
}
