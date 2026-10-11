package main

import (
	"fmt"
	"sync"
	"time"
)

// 锁

// 互斥锁
// 互斥锁 是传统的锁机制，保证同一时间只有一个线程可以访问共享资源。Go语言提供了 sync包 中的 Mutex结构体类型 来实现互斥锁。
// sync.Mutex 类型只有两个公开的指针方法：Lock() 和 Unlock()，分别用于加锁和解锁。
// Lock() 方法会阻塞当前线程，直到获取到锁为止；Unlock() 方法会释放锁，使其他线程可以获取锁。

var count = 0
var wg14_1 sync.WaitGroup
var mutex14_1 sync.Mutex

// 或
// var mutex14_2 sync.Mutex = sync.Mutex{}

var m1 = make(map[int]int, 0)

func fn14_1() {
	defer wg14_1.Done()

	// 加锁
	mutex14_1.Lock()

	count++

	fmt.Println("The count is ==>", count)

	time.Sleep(time.Millisecond * 5)

	// 解锁
	mutex14_1.Unlock()
}

func fn14_2(num int) {
	defer wg14_1.Done()

	// 加互斥锁 (因为 不同的协程 同时操作了 m1，所以需要加锁。)
	mutex14_1.Lock()

	sum := 1

	for i := 1; i < num; i++ {
		sum *= i
	}

	m1[num] = sum

	time.Sleep(time.Millisecond * 5)

	mutex14_1.Unlock()
}

func main() {
	// 1. 计数
	for range 20 {
		wg14_1.Add(1)

		go fn14_1()
	}

	// The count is ==> 20

	// 2. 计算阶乘
	for i := 1; i <= 20; i++ {
		wg14_1.Add(1)

		go fn14_2(i)
	}

	wg14_1.Wait()

	fmt.Printf("m1: %#v \n", m1) // m1: map[int]int{1:1, 2:1, 3:2, 4:6, 5:24, 6:120, 7:720, 8:5040, 9:40320, 10:362880, 11:3628800, 12:39916800, 13:479001600, 14:6227020800, 15:87178291200, 16:1307674368000, 17:20922789888000, 18:355687428096000, 19:6402373705728000, 20:121645100408832000}
}
