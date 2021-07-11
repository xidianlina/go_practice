package main

import "fmt"

func main() {
	s1 := []int{11, 22, 33, 44}
	foo(s1)
	fmt.Println(s1[1]) // 输出：23

	slice := []int{1, 2, 3, 4, 5}
	sliceModify(slice)
	fmt.Println(cap(slice))
	fmt.Println(slice) // [1 2 3 4 5]
	fmt.Printf("%p", slice)
}

func foo(s []int) {
	for index, _ := range s {
		s[index] += 1
	}
}

func sliceModify(slice []int) {
	slice = append(slice, 6)
	fmt.Printf("%p", slice)
}
