package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func index5_1Fnc(writer http.ResponseWriter, request *http.Request) {
	timeStart := time.Now()

	byteData, err := os.ReadFile("./index.html")

	if err != nil {
		writer.Write([]byte("read file error"))
		return
	}

	fmt.Println("read file success use time ==>", time.Since(timeStart))

	// 字节切片可以直接返回
	writer.Write(byteData)
}

func main() {
	http.HandleFunc("/", index5_1Fnc)

	fmt.Println("server start listen 8080")

	http.ListenAndServe(":8080", nil)
}
