package main

import (
	"fmt"
	"sync"
	"time"
)

// select 多 channel

var moneyChan3 = make(chan int, 0)
var nameChan3 chan string = make(chan string)

var doneChan = make(chan struct{}) // 声明一个用于关闭的信道

func pay3(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物 -->\n", name)

	time.Sleep(1 * time.Second)

	fmt.Printf("%s <-- 购物结束\n", name)

	moneyChan3 <- money
	nameChan3 <- name

	wait.Done()
}

func main() {
	time1 := time.Now()

	var wait1 sync.WaitGroup

	wait1.Add(3)

	go pay3("小明", 2, &wait1)
	go pay3("小红", 3, &wait1)
	go pay3("小刚", 5, &wait1)

	go func() {
		defer close(moneyChan3)
		defer close(nameChan3)

		defer close(doneChan)

		wait1.Wait()
	}()

	var moneyList []int
	var nameList []string

	var event = func() {
		for {
			select {
			case money := <-moneyChan3:
				moneyList = append(moneyList, money)
			case name := <-nameChan3:
				nameList = append(nameList, name)
			case <-doneChan:
				fmt.Println("结束")

				fmt.Println("所有人的购物结束，总耗时 ==>", time.Since(time1)) // 1.001414792s
				fmt.Println("所有人购物结束，花费为 ==>", moneyList)          // [2 3 5]
				fmt.Println("所有人购物结束，购物的人为 ==>", nameList)         // [小明 小红 小刚]
				return
			}
		}
	}

	event()

	fmt.Println("程序运行结束，总耗时 ==>", time.Since(time1)) // 1.001598209s
}
