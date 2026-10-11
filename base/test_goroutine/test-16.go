package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁

// 读写锁 可以让多个 读操作并发 同时读取，但是对于写操作来说，必须是独占的，完全互斥。
// 也就是说，当一个 goroutine 在写数据的时候，其他 goroutine 不能读也不能写。

// 补充：已有 goroutine 在等待写锁时，新来的读锁通常也会排队，
//      以避免写锁饥饿。RWMutex 不可重入，禁止拷贝。

// 读写锁由结构体类型 sync.RWMutex 提供，RWMutex 结构体中有两个方法：

// 一组是对写操作的锁定和解锁。简称 "写锁定" 和 "写解锁"。
// func (*RWMutex) Lock() // 写锁定
// func (*RWMutex) Unlock() // 写解锁

// 另一组表示对读操作的锁定和解锁。简称 "读锁定" 和 "读解锁"。
// func (*RWMutex) RLock() // 读锁定
// func (*RWMutex) RUnlock() // 读解锁

var wg16_1 sync.WaitGroup
var mutex16_1 sync.RWMutex

func fnWrite16_1() {
	defer wg16_1.Done()

	// 写锁定 (互斥)
	mutex16_1.Lock()

	time.Sleep(time.Second * 2)

	fmt.Println("写操作完成")

	mutex16_1.Unlock() // 写解锁
}

func fnRead16_2() {
	defer wg16_1.Done()

	// 读锁定 (共享)
	mutex16_1.RLock()

	time.Sleep(time.Second * 2)

	fmt.Println("执行读操作 ------")

	mutex16_1.RUnlock() // 读解锁
}

func main() {
	// 读写互斥锁的使用

	// 1.开启 10个 协程，执行写操作。

	for range 10 {
		wg16_1.Add(1)

		go fnWrite16_1()
	}

	// 2. 开启 10个 协程，执行读操作。
	for range 10 {
		wg16_1.Add(1)

		go fnRead16_2()
	}

	wg16_1.Wait()

	fmt.Println("主线程 操作完成。。。")
}
