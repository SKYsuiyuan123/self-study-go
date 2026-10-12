package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// 文件操作 —— 读取文件

// 文件操作 (并不能 真正读取完文件)
func testFn1_1() {
	// 只读方式打开文件
	file, err := os.Open("./hello.txt")

	if err != nil {
		fmt.Println("testFn1_1 err:", err)
		return
	}

	fmt.Println("file ==>", file) // file ==> &{0x528471080180}

	// 必须得关闭文件
	defer file.Close()

	// 读取文件内容
	slice1 := make([]byte, 1024) // 每次读取 1024 个字节

	n, err2 := file.Read(slice1)

	if err2 != nil {
		fmt.Println("testFn1_1 err2:", err2)
		return
	}

	fmt.Printf("读取了 %d 个字节 \n", n) //  读取了 25 个字
	fmt.Println("testFn1_1 读取到的内容 ==>", string(slice1[:n]))

	// 好好学习
	// 天天向上
}

func testFn1_2() {
	// 只读方式打开文件
	file, err := os.Open("./hello.txt")

	if err != nil {
		fmt.Println("testFn1_2 err ==>", err)
		return
	}

	defer file.Close()

	fmt.Println("file ==>", file) // file ==> &{0x528471080180} file 的地址

	// 读取文件内容
	var strSlice []byte
	slice1 := make([]byte, 128) // 每次读取 128 字节

	// 流式读取
	for {
		n, err2 := file.Read(slice1)

		if err2 == io.EOF {
			fmt.Println("读取完毕 ==>", err2)
			break
		}

		if err2 != nil {
			fmt.Println("testFn1_2 err2 ==>", err2)
			break
		}

		fmt.Printf("读取了 %d 个字节 \n", n) //  读取了 15 个字 (最后一次 只有15个字符)

		// 追加读取到的内容. 最后一次读取到的内容可能不足 128 个字节, 所以要用 slice1[:n]。
		// 或者每次读取完都把 slice1 清空, 但是这样会浪费内存空间, 所以不推荐。
		strSlice = append(strSlice, slice1[:n]...)

		// 读取完之后要清空 slice1, 否则会有脏数据
		// slice1 = make([]byte, 128)
	}

	fmt.Println("testFn1_2 读取到的内容 ==>", string(strSlice)) // 读取到的内容 ==>
}

// 使用 bufio 读取文件
func testFn1_3() {
	file, err := os.Open("./hello.txt")

	if err != nil {
		fmt.Println("testFn1_3 err ==>", err)
		return
	}

	defer file.Close()

	// bufio 读取文件
	reader := bufio.NewReader(file)
	var fileStr string // 字符串拼接性能较差, 复杂度 O(n^2), 因为每次拼接都会创建一个新的字符串, 所以不推荐使用.

	// 流式读取
	for {
		str, err2 := reader.ReadString('\n') // 按行读取, 以换行符为分隔符

		if err2 == io.EOF {
			fmt.Println("读取完毕 ==>", err2)

			fileStr += str // 最后一行可能没有换行符, 所以要加上
			break
		}

		if err2 != nil {
			fmt.Println("testFn1_3 err2 ==>", err2)
			break
		}

		fileStr += str
	}

	fmt.Println("testFn1_3 读取到的内容 ==>", fileStr) // 读取到的内容 ==>
}

// 使用 bufio 读取文件 (改进字符串拼接方式)(推荐)
func testFn1_4() {
	file, err := os.Open("./hello.txt")

	if err != nil {
		fmt.Println("testFn1_4 err ==>", err)
		return
	}

	defer file.Close()

	// bufio 读取文件
	reader := bufio.NewReader(file)
	var b strings.Builder // 使用 strings.Builder 拼接字符串, 性能更好 (复杂度 O(n))

	for {
		str, err2 := reader.ReadString('\n') // 按行读取, 以换行符为分隔符.
		b.WriteString(str)                   // 即使 err2 非 nil，str 也可能包含数据，所以先写入。

		if err2 == io.EOF {
			fmt.Println("读取完毕 ==>", err2)
			break
		}
		if err2 != nil {
			fmt.Println("testFn1_4 err2 ==>", err2)
			break
		}
	}

	fileStr := b.String() // 获取拼接后的字符串

	fmt.Println("testFn1_4 读取到的内容 ==>", fileStr) // 读取到的内容 ==>
}

// 使用 os.ReadFile 一次性读取整个文件内容
func testFn1_5() {

	// 绝对路径
	strByte, err := os.ReadFile("/Users/sky/Desktop/study_github/self-study-go/base/test_file/hello.txt")

	// 相对路径要注意 运行时命令行的所在目录
	// strByte, err := os.ReadFile("./hello.txt")

	if err != nil {
		fmt.Println("testFn1_5 err ==>", err)
		return
	}

	fmt.Println("testFn1_5 读取到的内容 ==>", string(strByte)) // 读取到的内容 ==>
}

// 流式读取文件 —— 按行读 (不推荐)
func testFn1_6() {
	file, err := os.Open("./hello.txt")

	if err != nil {
		fmt.Println("testFn1_6 err ==>", err)
		return
	}

	defer file.Close()

	buf := bufio.NewReader(file)

	for {
		line, isPrefix, err2 := buf.ReadLine() // 按行读取文件内容

		var full []byte // 用于存储完整的行内容

		fmt.Println("line ==>", string(line), "isPrefix ==>", isPrefix) // isPrefix == true 表示该行内容过长, 需要继续读取下一行

		for isPrefix {
			full = append(full, line...)          // 追加当前行内容
			line, isPrefix, err2 = buf.ReadLine() // 继续读取下一行

			if err2 != nil {
				break
			}
		}

		full = append(full, line...) // 追加最后一行内容

		if len(full) > 0 {
			fmt.Println("完整行内容 ==>", string(full)) // 输出完整行内容
		}

		if err2 == io.EOF {
			fmt.Println("读取完毕 ==>", err2)
			break
		}
		if err2 != nil {
			fmt.Println("testFn1_6 err2 ==>", err2)
			break
		}
	}
}

// 按 分隔符 读取文件内容 (不推荐)
func testFn1_7() {
	file, err := os.Open("./hello4.txt")

	if err != nil {
		fmt.Println("testFn1_7 err ==>", err)
		return
	}

	defer file.Close()

	buf := bufio.NewScanner(file)

	// buf.Split(bufio.ScanWords) // 按照单词分隔符读取文件内容, 也可以使用 bufio.ScanLines 按行读取, bufio.ScanBytes 按字节读取, bufio.ScanRunes 按字符读取
	buf.Split(bufio.ScanLines) // 按行读取文件内容

	var index int = 1 // 行号

	for buf.Scan() {
		fmt.Println("index ==>", index, "内容 ==>", buf.Text())
		index++
	}
}

func main() {
	timeStart := time.Now()

	// 打开文件
	// testFn1_1()

	// testFn1_2() 	// 流式读取文件

	// testFn1_3() // 使用 bufio 读取文件

	// testFn1_4() // 使用 bufio 读取文件 (改进字符串拼接方式)

	// testFn1_5() // 使用 os.ReadFile 一次性读取整个文件内容

	// testFn1_6() // 流式读取文件 —— 按行读

	testFn1_7() // 按 分隔符 读取文件内容

	fmt.Println("程序执行耗时 ==>", time.Since(timeStart)) // 程序执行耗时 ==> 1.019958ms
}
