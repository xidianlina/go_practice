package main

import (
	"fmt"
	"time"
)

func producer3(intChan chan int) {
	for i := 0; i < cap(intChan); i++ {
		fmt.Println("生产者：", i)
		intChan <- i
	}
	// 写完后关闭掉
	close(intChan)
}

func consumer3(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if ok {
			fmt.Println("消费者：", v)
		} else { // 读完了
			break
		}
		time.Sleep(time.Second)
	}
	exitChan <- true
	close(exitChan)
}

func main4() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	go producer3(intChan)
	go consumer3(intChan, exitChan)

	for {
		select {
		case _, ok := <-exitChan:
			if ok {
				fmt.Println("执行完毕！")
				// break // break只是跳出select循环，可配合lable跳出
				return
			}
		default:
			fmt.Println("读不到，执行其他的！")
			time.Sleep(time.Second) // 此处添加Sleep才会看到效果，否则打印太多了找不到输出
		}
	}
	fmt.Println("主线程结束！")
}
