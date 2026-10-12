package main

import (
	"fmt"
	"os"
)

// 目录操作

// 创建目录
func testFn4_1() {
	err := os.Mkdir("./abc", 0666)

	if err != nil {
		fmt.Println("testFn4_1 err ==>", err)
	}
}

func testFn4_2() {
	err := os.MkdirAll("./aaa/bbb/ccc", 0777) // 创建多级目录 必须带执行位 07xx, 否则创建好一级目录后，无法进入该目录创建子目录，直接报 permission denied 错误.

	if err != nil {
		fmt.Println("testFn4_2 err ==>", err)
	}
}

// 删除目录
func testFn4_3() {
	err := os.Remove("./abc")

	if err != nil {
		fmt.Println("testFn4_3 err ==>", err)
	}
}

// 删除多级目录
func testFn4_4() {
	// err := os.RemoveAll("./aaa/bbb/ccc") // 只会删除指定的目录，不会删除上级目录.

	err := os.RemoveAll("./aaa") // 删除指定目录及其所有子目录

	if err != nil {
		fmt.Println("testFn4_4 err ==>", err)
	}
}

// 目录读取
func testFn4_5() {
	dir, err := os.ReadDir("./")

	if err != nil {
		fmt.Println("testFn4_5 err ==>", err)
		return
	}

	fmt.Println("testFn4_5 dir ==>", dir)

	for index, entry := range dir {
		fmt.Printf("index: %d, entry: %v, name: %s, isDir: %t \n", index, entry, entry.Name(), entry.IsDir())

		info, _ := entry.Info()

		fmt.Println("info.Size() ==>", info.Size())
		fmt.Println("------------------------------------------")
	}
}

func main() {
	// testFn4_1() // 创建单级目录

	// testFn4_2() // 创建多级目录

	// testFn4_3() // 删除单级目录

	// testFn4_4() // 删除多级目录

	testFn4_5() // 目录读取
}
