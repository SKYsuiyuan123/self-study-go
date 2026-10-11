package main

import (
	"errors"
	"fmt"
	"time"
)

// select + time.After 实现超时控制

// 模拟超时请求
func fn13_1() <-chan string {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 3) // 模拟耗时操作
		ch <- "fn13_1 数据"
	}()

	return ch
}

func fn13_2() (data any, err error) {
	select {
	case data := <-fn13_1():
		return data, nil

	case <-time.After(time.Second * 2): // 超时控制
		return nil, errors.New("fn13_2 超时")
	}
}

func main() {
	// 测试超时请求
	var data, err = fn13_2()

	if err != nil {
		fmt.Println("fn13_2 err ==>", err) // fn13_2 err ==> fn13_2 超时
	} else {
		fmt.Printf("fn13_2 data: %#v, type: %T \n", data, data) // fn13_2 data: "fn13_1 数据", type: string
	}
}
