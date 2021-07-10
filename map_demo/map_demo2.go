package main

import "fmt"

/*
	运行结果是：
	m2 outer address 0x0, m=map[]
	inner: map[], 0x0
	inner: map[a:11], 0xc0000681e0
	post m2 outer address 0x0, m=map[]
*/
func main() {
	var m2 map[string]string //未初始化
	fmt.Printf("m2 outer address %p, m=%v \n", m2, m2)
	passMapNotInit(m2)
	fmt.Printf("post m2 outer address %p, m=%v \n", m2, m2)
}

func passMapNotInit(m map[string]string) {
	fmt.Printf("inner: %v, %p\n", m, m)
	m = make(map[string]string, 0)
	m["a"] = "11"
	fmt.Printf("inner: %v, %p\n", m, m)
}
