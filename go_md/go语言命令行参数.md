Go语言-命令行参数（os.Args, flag包）
======

> 大部分Go程序都是没有UI的，运行在纯命令行的模式下，该干什么全靠运行参数。

# 1. os.Args
程序获取运行他时给出的参数，可以通过os包来实现。
```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for idx, args := range os.Args {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}
}
```
运行起来得到的如下：
   ![args](http://github.com/xidianlina/go_practice/raw/master/picture/args.jpg)
> 可以看到，命令行参数包括了程序路径本身，以及通常意义上的参数。
> 程序中os.Args的类型是 []string ，也就是字符串切片。所以可以在for循环的range中遍历，还可以用 len(os.Args) 来获取其数量。

如果不想要输出os.Args的第一个值，也就是可执行文件本身的信息，可以修改上述程序：
```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for idx, args := range os.Args[1:] {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}
}
```

# 输出切片的所有元素，还有更简洁的方式：
```go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for idx, args := range os.Args {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}
	fmt.Println("\n-------------------------------")

	for idx, args := range os.Args[1:] {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}

	fmt.Println("\n-------------------------------")
	fmt.Println(strings.Join(os.Args[1:], "\n"))
}
```

# 2.flag包
flag包相比单纯的通过os.Args切片分析命令行参数，提供了更强的能力。
```go
package main

import (
	"flag"
	"fmt"
)

var b = flag.Bool("b", false, "bool类型参数")
var s = flag.String("s", "", "string类型参数")

func main() {
	flag.Parse()
	fmt.Println("-b:", *b)
	fmt.Println("-s:", *s)
	fmt.Println("其他参数:", flag.Args())
}

/*
------------------------------------
$ go run main.go
-b: false
-s:
其他参数： []
------------------------------------
$ go run main.go -b
-b: true
-s:
其他参数： []
------------------------------------
$ go run main.go -b -s test others
-b: true
-s: test
其他参数： [others]
------------------------------------
$ go run main.go  -help
Usage of /var/folders/0x/55rm67xj28z5v7z1r5pg11lr0000gp/T/go-build514692984/b001/exe/main:
  -b	bool类型参数
  -s string
    	string类型参数
exit status 2
 */
```

## 2.1定义参数
> 使用flag包，首先定义待解析命令行参数，也就是以"-"开头的参数，比如这里的 -b -s -help等。-help不需要特别指定，可以自动处理。

var b = flag.Bool("b", false, "bool类型参数")   
var s = flag.String("s", "", "string类型参数")  
原型:     
func Bool(name string, value bool, usage string) *bool  
func String(name string, value string, usage string) *string

> 通过flag.Bool和flag.String，建立了2个指针b和s，分别指向bool类型和string类型的变量。所以后续要通过 *b 和 *s 使用变量值。
  flag.Bool和flag.String的参数有3个，分别是命令行参数名称，默认值，提示字符串。

   ![flag](http://github.com/xidianlina/go_practice/raw/master/picture/flag.png)

## 2.2 解析参数
> flag使用前，必须首先解析:   
flag.Parse()

## 2.3 使用参数
通过flag方法定义好的参数变量指针，通过间接引用操作即可使用其内容：  
fmt.Println("-b:", *b)  
fmt.Println("-s:", *s)

## 2.4 未解析参数
参数中没有能够按照预定义的参数解析的部分，通过flag.Args()即可获取，是一个字符串切片。

注意:从第一个不能解析的参数开始，后面的所有参数都是无法解析的。即使后面的参数中含有预定义的参数：

# Usage
https://blog.csdn.net/guanchunsheng/article/details/79612153
