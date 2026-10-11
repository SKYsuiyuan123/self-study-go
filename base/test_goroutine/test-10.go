package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用 goroutine + channel 来统计 素数

const maxVal10_1 = 1000000

var wg10_1 sync.WaitGroup

// 存放数据
func puNum(ch chan int) {
	defer wg10_1.Done()
	defer close(ch) // 写入完成后关闭管道

	for i := 2; i < maxVal10_1; i++ {
		ch <- i
	}
}

// 统计素数
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	defer wg10_1.Done()

	for num := range intChan {
		flag := true // 假设 num 是素数

		for i := 2; i*i < num; i++ { // 这里可以优化为 i*i <= num (优化前耗时: 3.252676833s, 优化后耗时: 170.03025ms)
			if num%i == 0 {
				flag = false // num 不是素数
				break
			}
		}

		// num 是素数
		if flag {
			primeChan <- num
		}
	}

	// 如果一个 channel 关闭了就没法再给这个 channel 发送数据了。
	// close(primeChan) // 报错: panic: close of closed channel

	// close channel 的规则 一般是由生产者负责关。(发送到已关闭 和 重复 close 都会触发 panic)

	// 什么时候关闭 primeChan 呢？当所有的生产者都结束了，才可以关闭 primeChan。
	// 每次执行完 primeNum 给 exitChan 发送一个 true，表示这个 primeNum 执行完了。
	// 当 exitChan 数据量等于 开启的 primeNum 的协程数据量时，关闭 primeChan。

	exitChan <- true
}

// 打印素数
func printPrime(primeChan chan int) {
	defer wg10_1.Done()

	for val := range primeChan {
		fmt.Println("素数 val ==>", val)
	}
}

func main() {
	timeStart := time.Now()

	// 计算素数时，开启的协程数量。
	const goroutineNum = 8

	// 1. 存放数字
	intChan := make(chan int, 1000)

	// 2. 存放素数
	primeChan := make(chan int, 1000)

	// 3. 存放 标识 primeChan close
	exitChan := make(chan bool, 8) // 一共开启 8个协程，因此容量需要 16.

	// 4. 存放要统计的数字
	wg10_1.Add(1)
	go puNum(intChan)

	// 5. 统计素数
	for i := 0; i < goroutineNum; i++ {
		wg10_1.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	// 6. 打印素数
	wg10_1.Add(1)
	go printPrime(primeChan)

	// 判断 exitChan 是否存满值 (可以再定义一个方法，或者使用 匿名自执行函数)

	wg10_1.Add(1)

	go func() {
		defer wg10_1.Done()

		for i := 0; i < goroutineNum; i++ {
			<-exitChan
		}

		// exitChan 存满了 goroutineNum 个值，说明所有的 primeNum 都执行完了，可以关闭 primeChan 了。

		// 关闭 primeChan
		close(primeChan)

		// 关闭 exitChan
		close(exitChan)
	}()

	wg10_1.Wait()

	fmt.Println("主线程执行完毕 耗时 ==>", time.Since(timeStart))
}
