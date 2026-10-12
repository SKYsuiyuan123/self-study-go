package main

import "fmt"

// fmt.Print 使用

// 			println				  												fmt.Println
// 来源	内置函数，无需 import	标准库 fmt 包，			需 import
// 输出目标	stderr（标准错误）									stdout（标准输出）
// 用途	调试																	正式输出
// 格式化能力	很弱															强大（配 Printf）
// 对结构体/指针	只给粗略信息											调用 String()，可读
// 返回值	无																		(n int, err error)
// 稳定性	不保证																官方保证

func printDemo() {
	fmt.Printf("%s 哇，你好美 \n", "老婆")   // 老婆 哇，你好美
	fmt.Printf("%d \n", 3)            // 3
	fmt.Printf("%.2f \n", 3.1415926)  // 3.14
	fmt.Printf("%T %T \n", "你好", 2.5) // string float64
	fmt.Printf("%v \n", "你好")         // 你好
	fmt.Printf("%v \n", "")           // (打印时，看不出来是空字符串)
	fmt.Printf("%#v \n", "")          // "" (打印时 可以看出来是空字符串)

	var f1 = fmt.Sprintf("%.2f \n", 3.1415926)
	fmt.Printf("f1 ==> %s%T \n", f1, f1) // f1 ==> 3.14 string

	fmt.Println("----------------分割线---printDemo-----------------")
}

func main() {
	fmt.Print("Please input your score: ")

	var score int

	// fmt.Scanf("%d", &score)

	// 等价
	fmt.Scan(&score)

	fmt.Printf("%v, type: %T \n", score, score) // 20, type: int

	fmt.Println("----------------分割线---111-----------------")

	printDemo()
}
