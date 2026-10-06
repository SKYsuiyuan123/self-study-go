package main

import "fmt"

// fmt.Print 使用

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
