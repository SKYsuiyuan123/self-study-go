package main

import "fmt"

// 数组

// go 的 数组类型是 值类型

// 多维数组，只有最外层可以使用 [...]。完整用法 arr1 := [...][2]int{{1, 2}, {3, 4}}

func main() {
	// 1. 数组求平均值
	// 注意 float64() 的使用

	var arr1 = [3]int{1, 2, 3}

	var arr1Sum = 0
	for _, item := range arr1 {
		arr1Sum += item
	}

	arr1Avg := float64(arr1Sum) / float64(len(arr1))

	fmt.Printf("arr1 所有元素的和是: %v, 平均值: %.2f \n", arr1Sum, arr1Avg) // arr1 所有元素的和是: 6, 平均值: 2.00

	fmt.Println("----------------分割线---111-----------------")

	// 2. 默认零值

	var arr2 [4]string              // string 的默认零值是 ""
	fmt.Printf("arr2 %#v \n", arr2) // [4]string{"", "", "", ""}

	for index, item := range arr2 {
		fmt.Printf("arr2 index: %#v, item: %#v \n", index, item)
		// arr2 index: 0, item: ""
		// arr2 index: 1, item: ""
		// arr2 index: 2, item: ""
		// arr2 index: 3, item: ""
	}

	fmt.Println("----------------分割线---222-----------------")

	// 3. 自动推导数组长度

	arr3 := [...]int{1, 2, 3, 4}
	fmt.Println("arr3, len(arr3) ==>", arr3, len(arr3)) // [1 2 3 4] 4

	// 4. 初始化时跳过 中间的几项

	arr4 := [4]int{0: 2, 3: 3}
	fmt.Printf("arr4 ==> %#v, len(arr4): %v \n", arr4, len(arr4)) // arr4 ==> [4]int{2, 0, 0, 3}, len(arr4): 4

	fmt.Println("----------------分割线---333-----------------")

	// 5. 数组的 复制
	// 数组 是 值类型复制

	arr5 := [...]int{1, 2}
	arr6 := arr5
	arr6[1] = 3 // arr6 是新的内存的数组，不会改变 arr5 对应位置的内容。

	fmt.Printf("arr5 ==> %#v \n", arr5) // [2]int{1, 2}
	fmt.Printf("arr6 ==> %#v \n", arr6) // [2]int{1, 3}

	fmt.Println("----------------分割线---444-----------------")

	// 6. 多维数组

	arr7 := [2][3]int{{1, 2, 3}, {2, 3, 4}}
	fmt.Printf("arr7 ==> %#v \n", arr7) // [2][3]int{[3]int{1, 2, 3}, [3]int{2, 3, 4}}
	fmt.Println("arr7[0] ==>", arr7[0]) // [1 2 3]

	// 6.2 多维数组 外层自动推导数组长度

	arr8 := [...][2]int{{1, 2}, {3, 4}}
	fmt.Printf("arr8 ==> %#v \n", arr8)     // [2][2]int{[2]int{1, 2}, [2]int{3, 4}}
	fmt.Println("arr8 ==>", arr8)           // [[1 2] [3 4]]
	fmt.Println("len(arr8) ==>", len(arr8)) // 2
}
