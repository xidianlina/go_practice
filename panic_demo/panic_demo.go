package main

import "fmt"

/*
	首先顺序执行，会先将第一个defer延迟函数“入栈”，然后输出“bbbbbbb"，”cccccccc”，此时使用panic来触发一次宕机，
	panic接受一个任意类型的参数，会将该字符串输出，用作提示信息，之后的代码不再执行，所以后面的dddddd不会输出，
	而且第二个defer延迟函数也不会“入栈”，因为panic之后的代码不会继续执行，程序现在只会运行已经“入栈”的defer延迟函数，
	输出aaaaaa，在最后，会输出此次触发宕机的一些信息，所以执行结果如下：
	bbbbbb
	cccccc
	aaaaaa
	panic: hahahaha

	goroutine 1 [running]:
	main.main()
        /Users/lina/go/src/go_practice/panic_demo/panic_demo.go:11 +0x10b

	Process finished with exit code 2
*/
func main() {
	defer func() {
		fmt.Println("aaaaaa")
	}()
	fmt.Println("bbbbbb")
	fmt.Println("cccccc")
	panic("hahahaha")
	fmt.Println("ddddd")
	defer func() {
		fmt.Println("eeeeeeee")
	}()
}
