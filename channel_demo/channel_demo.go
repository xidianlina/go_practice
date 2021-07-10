package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		场景1：当需要不断从channel读取数据时
		原理：使用for-range读取channel，这样既安全又便利，当channel关闭时，for循环会自动退出，
			 无需主动监测channel是否关闭，可以防止读取已经关闭的channel，造成读到数据为通道所存储的数据类型的零值。
	*/
	c := make(chan int, 3)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		for x := range c {
			fmt.Println(x)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 4; i > 0; i-- {
			c <- i
		}
		close(c)
		fmt.Println(c == nil)
	}()

	wg.Wait()

	/*
		场景2：读channel，但不确定channel是否关闭时
		原理：读已关闭的channel会造成panic，如果不确定channel，需要使用ok进行检测。ok的结果和含义：
		     true：读到数据，并且通道没有关闭。
		     false：通道关闭，无数据读到。
	*/
	c2 := make(chan bool, 3)
	close(c2)
	if v, ok := <-c2; ok {
		fmt.Println(v)
	} else {
		fmt.Println(ok)
	}

	/*
		场景3：需要对多个通道进行同时处理，但只处理最先发生的channel时
		原理：select可以同时监控多个通道的情况，只处理未阻塞的case。
			 当通道为nil时，对应的case永远为阻塞，无论读写。特殊关注：普通情况下，对nil的通道写操作是要panic的。
	*/
	var wag sync.WaitGroup
	ch := make(chan int, 2)
	dh := make(chan string, 2)
	wag.Add(1)

	go func() {
		defer wag.Done()
		dh <- "joker"
		ch <- 10000
	}()

	select {
	case x := <-ch:
		fmt.Println(x)
	case y := <-dh:
		fmt.Println(y)
	}

	wag.Wait()
}

/*
	场景4：需要超时控制的操作
	原理：使用select和time.After，看操作和定时器哪个先返回，处理先完成的，就达到了超时控制的效果
*/
func doWithTimeOut(timeout time.Duration) (int, error) {
	select {
	case ret := <-do():
		return ret, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}

func do() <-chan int {
	outCh := make(chan int)
	go func() {
		// do work
	}()
	return outCh
}

/*
	场景5：并不希望在channel的读写上浪费时间
	原理：是为操作加上超时的扩展，这里的操作是channel的读或写。使用time实现channel无阻塞读写
*/
func unBlockRead(ch chan int) (x int, err error) {
	select {
	case x = <-ch:
		return x, nil
	case <-time.After(time.Microsecond):
		return 0, errors.New("read time out")
	}
}

func unBlockWrite(ch chan int, x int) (err error) {
	select {
	case ch <- x:
		return nil
	case <-time.After(time.Microsecond):
		return errors.New("read time out")
	}
}

/*
	场景6：使用channel传递信号，而不是传递数据时
	原理：没数据需要传递时，传递空struct。使用chan struct{}作为信号channel
*/
type Handler struct {
	stopCh chan struct{}
}

/*
	场景7：使用channel传递结构体数据时，传递结构体的指针而非结构体
	原理：channel本质上传递的是数据的拷贝，拷贝的数据越小传输效率越高，传递结构体指针，比传递结构体更高效
*/
