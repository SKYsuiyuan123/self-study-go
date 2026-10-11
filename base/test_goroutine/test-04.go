package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用 goroutine 查找素数

// 需求：统计 1 - 1000000 的素数个数，并打印这些素数。
// 素数：就是除了 1 和它本身之外，不能被其他数整除的数。

// 实现方法：
// 1. 传统方法，通过两层 for 循环判断各个数 是不是素数。(效率低 (4m50.472s))
// 2. 使用 并行或者并发的方式，将统计素数的任务分配给多个 goroutine 来执行。(效率高 (4.258725083s), 但是需要注意协程的数量，过多的协程会导致性能下降)
// 3. goroutine 结合 channel 来实现统计素数的个数。

const maxVal = 1000000

// 方法一：for 循环

func test4_1() {
	slices1 := make([]int, 0, maxVal)

	for num := 2; num <= maxVal; num++ {
		var flag = true // 假设 num 是素数

		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false // num 不是素数
			}
		}

		if flag {
			slices1 = append(slices1, num)
		}
	}

	fmt.Printf(`1 - %d 的素数个数为：%d, 分别是: %v`, maxVal, len(slices1), slices1)
}

// 方法二：多协程并行计算

const goroutinePage = 100000

var wg4_1 = sync.WaitGroup{}

func test4_2(n int) {
	defer wg4_1.Done() // 在函数结束时调用 Done()，表示该协程完成。

	slices1 := make([]int, 0, goroutinePage)

	for num := (n-1)*goroutinePage + 1; num < n*goroutinePage; num++ {
		if num == 1 {
			continue
		}

		var flag = true // 假设 num 是素数

		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false // num 不是素数
				break
			}
		}

		if flag {
			slices1 = append(slices1, num)
		}
	}

	fmt.Println((n-1)*goroutinePage+1, " - ", n*goroutinePage, "内的素数有 ==>", slices1)
}

func main() {
	timeStart := time.Now() // 记录开始时间

	// 方法一：单线程计算
	// test4_1()

	// 方法二：多协程计算
	count := maxVal / goroutinePage

	fmt.Println("count ==>", count)

	for i := 1; i <= count; i++ {
		wg4_1.Add(1) // 在开启协程前调用 Add(1)，表示有一个协程需要等待。

		go test4_2(i) // 开启协程
	}

	wg4_1.Wait() // 等待所有协程完成

	fmt.Println("共耗时 ==>", time.Since(timeStart)) // 计算耗时
}
