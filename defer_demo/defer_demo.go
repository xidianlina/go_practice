package main

import "fmt"

func main() {
	//defer_test1()
	//defer_test2()
	res := defer_test3()
	fmt.Println(res)
}

//3 2 1 0
func defer_test1() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i, "\t")
	}
}

func defer_test2() {
	i := 0
	defer fmt.Println(i) //输出0，因为i此时就是0
	i++
	defer fmt.Println(i) //输出1，因为i此时就是1
	return
}

/*
	返回值为 2
	defer是在return调用之后才执行的。 但defer代码块的作用域仍然在函数之内，因此defer仍然可以读取函数内的变量。
	当执行return 1 之后，i的值就是1. 此时，defer代码块开始执行，对i进行自增操作。 因此输出2.
*/
func defer_test3() (i int) {
	defer func() { i++ }()
	return 1
}
