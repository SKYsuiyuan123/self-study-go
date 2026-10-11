package main

import (
	"fmt"
	"time"
)

// goroutine select 多路复用

// 在某些场景下，我们需要同时监听多个管道的读写操作，这时就可以使用 select 来实现多路复用。

func main() {
	// 1. 定义一个 int 管道

	intChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 2. 定义一个 string 管道

	strChan := make(chan string, 5)

	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 3. 使用 channel 来获取 channel 里面的数据的时候，不需要关闭 channel
	// (除非你想要通知接收方没有数据了) (关闭 channel 之后，接收方会收到一个零值)
	// (关闭后 取数据时会死循环)

	// close(strChan)

	// 4. 使用 select 来实现多路复用
	for {
		select {
		case intVal := <-intChan:
			fmt.Println("intVal ==>", intVal)
			time.Sleep(time.Millisecond * 50)

		case strVal, ok := <-strChan:
			fmt.Printf("strVal: %#v, type: %T, ok: %v \n", strVal, strVal, ok)
			time.Sleep(time.Millisecond * 30)

		default:
			fmt.Println("没有数据了，退出循环")
			return // 必须用 return 来终止程序，否则无限循环执行 default 内容。
		}
	}
}
