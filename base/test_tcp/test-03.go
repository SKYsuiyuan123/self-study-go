package main

import (
	"fmt"
	"net/http"
)

// 网络编程 —— http 服务端

func index3_1(writer http.ResponseWriter, request *http.Request) {
	// writer.Write([]byte("index3_1"))

	fmt.Println("index3_1 request.url ==>", request.URL)

	writer.Write([]byte("Hello I'm is http Server index3_1"))
}

func main() {
	http.HandleFunc("/", index3_1)

	fmt.Println("Go http server is running on port 8082...")

	http.ListenAndServe(":8082", nil)
}
