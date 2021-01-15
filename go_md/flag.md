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

