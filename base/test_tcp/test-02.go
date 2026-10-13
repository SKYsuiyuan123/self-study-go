package main

import (
	"fmt"
	"io"
	"net"
)

// web 编程 —— TCP 客户端

func main() {
	connect, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println("err ==>", err)
		return
	}

	// connect.Write([]byte("你好，我是客户端"))

	// 接收服务端请求
	for {
		var byteData = make([]byte, 1024)

		n, err2 := connect.Read(byteData)

		if err2 != nil || err2 == io.EOF {
			fmt.Println("err2 ==>", err2)
			break
		}

		fmt.Println("Server Send: ", string(byteData[:n])) // hello world
	}

	for {
		var text string
		fmt.Printf("请输入要发送的内容：")

		fmt.Scan(&text)

		if text == "exit" {
			connect.Close()
			break
		}

		connect.Write([]byte(text))
	}
}
