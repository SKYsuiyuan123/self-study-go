package main

import "fmt"

// 接口 —— 判断具体的 结构体类型

type usb2_1er interface {
	start()
	stop()
}

type computer2_1 struct{}

func (c computer2_1) work2_1(usb usb2_1er) {
	// 可以判断 usb 的具体类型 (类型断言)
	if _, ok := usb.(phone2_1); ok { // 手机可以多执行一个 stop 方法
		usb.start()
		usb.stop()
	} else {
		usb.start()
	}
}

type phone2_1 struct {
	name string
}

// 手机 要实现 usb 接口的话，必须得实现 usb 接口里的所有方法。

func (p phone2_1) start() {
	fmt.Println(p.name, "手机开始工作")
}

func (p phone2_1) stop() {
	fmt.Println(p.name, "手机停止工作")
}

type cameraType2_1 struct{}

func (c cameraType2_1) start() {
	fmt.Println("相机开始工作")
}

func (c cameraType2_1) stop() {
	fmt.Println("相机停止工作")
}

func (c cameraType2_1) run() {
	fmt.Println("相机拍照")
}

func main() {
	// 1. 创建一个电脑对象
	var c1 = computer2_1{}

	var p1 = phone2_1{
		name: "小米",
	}

	var camera1 = cameraType2_1{}

	c1.work2_1(p1) // p1 和 camera1 都实现了 usb 接口，所以都可以传入 work2_1 方法中。
	// 打印：小米 手机开始工作
	// 打印：小米 手机停止工作

	c1.work2_1(camera1) // 相机开始工作

	fmt.Println("----------------分割线---111-----------------")
}
