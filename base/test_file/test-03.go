package main

import (
	"fmt"
	"io"
	"os"
)

// 文件操作 —— 复制、移动、拷贝、删除文件

// 一次性复制文件
func testFn3_1() {
	// 读文件
	byteStr, err := os.ReadFile("./hello5.txt") // 读取文件内容, 返回 []byte 类型的数据.

	if err != nil {
		fmt.Println("testFn3_1 err ==>", err)
		return
	}

	// 写文件
	err2 := os.WriteFile("./hello6.txt", byteStr, 0666) // 一次性写入文件内容, 如果文件不存在则创建新文件, 如果文件存在则覆盖原有内容.

	if err2 != nil {
		fmt.Println("testFn3_1 err2 ==>", err2)
		return
	}

	fmt.Println("testFn3_1 success...")
}

// 文件流复制文件
func testFn3_2() {
	file, err1 := os.Open("./hello5.txt")                                         // 打开文件, 返回 *os.File 类型的文件对象.
	writeFile, err2 := os.OpenFile("./hello7.txt", os.O_CREATE|os.O_WRONLY, 0666) // 创建文件, 如果文件不存在则创建新文件, 如果文件存在则覆盖原有内容.

	if err1 != nil || err2 != nil {
		fmt.Println("testFn3_2 err1 ==>", err1)
		fmt.Println("testFn3_2 err2 ==>", err2)
		return
	}

	// 不关闭文件，程序长时间运行会 导致文件句柄泄漏，最终导致程序崩溃。
	defer file.Close()
	defer writeFile.Close()

	var tempSlice = make([]byte, 128) // 创建一个临时切片, 用于存储读取的文件内容.

	// 循环读取文件内容, 并写入到新文件中.一边读取，一边写入。
	for {
		n1, err3 := file.Read(tempSlice) // 读取文件内容, 返回读取的字节数和错误信息.

		if err3 == io.EOF {
			break // 文件读取完毕, 跳出循环.
		}
		if err3 != nil {
			fmt.Println("testFn3_2 err3 ==>", err3)
			return
		}

		_, err4 := writeFile.Write(tempSlice[:n1]) // 写入文件内容, 返回写入的字节数和错误信息.

		if err4 != nil {
			fmt.Println("testFn3_2 err4 ==>", err4)
			return
		}
	}

	fmt.Println("testFn3_2 success...")
}

// 文件流复制文件 (io.Copy)
func testFn3_2_2() {
	file, err1 := os.Open("./hello5.txt")
	writeFile, err2 := os.OpenFile("./hello9.txt", os.O_CREATE|os.O_WRONLY, 0666)

	if err1 != nil || err2 != nil {
		fmt.Println("testFn3_2 err1 ==>", err1)
		fmt.Println("testFn3_2 err2 ==>", err2)
		return
	}

	// 不关闭文件，程序长时间运行会 导致文件句柄泄漏，最终导致程序崩溃。
	defer file.Close()
	defer writeFile.Close()

	// 拷贝文件
	n, err3 := io.Copy(writeFile, file) // 将文件内容从 file 复制到 writeFile 中, 返回复制的字节数和错误信息.

	if err3 != nil {
		fmt.Println("testFn3_2 err3 ==>", err3)
		return
	}

	fmt.Println("n ==>", n) // n ==> 10
}

// 删除文件
func testFn3_3() {
	err := os.Remove("./hello7.txt") // 删除文件, 如果文件不存在则返回错误信息.

	// 如果文件不存在
	if os.IsNotExist(err) {
		fmt.Println("testFn3_3 ./hello7.txt 文件不存在")
		return
	}

	if err != nil {
		fmt.Println("testFn3_3 err ==>", err)
		return
	}
}

// 文件重命名
func testFn3_4() {
	err := os.Rename("./hello6.txt", "./hello8.txt") // 重命名文件, 如果文件不存在则返回错误信息.

	if err != nil {
		fmt.Println("testFn3_4 err ==>", err)
		return
	}

	fmt.Println("testFn3_4 success...")
}

func main() {
	// testFn3_1() // 一次性复制文件

	// testFn3_2() // 文件流复制文件

	testFn3_2_2() // 文件流复制文件 (io.Copy)

	// testFn3_3() // 删除文件

	// testFn3_4() // 文件重命名
}
