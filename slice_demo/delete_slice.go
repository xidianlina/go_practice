package main

import "fmt"

func delete_slice(index int, s []int) []int {
	s1 := s[:index]
	s1 = append(s1, s[index+1:]...)
	return s1
}

func main5() {
	slice := []int{1, 2, 3, 4, 5}
	for k, v := range slice {
		fmt.Println(k, v)
	}
	slice = delete_slice(2, slice)
	fmt.Println("-----------")
	for k, v := range slice {
		fmt.Println(k, v)
	}
}
