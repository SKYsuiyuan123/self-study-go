package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// 文件操作 —— 写入文件

func testWriteFn2_1() {
	file, err := os.OpenFile("./hello3.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666) // 创建文件, 如果文件存在则清空内容, 如果文件不存在则创建新文件.
	//  0666 是文件权限, 代表所有用户都可以读写该文件.0777 是文件权限, 代表所有用户都可以读写执行该文件.

	if err != nil {
		fmt.Println("testWriteFn1_6 err ==>", err)
		return
	}

	defer file.Close()

	for range 4 {
		// 写入文件内容 —— string
		// file.WriteString("春眠不觉晓\r\n")

		// 写入文件 —— []byte
		str := "春眠不觉晓2\r\n"
		byteData := []byte(str)
		file.Write(byteData)
	}
}

// bufio 写入文件
func testWriteFn2_2() {
	file, err := os.OpenFile("./hello4.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666) // 创建文件, 如果文件存在则清空内容, 如果文件不存在则创建新文件.

	if err != nil {
		fmt.Println("testWriteFn2_2 err ==>", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file) // 创建一个带缓冲的写入器, 可以提高写入效率.

	for range 4 {
		// 写入缓存
		writer.WriteString("hello go, go, go\r\n") // 写入文件内容
	}

	// Flush 将缓冲区的内容写入文件, 如果不调用 Flush, 则可能会导致部分内容未写入文件.
	writer.Flush()
}

// os.WriteFile 一次性写入文件内容
func testWriteFn2_3() {
	str := "Hello Go!"
	err := os.WriteFile("./hello5.txt", []byte(str), 0666) // 一次性写入文件内容, 如果文件不存在则创建新文件, 如果文件存在则覆盖原有内容.

	if err != nil {
		fmt.Println("testWriteFn2_3 err ==>", err)
		return
	}
}

func main() {
	timeStart := time.Now()

	// testWriteFn2_1() // 写入文件内容

	// testWriteFn2_2() // bufio 写入文件

	testWriteFn2_3() // os.WriteFile 一次性写入文件内容

	fmt.Println("程序执行耗时 ==>", time.Since(timeStart)) // 程序执行耗时 ==> 1.019958ms
}
