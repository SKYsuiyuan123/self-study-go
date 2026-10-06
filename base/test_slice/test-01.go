package main

import (
	"fmt"
	"slices"
)

// 切片

// 1. 切片是引用类型
// 2. 切片只声明不赋值，那么它的值就是 nil
// 3. 切片的底层是一个数组，可以基于数组定义切片，使用 arr1[:]
// 4. len() 获取切片的长度，cap() 获取切片的容量。
// 		长度：切片的长度就是它所包含的元素个数。
// 		容量：切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
// 5. 切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度(len)、和切片的容量(cap)。

func main() {
	// 1. 切片初始化
	slice1 := []string{"a", "b", "c"}
	slice2 := []byte{'a', 'b', 'c'}

	fmt.Println("slice1 ==>", slice1) // [a b c]
	fmt.Println("slice2 ==>", slice2) // [97 98 99]

	fmt.Printf("slice1 val: %#v -- type: %T \n", slice1, slice1) // slice1 val: []string{"a", "b", "c"} -- type: []string
	fmt.Printf("slice2 val: %#v -- type: %T \n", slice2, slice2) // slice2 val: []byte{0x61, 0x62, 0x63} -- type: []uint8

	fmt.Println("----------------分割线---111-----------------")

	// 2. 获取切片 长度 和 容量

	fmt.Printf("slice1 len(slice1): %v cap(slice1) %v \n", len(slice1), cap(slice1)) // slice1 len(slice1): 3 cap(slice1) 3
	fmt.Printf("slice1 len(slice2): %v cap(slice2) %v \n", len(slice2), cap(slice2)) // slice1 len(slice2): 3 cap(slice2) 3

	fmt.Println("----------------分割线---222-----------------")

	// 3. 切片只声明不赋值

	var slice3 []int

	// 3.1 只声明不赋值是 nil
	if slice3 == nil {
		fmt.Println("切片只声明不赋值，那么它的值就是 nil") // 切片只声明不赋值，那么它的值就是 nil
	}

	// slice3[0] = 2                                 // nil slice 可以正常索引，但只能读，不能写. 会触发 panic。
	// fmt.Printf("slice3[0] ==> %#v \n", slice3[0]) // nil slice 读取元素时，也会触发 panic。

	// 3.2 长度 和 容量都是 0

	fmt.Printf("slice3: %v -- type: %T, len: %v, cap: %v \n", slice3, slice3, len(slice3), cap(slice3)) // slice3: [] -- type: []int, len: 0, cap: 0

	// 3.3 赋值为空

	var slice4 = []int{}

	// 3.3 赋值为空 则不是 nil
	if slice4 != nil {
		fmt.Println("切片赋值为空，那么它的值 不是 nil") // 切片赋值为空，那么它的值 不是 nil
	}

	// slice4[0] = 2                                 // 空 slice 可以正常索引，但只能读，不能写. 会触发 panic。
	// fmt.Printf("slice4[0] ==> %#v \n", slice4[0]) // 空 slice 读取元素时，也会触发 panic。

	fmt.Println("----------------分割线---333-----------------")

	// 4. 切片扩容
	// slice 扩容，当容量放不下长度时，就会一直扩容。扩容大小每次都会 乘以 2。新版本有变化。
	// 容量(cap): 开头到结尾指针的距离。

	// 4.1 只声明的切片扩容
	var slice5 []int

	// slice1[10] = "d" // 不能这样扩容, 会触发 panic。

	slice5 = append(slice5, 1) // 添加元素 —— 扩容操作 ✅ append 会给它分配底层数组

	fmt.Printf("slice5 ==> %#v \n", slice5)                                                  // []int{1}
	fmt.Printf("slice5: -- type: %T, len: %v, cap: %v \n", slice5, len(slice5), cap(slice5)) // type: []int, len: 1, cap: 1

	slice5 = append(slice5, 3, 2)
	fmt.Printf("slice5: %#v \n", slice5)                               // slice5: []int{1, 3, 2}
	fmt.Printf("slice5 len: %v, cap: %v \n", len(slice5), cap(slice5)) // slice5 len: 3, cap: 3

	slice5 = append(slice5, 4)                                         // cap 由 3 变为 6
	fmt.Printf("slice5 len: %v, cap: %v \n", len(slice5), cap(slice5)) // slice5 len: 4, cap: 6

	// 4.2 在切片中插入内容

	slice5_2 := []int{1, 2, 3, 5}

	// 4.2.1 插入切片前边
	slice5_2 = append([]int{-1, 0, 1}, slice5_2...)

	fmt.Printf("slice5_2 ==> %#v \n", slice5_2) // []int{-1, 0, 1, 1, 2, 3, 5}

	// 4.2.2 插入切片中间
	// 中间插入 ==> [-1, 0, 1, 1, 2] [插入的元素] [3, 5]

	idx1 := 5 // 在索引 5 之前插入

	slice5_2 = append(slice5_2[:idx1], append([]int{9, 9, 9}, slice5_2[idx1:]...)...)

	fmt.Printf("slice5_2 ==> %#v \n", slice5_2) // []int{-1, 0, 1, 1, 2, 9, 9, 9, 3, 5}

	fmt.Println("----------------分割线---444-----------------")

	// 5. 切片 截取
	// 截取语法: slice1[开始索引:结束索引] 不包含结束索引。(左闭右开 区间) (0 索引 和 结束索引 可以省略)

	slice6 := []int{55, 56, 57, 58, 59}

	slice7 := slice6[1:3]

	slice8 := slice6[:] // 从 索引 0 截取到 结尾

	fmt.Printf("slice7 ==> %#v \n", slice7) // []int{56, 57}
	fmt.Printf("slice8 ==> %#v \n", slice8) // []int{55, 56, 57, 58, 59}

	fmt.Println("----------------分割线---555-----------------")

	// 6. 切片合并

	slice10 := []string{"php", "go"}
	slice12 := []string{"node", "java"}

	// slice11 := []int{1, 2, 3}
	// slice11 = append(slice11, slice12) // 不同类型的切片 不能合并

	slice12 = append(slice12, slice10...)
	fmt.Printf("slice12: %#v \n", slice12) // slice12: []string{"node", "java", "php", "go"}

	fmt.Println("----------------分割线---666-----------------")

	// 7. 切片复制

	// 使用 copy 函数复制切片

	// 切片复制完之后，新的 切片 内容变更就不会 影响旧切片 的内容。

	// 7.1 复制后的 新切片和 旧切片 长度和容量完全一致。
	slice13 := []int{1, 2, 3, 5}

	slice14 := make([]int, len(slice13), cap(slice13))

	// 或者
	// slice14 := make([]int, 4, 4)

	copy(slice14, slice13) // 将 slice13 的内容复制给 slice14

	slice14[0] = 6 // 不会改变 slice13 对应位置 的元素

	fmt.Printf("slice13: %#v \n", slice13) // slice13: []int{1, 2, 3, 5}
	fmt.Printf("slice14: %#v \n", slice14) // slice14: []int{6, 2, 3, 5}

	// 7.2 复制后的 新切片 长度和容量小于 原来的
	slice15 := make([]int, 2, 2)

	copy(slice15, slice13)

	fmt.Printf("slice15: %#v \n", slice15) // slice15: []int{1, 2}

	// 7.3 复制后的 新切片 长度和容量大于 原来的
	slice16 := make([]int, 6, 6)

	copy(slice16, slice13) // 复制后 新的切片 没有对应位置元素的 值都是 默认零值。

	fmt.Printf("slice16: %#v, len: %v, cap: %v \n", slice16, len(slice16), cap(slice16)) // slice16: []int{1, 2, 3, 5, 0, 0}, len: 6, cap: 6

	// 7.4 使用 Slices.clone 克隆一个切片

	slice17 := slices.Clone(slice13)

	slice17[0] = 100

	fmt.Printf("slice13: %#v \n", slice13) // slice13: []int{1, 2, 3, 5}
	fmt.Printf("slice17: %#v \n", slice17) // slice17: []int{100, 2, 3, 5}

	fmt.Println("----------------分割线---777-----------------")

	// 8. 删除切片元素

	// 8.1 可以使用 切片本身的特性来删除元素。

	slice20 := []int{1, 2, 3, 4}

	// 删除掉 第二个元素
	slice20 = append(slice20[:2], slice20[3:]...)
	fmt.Printf("slice20: %#v \n", slice20) // slice20: []int{1, 2, 4}

	// 8.2 使用 slices.Delete(切片, 起始索引, 结束索引) 删除切片元素，返回一个新的切片。

	slice21 := []int{1, 2, 3, 4}
	newSlice21 := slices.Delete(slice21, 1, 2) // 只包含开始索引，不包含结束索引。

	fmt.Printf("slice21: %#v \n", slice21)       // slice21: []int{1, 3, 4, 0}
	fmt.Printf("newSlice21: %#v \n", newSlice21) // newSlice21: []int{1, 3, 4}

	fmt.Println("----------------分割线---888-----------------")

	// 9. 利用切片特性 修改 string 类型的值

	str1 := "你好golang"
	runeStr1 := []rune(str1)

	fmt.Printf("runeStr1: %#v \n", runeStr1) // runeStr1: []int32{20320, 22909, 103, 111, 108, 97, 110, 103}

	runeStr1[0] = '不' // 注意：这里是单引号(rune 类型)

	str2 := string(runeStr1)

	fmt.Printf("str2 val: %#v -- type: %T \n", str2, str2) // str2 val: "不好golang" -- type: string
}
