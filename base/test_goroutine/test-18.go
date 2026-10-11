package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.Map

// 线程安全的 map

func main() {
	var maps = sync.Map{}

	go func() {
		for {
			maps.Store("key", "value") // 底层是使用了读写锁来保证线程安全
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)

			val, ok := maps.Load("key")
			fmt.Printf("val: %v, ok: %v\n", val, ok) // 每隔两秒打印: val: value, ok: true
		}
	}()

	// 让主线程不退出
	select {}
}
