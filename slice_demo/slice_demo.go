package main

import "fmt"

func main2() {
	//1.使用内置的make函数
	slice := make([]string, 5) //只指定长度，则默认容量和长度相等
	/*指定长度和容量，容量不能小于长度。声明一个长度为5、数据类型为string的底层数组，
	  然后从这个底层数组中从前向后取3个元素(即index从0到2)作为slice的结构。*/
	slice = make([]string, 3, 5)
	for k, v := range slice {
		fmt.Println(k, v)
	}

	//2.使用切片字面量
	slice = []string{"dog", "cat", "bear"} //其长度和容量都是3
	slice = []string{99: "0"}              //使用索引声明切片,创建了一个长度为100的切片
	for k, v := range slice {
		fmt.Println(k, v)
	}

	//3.声明时不做任何初始化就会创建一个nil切片
	//var s []int
	//s := *new([]int) //new 产生的是指针，需要用*

	//4.声明空切片
	//s1 = make([]int, 0) //使用make
	//s2 := []int{}        //使用切片字面量
}
