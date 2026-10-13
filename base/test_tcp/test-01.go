package main

import (
	"fmt"
	"net"
	"time"
)

// web 编程 —— TCP 服务端

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	listen, err := net.ListenTCP("tcp", addr)

	if err != nil {
		fmt.Println("err ==>", err)
		return
	}

	fmt.Println("Tcp server bind addr ==>", addr.String())

	for {
		connect, err2 := listen.AcceptTCP()

		if err2 != nil {
			fmt.Println("err2 ==>", err2)
			break
		}

		fmt.Println("connect ==>", connect.RemoteAddr().String()) // connect ==> 127.0.0.1:57677

		connect.Write([]byte("hello world"))

		// 两秒之后关闭连接
		time.Sleep(time.Second * 2)

		connect.Close()
	}
}
