package main

import "fmt"

// goroutine channel

// go 的并发模型是 CSP（Communicating Sequential Processes）提倡通过通信来共享内存，而不是通过共享内存来通信。
// go 语言的并发模型是基于协程（goroutine）和通道（channel）的。

// channel 是一种特殊的类型。管道像一个环形队列，总是遵循 先入先出的规则，保证收发数据的顺序。
// 每一个管道都是一个具体类型的管道，也就是 channel 的元素类型是固定的。channel 的元素类型可以是任意类型，包括指针类型、结构体类型、接口类型等。
// 声明 channel 时，需要指定 channel 的元素类型。

// 管道满了就会阻塞。

// channel 是一种引用类型。
// 声明格式：var cha chan int

// 声明的管道是 nil，不能直接使用，需要使用 make() 函数来分配内存空间。

// 管道有 发送(send）、接收(receive) 和 关闭(close) 三种操作。

// 发送和接收操作都是阻塞的，直到另一端准备好。
// 发送和接收都可以使用 <- 操作符来实现。

func main() {
	// 1. 创建管道

	ch1 := make(chan int, 3) // 创建一个容量为 3 的管道

	// 给管道里面写入数据
	ch1 <- 10
	ch1 <- 20
	ch1 <- 30

	// 2. 获取管道内容

	a1, ok1 := <-ch1 // 从管道中读取数据

	fmt.Printf("从管道中读取的数据: %d, 是否成功读取: %v \n", a1, ok1) // 从管道中读取的数据: 10, 是否成功读取: true

	// 从管道里面去除值 (如果管道中没有值，则会阻塞，直到有值为止) (取出 20)
	<-ch1

	b1 := <-ch1
	fmt.Println("b1 ==>", b1) // b1 ==> 30

	ch1 <- 40
	ch1 <- 50

	// 打印管道的长度和容量
	fmt.Printf("ch1 的值: %v, 长度: %d, 容量: %d \n", ch1, len(ch1), cap(ch1)) // ch1 的值: 0x99b11a2e000, 长度: 2, 容量: 3

	<-ch1
	_, ok2 := <-ch1             // 从管道中读取数据
	fmt.Println("ok2 ==>", ok2) // ok2 ==> true

	// 3. 关闭管道

	close(ch1)

	_, ok3 := <-ch1
	fmt.Println("ok3 ==>", ok3) // ok3 ==> false,只管管道被关闭，并且管道中没有数据了，ok3 就会返回 false

	fmt.Println("----------------分割线---111-----------------")

	// 4. 管道的类型 (引用类型)

	ch2 := make(chan int, 5) // 创建一个容量为 5 的管道

	ch2 <- 30
	ch2 <- 33
	ch2 <- 35

	ch3 := ch2 // 将 ch2 的引用赋值给 ch3, ch2 和 ch3 指向同一个管道

	fmt.Println("ch2 ==>", ch2) // ch2 ==> 0x1eda6b234000
	fmt.Println("ch3 ==>", ch3) // ch3 ==> 0x1eda6b234000

	ch3 <- 50

	<-ch2
	<-ch2
	<-ch3

	c1 := <-ch2
	fmt.Println("c1 ==>", c1) // c1 ==> 50

	fmt.Println("----------------分割线---222-----------------")

	ch5 := make(chan int, 1) // 创建一个容量为 1 的管道

	ch5 <- 1
	// ch5 <- 2 // 管道阻塞 (管道放满了就会阻塞) fatal error: all goroutines are asleep - deadlock!

	// ch6 := make(chan string) // 创建一个容量为 0 的管道 (无缓冲区的管道) 没有指定管道大小，那么就不能写入数据。写入时会立刻报 deadlock.
	ch6 := make(chan string, 2)

	ch6 <- "数据一"
	ch6 <- "数据二"

	m1 := <-ch6
	m2 := <-ch6
	// m3 := <-ch6 // 管道阻塞 (管道已被取空了) fatal error: all goroutines are asleep - deadlock!

	// 在没有使用 协程 的情况下，如果我们的管道数据已经全部取出，再取就会报 deadlock。

	fmt.Println("m1, m2 ==>", m1, m2) // m1, m2 ==> 数据一 数据二

	fmt.Println("----------------分割线---222-----------------")

	// 5. 管道里放入数后，立刻消费掉，就不会出现死锁

	ch7 := make(chan int, 1)

	ch7 <- 1
	m5 := <-ch7

	fmt.Println("m5 ==>", m5) // m5 ==> 1

	ch7 <- 2
	m6 := <-ch7

	fmt.Println("m6 ==>", m6) // m6 ==> 2
}
