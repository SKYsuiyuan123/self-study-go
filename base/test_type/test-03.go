package main

import "fmt"

// 使用自定义类型 来处理 错误码 和 错误信息

type httpCode int

// 定义错误码
const (
	successCode      httpCode = 0    // 成功
	serviceErrorCode httpCode = 1001 // 服务端错误
	networkErrorCode httpCode = 1002 // 网络错误
)

func (c httpCode) getMessage1() string {
	switch c {
	case successCode:
		return "成功"
	case serviceErrorCode:
		return "服务端错误"
	case networkErrorCode:
		return "网络错误"
	default:
		return "未知错误"
	}
}

// 封装

func (c httpCode) ok() (code httpCode, message string) {
	return c, c.getMessage1()
}

func webServer(name string) (code httpCode, message string) {
	if name == "1" {
		return serviceErrorCode.ok() // 返回服务端错误码和对应的错误信息，代码更简洁。
	}

	if name == "2" {
		return networkErrorCode.ok()
	}

	return successCode.ok() // 封装的更好

	// 或者
	// return successCode, successCode.getMessage1()
}

func main() {
	res1Code, res1Msg := webServer("1")
	fmt.Println("res1Code:", res1Code, "res1Msg:", res1Msg) // res1Code: 1001 res1Msg: 服务端错误

	res2Code, res2Msg := webServer("2")
	fmt.Println("res2Code:", res2Code, "res2Msg:", res2Msg) // res2Code: 1002 res2Msg: 网络错误

	res3Code, res3Msg := webServer("3")
	fmt.Println("res3Code:", res3Code, "res3Msg:", res3Msg) // res3Code: 0 res3Msg: 成功
}
