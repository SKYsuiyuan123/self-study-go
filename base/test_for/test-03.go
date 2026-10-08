package main

import "fmt"

// switch

// fallthrough 的关键特性：
// 	1. fallthrough 无条件穿透：它不看下一个 case 的条件，直接执行下一个 case 的语句体。
// 	2. 只能穿透一层：一个 fallthrough 只跳到紧邻的下一个 case。想继续穿，下一个 case 里必须再写 fallthrough。
//  3. 不能用在最后一个 case：最后一个 case 里写 fallthrough 会编译错误 cannot fallthrough final case in switch。
//  4. 不能跨过 default：如果 fallthrough 后面紧跟的是 default，行为要小心（它会穿到 default）。
//  5. fallthrough 和条件无关：这是最容易误解的点——它不会“如果下一个条件也满足才穿”，而是强制穿。

func main() {
	// 1. fallthrough 无条件执行下一个 case 的语句体

	// switch 无表达式 (case 后必须是 布尔(bool) 表达式)

	age1 := 33

	switch {
	case age1 < 24:
		fmt.Println("学习")
	case age1 <= 50: // 隐含 age1 >= 24
		fmt.Println("努力工作")
	case age1 < 60: // 隐含 age1 > 50
		fmt.Println("努力活到领养老金")
	default: // 隐含 age1 >= 60
		fmt.Println("多出门转转")
	}

	// 打印：
	// 	努力工作

	fmt.Println("----------------分割线---111-----------------")

	switch {
	case age1 < 24:
		fmt.Println("学习")
	case age1 >= 24 && age1 <= 50:
		fmt.Println("努力工作")
		fallthrough // 只会穿透 紧邻的一层 case，无法多层穿透。如果想穿透下一层 case 则需要在下一层 case 中加入 fallthrough
	case age1 > 50 && age1 < 60:
		fmt.Println("努力活到领养老金")
		fallthrough // 无条件执行下一个 case 的语句体，会忽略下一条 case 的判断条件。(互斥的区间不需要也不应该用 fallthrough，否则一旦命中一个，后面的会被强行执行。)
	case age1 >= 60:
		fmt.Println("多出门转转")
		// fallthrough // 不能用在最后一个 case，会编译报错。
	}

	// 打印：
	// 	努力工作
	// 	努力活到领养老金
	// 	多出门转转

	// 1.2 fallthrough 设计的本意是处理区间重叠、层层递进的场景 (fallthrough 只适合“满足高档则同时满足低档”这种递进累积的场景。)

	score := 88

	switch {
	case score >= 90:
		fmt.Println("优秀")
		fallthrough
	case score >= 80:
		fmt.Println("良好")
		fallthrough
	case score >= 60:
		fmt.Println("及格")
		fallthrough
	default:
		fmt.Println("分级完成")
	}

	// 打印：
	// 	良好
	// 	及格
	// 	分级完成

	fmt.Println("----------------分割线---222-----------------")

	// 2. switch case 匹配多个值

	day := "星期六"

	switch day {
	case "星期一", "星期二", "星期三", "星期四", "星期五":
		fmt.Println("工作时间")

	case "星期六", "星期日":
		fmt.Println("休息时间")
	}

	// 打印：
	// 	休息时间
}
