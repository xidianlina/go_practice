package main

import (
	"fmt"
	"time"
)

func producer2(intChan chan int) {
	for i := 0; i < cap(intChan); i++ {
		fmt.Println("生产者：", i)
		intChan <- i
	}

	// 写完后关闭掉
	close(intChan)
}

func consumer2(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if ok {
			fmt.Println("消费者：", v)
		}

		time.Sleep(time.Second)

		exitChan <- true
		close(exitChan)
	}
}

func main3() {
	intChan := make(chan int, 10)  // “生产者”和“消费者”之间互相通信的桥梁，这里假设生产的元素就是int类型的数字
	exitChan := make(chan bool, 1) // 退出的channel，因为仅做为一个标志所以空间为一个元素就够了
	go producer2(intChan)
	go consumer2(intChan, exitChan)

	// 1) for循环的等待判断
	// for {
	// 	_, ok := <-exitChan
	// 	if !ok {
	// 		break
	// 	}
	// }
	// 2) for range 阻塞，等待关闭close channel
	for ok := range exitChan {
		fmt.Println(ok)
	}
	fmt.Println("主线程结束！")
}
