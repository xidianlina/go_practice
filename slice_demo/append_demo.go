package main

import "fmt"

func main3() {
	my_slice := []int{11, 22, 33, 44, 55}
	new_slice := my_slice[1:3]

	// append()追加一个元素2323，返回新的slice
	app_slice := append(new_slice, 2323)

	for k, v := range new_slice {
		fmt.Println(k, v)
	}
	fmt.Println("--------")
	for k, v := range app_slice {
		fmt.Println(k, v)
	}
}
