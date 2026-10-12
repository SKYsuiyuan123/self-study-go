package main

import "fmt"

// init 函数

// 建议：少用 init, 优先使用显式构造函数，因为 init 隐式执行、无法返回 error、不好测试。

// init 是 Go 的特殊函数，在包初始化时自动执行，用于初始化包级变量、设置配置、注册插件等。

// 特性：
//  1. 不能被显式调用，在 main 之前自动执行。
//  2. 同一包可定义多个，按源文件出现顺序执行。
//  3. 不能有参数，也不能有返回值。
//  4. 不能作为参数传递。
//  5. 跨包时，被导入的包先于导入者初始化。(如果包 A 导入了包 B，那么 B 的 init 一定先于 A 的 init 执行。)

var db int

func init() {
	fmt.Println("init 函数执行，初始化 db")

	db = 10

	fmt.Println("db 初始化完成，值为：", db) // db 初始化完成，值为： 10
}

func init() {
	fmt.Println("第二个 init 函数执行")
	fmt.Println("init2 db 的值为 ==>", db) // 10
}

// 函数的名字必须是 init 否则就是普通函数
func init3() {
	fmt.Println("第三个 init 函数执行")
}

func main() {
	fmt.Println("main 函数执行, db 的值为 ==>", db) // 10
}
