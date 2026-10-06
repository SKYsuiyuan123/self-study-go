package main

import "fmt"

// 常量类型

// const 来定义枚举类型 (iota 只能出现在 const 初始化常量里)

const (
	bei_jing  = iota // 第一行的 iota 默认值是 0，每行的 iota 都会累加 1.
	shang_hai        // int 1
	shen_zhen        // int 2
)

const (
	bei_jing2  = iota * 10
	shang_hai2 // 自动累加 10. iota = 1, 所以 shang_hai2 = 1 * 10 = 10
	shen_zhen2
)

const (
	bei_jing3, bei_jing4   = iota + 1, iota + 2 // 1, 2
	shang_hai3, shang_hai4                      // iota = 1, 所以 shang_hai3 = 1 + 1 = 2, shang_hai4 = 1 + 2 = 3
	shen_zhen3, shen_zhen4                      // iota = 2, 所以 shen_zhen3 = 2 + 1 = 3, shen_zhen4 = 2 + 2 = 4
	shen_zhen5, shen_zhen6 = iota * 2, iota * 3 // iota = 3, 所以 shen_zhen5 = 3 * 2 = 6, shen_zhen6 = 3 * 3 = 9
	shen_zhen7, shen_zhen8                      // iota = 4, 所以 shen_zhen7 = 4 * 2 = 8, shen_zhen8 = 4 * 3 = 12
)

func main() {

	fmt.Println("bei_jing ==>", bei_jing)   // 0
	fmt.Println("shang_hai ==>", shang_hai) // 1
	fmt.Println("shen_zhen ==>", shen_zhen) // 2

	fmt.Println("----------------分割线---111-----------------")

	fmt.Println("bei_jing2 ==>", bei_jing2)   // 0
	fmt.Println("shang_hai2 ==>", shang_hai2) // 10
	fmt.Println("shen_zhen2 ==>", shen_zhen2) // 20

	fmt.Println("----------------分割线---222-----------------")
}
