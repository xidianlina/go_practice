package main

import "fmt"

type MyFunc func(string) string

func testFunc(s string, myFunc MyFunc) {
	myFunc(s)
}

func protect(unprotected func(...interface{})) func(...interface{}) {
	return func(args ...interface{}) {
		fmt.Println("protected")
		unprotected(args...)
	}
}

func main() {
	testFunc("happy", func(s string) string {
		fmt.Println(s)
		return s
	})

	protect(func(args ...interface{}) {
		for k, v := range args {
			fmt.Println(k, v)
		}
	})([]string{"a", "b", "c"})
}
