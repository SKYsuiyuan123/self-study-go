package main

import "fmt"

// 泛型 —— 切片、map

// 泛型切片
type mySliceType1[T any] []T

// 泛型 map
type myMapType1[K string | int, V any] map[K]V // map 的 key 不能是 any 类型，只能是 string、int、float、bool、指针、channel、interface 等类型

func main() {
	// 1. 泛型切片

	// 1.1
	var slice1 mySliceType1[string]
	slice1 = append(slice1, "hello")

	fmt.Printf("slice1: %#v \n", slice1) // slice1: main.mySliceType1[string]{"hello"}

	// 1.2
	slice2 := mySliceType1[int]{1, 2, 3}
	fmt.Printf("slice2: %#v \n", slice2) // slice2: main.mySliceType1[int]{1, 2, 3}

	fmt.Println("----------------分割线---111-----------------")

	// 2. 泛型 map

	// 2.1
	var map1 myMapType1[string, int]
	map1 = make(myMapType1[string, int])

	map1["a"] = 1

	fmt.Printf("map1: %#v \n", map1) // map1: main.myMapType1[string,int]{"a":1}

	// 2.2
	var map2 = make(myMapType1[string, string])
	map2["b"] = "hello"

	fmt.Printf("map2: %#v \n", map2) // map2: main.myMapType1[string,string]{"b":"hello"}
}
