package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// atomic 原子性操作 可以替代锁

var (
	n1       int64
	n2       int64 = 0
	lock15_1       = sync.Mutex{}
)

func fn16_1() {
	for i := 0; i < 10000; i++ {
		lock15_1.Lock() // 锁控制 并发

		n1++

		lock15_1.Unlock()
	}

	fmt.Println("fn16_1 n1 ==>", n1)
}

func fn16_2() {
	for range 10000 {
		atomic.AddInt64(&n2, 1) // 使用原子操作来保证并发安全
	}

	fmt.Println("fn16_2 n2 ==>", n2)
}

func main() {
	wg1 := sync.WaitGroup{}

	wg1.Add(1)
	go func() {
		defer wg1.Done()

		fn16_1()
	}()

	wg1.Add(1)
	go func() {
		defer wg1.Done()

		fn16_2()
	}()

	wg1.Add(1)
	go func() {
		defer wg1.Done()

		fn16_2()
	}()

	wg1.Wait()

	fmt.Println("main n1 ==>", n1) // main n1 ==> 10000
	fmt.Println("main n2 ==>", n2) // main n2 ==> 20000
}
