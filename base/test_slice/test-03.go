package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

// 切片 —— 排序

// 标准 sot 包。
// 升序：sort.Ints(), sort.Float64s(), sort.Strings()函数。默认都是从小到大排序。
// 降序：
// 	sort.Sort(sort.Reverse(sort.IntSlice(intList))),
// 	sort.Sort(sort.Reverse(sort.Float64Slice(float8List))),
// 	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))

// 升序 sort.Ints(slice100)
// 降序 sort.Sort(sort.Reverse(sort.IntSlice(slice100)))

func main() {
	slice1 := []int{2, 5, 3, 7, 8}

	// 1. 使用 sort.Ints 排序 (默认 升序)

	sort.Ints(slice1)
	fmt.Printf("slice1 %#v \n", slice1) // []int{2, 3, 5, 7, 8}

	// 2. 使用 sort.Sort(sort.Reverse(sort.IntSlice(slice1))) 降序

	sort.Sort(sort.Reverse(sort.IntSlice(slice1)))
	fmt.Printf("slice1 %#v \n", slice1) // []int{8, 7, 5, 3, 2}

	fmt.Println("----------------分割线---111-----------------")

	// 3. Go 1.21+ 优先使用 slice.SortFunc 排序 (速度更快)

	slice2 := []int{1, 2, 5, 6, 3, 2}

	// 3.1 降序
	slices.SortFunc(slice2, func(a, b int) int {
		return b - a
	})

	fmt.Printf("slice2 %#v \n", slice2) // []int{6, 5, 3, 2, 2, 1}

	// 3.2 升序
	slices.SortFunc(slice2, func(a, b int) int {
		return a - b
	})

	fmt.Printf("slice2 %#v \n", slice2) // []int{1, 2, 2, 3, 5, 6}

	// 3.3 字符串排序
	slice3 := []string{"a", "b", "c", "f", "e", "c"}

	// slices.Sort(slice3) // 标准库用的是更高效的排序（pdqsort）

	// 或

	// slices.SortFunc(slice3, func(a, b string) int {
	// 	return strings.Compare(a, b)
	// })

	// 或

	slices.SortFunc(slice3, strings.Compare)

	fmt.Printf("slice3 %#v \n", slice3) // []string{"a", "b", "c", "c", "e", "f"}

	fmt.Println("----------------分割线---222-----------------")

	// 4. 手动排序 (从小到大)

	// 4.1 选择排序 O(n²)

	numSlice1 := []int{9, 8, 7, 2, 3, 1}

	for i := 0; i < len(numSlice1)-1; i++ { // 最后一位不要再比较
		minIndex := i

		for j := i + 1; j < len(numSlice1); j++ {
			if numSlice1[j] < numSlice1[minIndex] { // ✅ 和当前最小值比
				minIndex = j
			}
		}

		// 如果某一轮 minIndex 没变，说明当前位置已经是最小值，可以跳过交换
		if minIndex != i {
			// 每次循环完 i 后 交换一次值，减少了交互的次数。
			numSlice1[i], numSlice1[minIndex] = numSlice1[minIndex], numSlice1[i]
		}
	}

	fmt.Printf("numSlice1 ==> %#v \n", numSlice1) // []int{1, 2, 3, 7, 8, 9}

	// 标准库用的是更高效的排序（pdqsort）

	numSlice2 := []int{9, 8, 7, 2, 3, 1}

	slices.Sort(numSlice2)
	fmt.Printf("slices.Sort(numSlice2) ==> %#v \n", numSlice2) // []int{1, 2, 3, 7, 8, 9}

	// 4.2 冒泡排序

	numSlice3 := []int{9, 8, 7, 2, 3, 1, 2, 2}

	for i := 0; i < len(numSlice3)-1; i++ { // 最后一轮只剩一个元素，天然有序，不用再跑
		swapped := false

		for j := 0; j < len(numSlice3)-1-i; j++ { // -1-i: 后边的排序过了，不用再参与排序。
			if numSlice3[j] > numSlice3[j+1] {
				numSlice3[j], numSlice3[j+1] = numSlice3[j+1], numSlice3[j]
				swapped = true
			}
		}

		// 如果某一轮一次交换都没发生，说明数组已经有序，可以直接结束。对“接近有序”的数据提升巨大
		if !swapped {
			break // 已经有序，提前结束
		}
	}

	fmt.Printf("numSlice3 ==> %#v \n", numSlice3) // []int{1, 2, 2, 2, 3, 7, 8, 9}

	// 4.3 改进冒泡排序

	numSlice4 := []int{9, 8, 7, 2, 3, 1}

	n := len(numSlice4)
	for n > 1 {
		newN := 0
		for j := 0; j < n-1; j++ {
			if numSlice4[j] > numSlice4[j+1] {
				numSlice4[j], numSlice4[j+1] = numSlice4[j+1], numSlice4[j]
				newN = j + 1 // 进一步优化：每轮记录最后一次发生交换的位置，该位置之后已经有序，下一轮只需比到这里。
			}
		}
		n = newN
	}

	fmt.Printf("numSlice4 ==> %#v \n", numSlice4) // []int{1, 2, 3, 7, 8, 9}
}
