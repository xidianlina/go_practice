package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func run(task string) {
	fmt.Println(task, "start....")
	time.Sleep(time.Second * 2)

	// 每个goroutine运行完毕后就释放等待组的计数器
	wg.Done()
}

func testWaitGroup() {
	// 需要开启几个goroutine就给等待组的计数器赋值为多少，这里为2
	wg.Add(2)
	for i := 1; i < 3; i++ {
		taskName := "task" + strconv.Itoa(i)
		go run(taskName)
	}

	// 等待，等待所有的任务都释放
	wg.Wait()
	fmt.Println("所有任务结束...")
}

func testChannelSign() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("任务1 结束了...")
				return
			default:
				fmt.Println("任务1 正在运行中...")
				time.Sleep(time.Second * 2)
			}
		}
	}()

	// 运行10s后停止
	time.Sleep(time.Second * 10)
	fmt.Println("需要停止任务1...")
	stop <- true
	time.Sleep(time.Second * 3)
}

func testWithCancel() {
	/*
			context.Background() 返回一个空的 Context，这个空的 Context 一般用于整个 Context 树的根节点。
			然后使用 context.WithCancel(parent) 函数，创建一个可取消的子 Context，然后当作参数传给 goroutine 使用，
		这样就可以使用这个子 Context 跟踪这个 goroutine。
	*/
	ctx, cancel := context.WithCancel(context.Background())
	// 开启goroutine，传入ctx
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("任务1 结束了...")
				return
			default:
				fmt.Println("任务1 正在运行中...")
				time.Sleep(time.Second * 2)
			}
		}
	}(ctx)

	// 运行10s后停止
	time.Sleep(time.Second * 10)
	fmt.Println("需要停止任务1...")
	// 使用context的cancel函数停止goroutine
	cancel()
	// 为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(time.Second * 3)
}

// 使用context控制多个goroutine
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "退出,停止了...")
			return
		default:
			fmt.Println(name, "运行中...")
			time.Sleep(time.Second * 2)
		}
	}
}

func testWatch() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "【任务1】")
	go watch(ctx, "【任务2】")
	go watch(ctx, "【任务3】")

	time.Sleep(time.Second * 10)
	fmt.Println("通知任务停止...")
	cancel()
	time.Sleep(time.Second * 5)
	fmt.Println("真的停止了...")
}

var key string = "name"

// 使用通过context向goroutinue传递值
func watchKey(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), "退出，停止了...")
			return
		default:
			fmt.Println(ctx.Value(key), "运行中...")
			time.Sleep(2 * time.Second)
		}
	}
}

func testValue() {
	ctx, cancel := context.WithCancel(context.Background())
	// 给ctx绑定键值，传递给goroutine
	valueCtx := context.WithValue(ctx, key, "[监控1]")

	// 启动goroutine
	go watchKey(valueCtx)

	time.Sleep(time.Second * 10)
	fmt.Println("该结束了...")
	// 运行结束函数
	cancel()
	time.Sleep(time.Second * 3)
	fmt.Println("真的结束了...")
}

func main() {
	testWaitGroup()
	fmt.Println("---------------------------------------")
	testChannelSign()
	fmt.Println("---------------------------------------")
	testWithCancel()
	fmt.Println("---------------------------------------")
	testWatch()
	fmt.Println("---------------------------------------")
	testValue()
	fmt.Println("---------------------------------------")
}
