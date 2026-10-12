package main

import (
	"fmt"
	"reflect"
	"strings"
)

// 反射 —— 修改结构体

type student5_1 struct {
	Name1 string `big:"_"` // 这里不能写成 `big: "_"`
	Name2 string
}

func setStructField(obj any) {
	v1 := reflect.ValueOf(obj).Elem() // 获取指针指向的值
	t1 := reflect.TypeOf(obj).Elem()  // 获取指针指向的类型

	for i := 0; i < v1.NumField(); i++ {
		val := v1.Field(i)
		big := t1.Field(i).Tag.Get("big")

		fmt.Println("value ==>", val)

		if big == "" {
			continue
		}

		// 修改值
		val.SetString(strings.ToUpper(val.String()))
	}
}

func main() {
	s1 := student5_1{
		Name1: "blue",
		Name2: "sky",
	}

	setStructField(&s1)

	fmt.Printf("s1: %#v \n", s1) // s1: main.student5_1{Name1:"BLUE", Name2:"sky"}
}
