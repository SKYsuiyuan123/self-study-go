package main

import "fmt"

// 切片 —— 复制 与 截取

// 主要关注 复制后的切片与原切片的关系

func main() {
	// 1. 复制后的切片 与原切片 共享前 len 个元素 与 容量 (cap).

	arr1 := make([]int, 3, 5)

	fmt.Printf("arr1 ==> %#v \n", arr1) // []int{0, 0, 0}

	arr1[0], arr1[1], arr1[2] = 2, 7, 9

	brr1 := arr1 // 把 arr1 复制给 brr1，共用 前三个 元素。

	brr1[0] = 4

	fmt.Println("arr1[0] ==>", arr1[0]) // 4
	fmt.Println("brr1[0] ==>", brr1[0]) // 4

	fmt.Println("----------------分割线---111-----------------")

	// 2. append 返回后的切片 与原切片 共享前 len 个元素 与 容量 (cap).

	brr2 := append(arr1, 8)

	fmt.Printf("arr1 ==> %#v \n", arr1)                // []int{4, 7, 9}
	fmt.Printf("brr2 ==> %#v \n", brr2)                // []int{4, 7, 9, 8}
	fmt.Printf("&brr2 == &arr1 %v \n", &brr2 == &arr1) // false, 内存地址是 不一样的。

	fmt.Printf("arr1 len: %d, cap: %d \n", len(arr1), cap(arr1)) // arr1 len: 3, cap: 5
	fmt.Printf("brr2 len: %d, cap: %d \n", len(brr2), cap(brr2)) // brr2 len: 4, cap: 5

	fmt.Println("----------------分割线---222-----------------")

	arr1[0] = 1

	brr2 = append(brr2, 6)

	arr1[1] = 2

	fmt.Printf("arr1 ==> %#v \n", arr1) // []int{1, 2, 9}
	fmt.Printf("brr2 ==> %#v \n", brr2) // []int{1, 2, 9, 8, 6}

	fmt.Println("----------------分割线---333-----------------")

	// 3. append 扩容

	// 容量(cap) 不够时，需要申请一块新的内存空间 来存储切片。此时新地址的 前几个元素 就跟 arr1 没有关系了。

	brr2 = append(brr2, 6) // 会先把 brr2 旧的几个元素拷贝过来，然后追加新的元素。(brr2 旧的内存地址会舍弃掉)

	arr1[2] = 3 // brr2 已经是新的内存空间了，不会再受到 arr1 元素更改的影响。

	fmt.Printf("arr1 ==> %#v \n", arr1) // []int{1, 2, 3}
	fmt.Printf("brr2 ==> %#v \n", brr2) // []int{1, 2, 9, 8, 6, 6}

	fmt.Printf("arr1 len: %d, cap: %d \n", len(arr1), cap(arr1)) // arr1 len: 3, cap: 5
	fmt.Printf("brr2 len: %d, cap: %d \n", len(brr2), cap(brr2)) // brr2 len: 6, cap: 10

	fmt.Println("----------------分割线---444-----------------")

	// 4. 截取

	arr2 := []int{3, 4, 5}
	arr3 := arr2[:]

	fmt.Printf("arr2 ==> %#v \n", arr2) // []int{3, 4, 5}
	fmt.Printf("arr3 ==> %#v \n", arr3) // []int{3, 4, 5}

	arr3[0] = 1

	fmt.Printf("arr2 ==> %#v \n", arr2) // []int{1, 4, 5}
	fmt.Printf("arr3 ==> %#v \n", arr3) // []int{1, 4, 5}

	fmt.Printf("arr2 len: %d, cap: %d \n", len(arr2), cap(arr2)) // arr2 len: 3, cap: 3
	fmt.Printf("arr3 len: %d, cap: %d \n", len(arr3), cap(arr3)) // arr3 len: 3, cap: 3

	// 4.2 截取 后 扩容

	arr3 = append(arr3, 9) // arr3 扩容后 就跟 arr2 不再共享前几个元素

	arr3[1] = 2

	fmt.Printf("arr2 ==> %#v \n", arr2) // []int{1, 4, 5}
	fmt.Printf("arr3 ==> %#v \n", arr3) // []int{1, 2, 5, 9}

	fmt.Printf("arr2 len: %d, cap: %d \n", len(arr2), cap(arr2)) // arr2 len: 3, cap: 3
	fmt.Printf("arr3 len: %d, cap: %d \n", len(arr3), cap(arr3)) // arr3 len: 4, cap: 6
}
