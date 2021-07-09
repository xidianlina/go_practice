package main

import (
	"flag"
	"fmt"
)

var bool_flag = flag.Bool("b", false, "bool类型参数")
var str_flag = flag.String("s", "", "string类型参数")

func main() {
	flag.Parse()
	fmt.Println("-b:", *bool_flag)
	fmt.Println("-s:", *str_flag)
	fmt.Println("其他参数:", flag.Args())
}

/*
------------------------------------
$ go run slice.go
-b: false
-s:
其他参数： []
------------------------------------
$ go run slice.go -b
-b: true
-s:
其他参数： []
------------------------------------
$ go run slice.go -b -s test others
-b: true
-s: test
其他参数： [others]
------------------------------------
$ go run slice.go  -help
Usage of /var/folders/0x/55rm67xj28z5v7z1r5pg11lr0000gp/T/go-build514692984/b001/exe/main:
  -b	bool类型参数
  -s string
    	string类型参数
exit status 2
*/
