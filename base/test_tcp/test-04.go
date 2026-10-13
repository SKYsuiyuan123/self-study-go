package main

import (
	"fmt"
	"io"
	"net/http"
)

// 网络编程 —— http 客户端

func main() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Println("err ==>", err)
		return
	}

	byteData, _ := io.ReadAll(response.Body)

	fmt.Println("byteData ==>", string(byteData))
}
