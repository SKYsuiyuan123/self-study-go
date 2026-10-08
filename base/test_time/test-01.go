package main

import (
	"fmt"
	"time"
)

// time 包

func main() {
	// 1. 获取时间 Time

	now1 := time.Now() // 返回的是一个 Time 类型的值，表示当前的本地时间

	fmt.Println("now1 ==>", now1) // 2026-10-09 00:06:15.775954 +0800 CST m=+0.000217876

	year1 := now1.Year()     // 获取年份
	month1 := now1.Month()   // 获取月份
	day1 := now1.Day()       // 获取日期
	hour1 := now1.Hour()     // 获取小时
	minute1 := now1.Minute() // 获取分钟
	second1 := now1.Second() // 获取秒数

	fmt.Printf("year1: %d, month1: %d, day1: %d \n", year1, month1, day1)         // year1: 2026, month1: 10, day1: 9
	fmt.Printf("hour1: %d, minute1: %d, second1: %d \n", hour1, minute1, second1) // hour1: 0, minute1: 9, second1: 10

	fmt.Printf("%d-%d-%d %d:%d:%d \n", year1, month1, day1, hour1, minute1, second1) // 2026-10-9 0:10:16

	// 整数宽度不够 补0
	fmt.Printf("%04d-%02d-%02d %02d:%02d:%02d \n", year1, month1, day1, hour1, minute1, second1) // 2026-10-09 00:11:58

	fmt.Println("----------------分割线---111-----------------")

	// 2. 格式化 Time

	// 03: 是 12小时制，15: 24小时制。

	time1 := now1.Format("2006-01-02 15:04:05") // 格式化时间，返回字符串 (常用)
	time2 := now1.Format("2006/01/02 03:04:05")
	time3 := now1.Format("2006/01/02 03:04:05 PM") // 12小时制

	fmt.Println("time1 ==>", time1) // time1 ==> 2026-10-09 00:15:36
	fmt.Println("time2 ==>", time2) // time2 ==> 2026/10/09 12:15:36
	fmt.Println("time3 ==>", time3) // time3 ==> 2026/10/09 12:16:10 AM

	time4 := now1.Format(time.RFC3339)
	fmt.Println("time4 ==>", time4) // time4 ==> 2026-10-09T00:18:07+08:00

	time5 := now1.Format("15:04")
	fmt.Println("time5 ==>", time5) // time5 ==> 00:18

	fmt.Println("----------------分割线---222-----------------")

	// 3. 获取当前时间戳

	unixTime1 := now1.Unix()                // 获取当前时间戳，单位为秒
	fmt.Println("unixTime1 ==>", unixTime1) // unixTime1 ==> 1791476393

	unixTime2 := now1.UnixMilli()           // 获取当前时间戳，单位为毫秒
	fmt.Println("unixTime2 ==>", unixTime2) // unixTime2 ==> 1791476562360

	// 4. 格式化时间戳

	time6 := time.Unix(unixTime1, 0).Format("2006-01-02 15:04:05") // 将时间戳格式化为字符串
	fmt.Println("time6 ==>", time6)                                // time6 ==> 2026-10-09 00:21:26

	time7 := time.UnixMilli(unixTime2).Format("2006-01-02 15:04:05") // 将时间戳格式化为字符串
	fmt.Println("time7 ==>", time7)                                  // time7 ==> 2026-10-09 00:23:44

	fmt.Println("----------------分割线---333-----------------")

	// 5. 解析 时间字符串

	var timeStr = "2020-04-06 15:38:06"

	// `loc = LoadLocation("Local")`: 表示使用本地时区解析时间字符串 (重新读取，拿到 新时区)
	// `time.Local`: 是程序初始化时调用 `time.Now().Location()` 获取的本地时区，之后不会再更新。

	// 运行时操作系统的时区发生了变化，`time.Local` 不会更新，而 `LoadLocation("Local")` 会重新读取系统时区。
	// 业务优先直接使用 `time.Local`。

	// 拓展：如果要指定东八时区，不要用 Local，用 `time.FixedZone("CST",8*3600)`，这样不受系统时区影响。
	// 有些容器没有时区文件，LoadLocation("Local") 会报错，FixedZone 不会。

	loc, err := time.LoadLocation("Local")      // 重新读取系统时区
	fmt.Printf("loc: %v, err: %v \n", loc, err) // loc: Local, err: <nil>

	time10, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc) // 解析时间字符串，返回 Time 类型的值
	fmt.Println("time10 ==>", time10)                                      // time10 ==> 2020-04-06 15:38:06 +0800 CST

	time10_2, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local) // 解析时间字符串，返回 Time 类型的值
	fmt.Println("time10_2 ==>", time10_2)                                           // time10_2 ==> 2020-04-06 15:38:06 +0800 CST

	time10_3, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.FixedZone("CST", 8*3600)) // 解析时间字符串，返回 Time 类型的值
	fmt.Println("time10_3 ==>", time10_3)                                                              // time10_3 ==> 2020-04-06 15:38:06 +0800 CST

	time11_3 := time10_3.UnixMilli()
	fmt.Println("time11_3 ==>", time11_3) // time11_3 ==> 1586158686000

	fmt.Println("----------------分割线---444-----------------")

	// 6. 时间间隔 Duration

	fmt.Println("time.Millisecond ==>", time.Millisecond) // time.Millisecond ==> 1ms
	fmt.Println("time.Second ==>", time.Second)           // time.Second ==> 1s
	fmt.Println("time.Minute ==>", time.Minute)           // time.Minute ==> 1m0s
	fmt.Println("time.Hour ==>", time.Hour)               // time.Hour ==> 1h0m0s

	fmt.Println("----------------分割线---555-----------------")

	// 7. 时间间隔的加减

	now2 := time.Now()

	time15 := now2.Add(time.Hour * 2) // 两小时后的时间
	fmt.Println("time15 ==>", time15) // time15 ==> 2026-10-09 04:26:20.830199 +0800 CST m=+7200.000267751

	fmt.Println("----------------分割线---666-----------------")

	// 8. 时间间隔的计算

	now3 := time.Now()

	time.Sleep(time.Second * 3)

	time20 := time.Now()

	// 8.1 计算 t1 到 t2 的时间间隔
	diff1 := time20.Sub(now3)       // 计算时间间隔，返回 Duration 类型的值
	fmt.Println("diff1 ==>", diff1) // diff1 ==> 3.00101675s

	// 8.2
	d1 := time.Duration(time.Second * 2) // 2秒的时间间隔
	fmt.Println("d1 ==>", d1)            // d1 ==> 2s

	time21 := time20.Add(d1) // 在 t2 的基础上加上 d1 的时间间隔，返回 Time 类型的值

	fmt.Println("time20 ==>", time20) // time20 ==> 2026-10-09 02:31:28.159644 +0800 CST m=+3.001479376
	fmt.Println("time21 ==>", time21) // time21 ==> 2026-10-09 02:31:11.254611 +0800 CST m=+5.001540668

	fmt.Println("----------------分割线---777-----------------")

	// 9. 定时器

	ticker1 := time.NewTicker(time.Second * 1) // 创建一个定时器，每隔 1 秒触发一次

	n := 0

	for t := range ticker1.C {
		fmt.Println("ticker.c ==>", t)

		// ticker.c ==> 2026-10-09 02:37:29.962979542 +0800 CST m=+4.001773001
		// ticker.c ==> 2026-10-09 02:37:30.96297175 +0800 CST m=+5.001770126
		// ticker.c ==> 2026-10-09 02:37:31.962969458 +0800 CST m=+6.001772584
		// ticker.c ==> 2026-10-09 02:37:32.962964333 +0800 CST m=+7.001772626
		// ticker.c ==> 2026-10-09 02:37:33.962957875 +0800 CST m=+8.001771043

		n++

		if n >= 5 {
			ticker1.Stop() // 停止定时器
			break
		}
	}

	fmt.Println("----------------分割线---888-----------------")

	// 10. 定时器休眠

	fmt.Println("Test time.Sleep start ...")

	time.Sleep(time.Second * 3) // 休眠 3 秒

	fmt.Println("Test time.Sleep end ...")
}
