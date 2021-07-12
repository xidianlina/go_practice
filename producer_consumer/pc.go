package main

import (
	"fmt"
	"time"
)

func producer4(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i //主线程不能产生死锁,所以此处报错
		fmt.Println("send:", i)
	}
}

//自动消费
func autoConsumer(ch <-chan int) {
	for {
		select {
		case ws := <-ch:
			fmt.Println("fmt print", ws)
		default:
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func main() {
	ch := make(chan int)
	//只生产和消费10条记录
	//持续生产与消费, high起来
	go func() {
		for { //不断生产,一次10条
			producer4(ch)
		}
	}()
	go autoConsumer(ch)

	for {
		//心跳
		time.Sleep(time.Second)
	}

}
