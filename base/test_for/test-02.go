package main

import "fmt"

// for - label

func main() {

label1:
	for i := 0; i < 10; i++ {
		for k := 0; k < 10; k++ {
			for j := 0; j < 10; j++ {
				if i == 5 {
					break label1 // 使用 break label 跳出多层循环
				}

				fmt.Println("i1, k1, j1 ==>", i, k, j)
			}
		}
	}

	fmt.Println("----------------分割线---111-----------------")

label2:
	for i := 0; i < 5; i++ {
		for k := 0; k < 5; k++ {
			for j := 0; j < 5; j++ {
				if i == 3 {
					continue label2 // 使用 continue label 跳到最外层 for 继续下一次 循环
				}

				fmt.Println("i2, k2, j2 ==>", i, k, j)
			}
		}
	}
}
