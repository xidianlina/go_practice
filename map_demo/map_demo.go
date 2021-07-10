package main

import "fmt"

/*
	运行结果是：
	m outer address 0xc000068180, m=map[1:0]
	m inner address 0xc000068180
	post m inner address 0xc000068180
	post m outer address 0xc000068180, m=map[1:0 11111111:11111111]
*/
func main() {
	m := make(map[string]string)
	m["1"] = "0"
	fmt.Printf("m outer address %p, m=%v \n", m, m)
	passMap(m)
	fmt.Printf("post m outer address %p, m=%v \n", m, m)
}

func passMap(m map[string]string) {
	fmt.Printf("m inner address %p \n", m)
	m["11111111"] = "11111111"
	fmt.Printf("post m inner address %p \n", m)
}
