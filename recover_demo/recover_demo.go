package main

import "fmt"

func main() {
	do()
}

func do() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from run")
		}
	}()

	run() //直接调用
}

func run() {
	panic("panic")
}
