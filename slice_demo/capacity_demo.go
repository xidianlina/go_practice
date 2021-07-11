package main

import "fmt"

func main6() {
	my_slice := []int{11, 22, 33, 44, 55}
	new_slice := append(my_slice, 66)

	my_slice[3] = 444 // 修改旧的底层数组

	fmt.Println(my_slice)  // [11 22 33 444 55]
	fmt.Println(new_slice) // [11 22 33 44 55 66]

	fmt.Println(len(my_slice), ":", cap(my_slice))   // 5:5
	fmt.Println(len(new_slice), ":", cap(new_slice)) // 6:10
}
