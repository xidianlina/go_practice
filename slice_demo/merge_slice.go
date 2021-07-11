package main

import "fmt"

func main7() {
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	s3 := append(s1, s2...)
	fmt.Println(s3) // [1 2 3 4]

	s4 := []int{7, 8}
	s5 := []int{5, 6}
	s := append(s1, append(s2, append(s4, s5...)...)...)
	fmt.Println(s) // [1 2 3 4 7 8 5 6]
}
