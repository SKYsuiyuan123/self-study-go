package main

import "fmt"

// interface —— 多态

// 定义属性

type apiPay struct { // 支付宝支付的结构体
	appId string
}

// 通过结构体与函数进行绑定
func (a apiPay) pay(money float64) (string, float64) {
	fmt.Printf("支付宝支付成功, money: %.2f, appId: %s \n", money, a.appId)
	return "支付宝支付成功", money
}

type weChatPay struct { // 微信支付的结构体
	wxId string
}

func (w weChatPay) pay(money float64) (string, float64) {
	fmt.Printf("微信支付成功, money: %.2f, wxId: %s \n", money, w.wxId)
	return "微信支付成功", money
}

// 定义支付接口

type payment7_1 interface {
	pay(money float64) (string, float64)
}

func createOrder(p payment7_1, money float64) { // p 是一个接口类型的参数，传入的参数必须实现了 payment7_1 接口中的所有方法。
	str, num := p.pay(money)
	fmt.Printf("支付结果: %s, 支付金额: %.2f \n", str, num)
}

func main() {
	// 1. 创建结构体

	// 1.1 创建支付宝支付对象
	aliPay1 := apiPay{
		appId: "zfb_123",
	}

	aliPay1.pay(10.2) // 支付宝支付成功, money: 10.20, appId: zfb_123

	// 1.2 创建微信支付对象
	wxPay1 := weChatPay{
		wxId: "wx_456",
	}

	str, money := wxPay1.pay(20.0) // 微信支付成功, money: 20.00, wxId: wx_456

	fmt.Printf("支付结果: %s, 支付金额: %.2f \n", str, money) // 支付结果: 微信支付成功, 支付金额: 20.00

	fmt.Println("----------------分割线---111-----------------")

	// 2. 创建订单

	// 2.1 使用 支付对象 创建订单

	createOrder(aliPay1, 66.66)
	createOrder(wxPay1, 77.77)

	// 打印：
	// 支付宝支付成功, money: 66.66, appId: zfb_123
	// 支付结果: 支付宝支付成功, 支付金额: 66.66
	// 微信支付成功, money: 77.77, wxId: wx_456
	// 支付结果: 微信支付成功, 支付金额: 77.77

	// 2.2 使用匿名结构体创建订单
	createOrder(apiPay{appId: "zfb_789"}, 88.88)
	createOrder(weChatPay{wxId: "wx_987"}, 99.99)

	// 打印：
	// 支付宝支付成功, money: 88.88, appId: zfb_789
	// 支付结果: 支付宝支付成功, 支付金额: 88.88
	// 微信支付成功, money: 99.99, wxId: wx_987
	// 支付结果: 微信支付成功, 支付金额: 99.99
}
