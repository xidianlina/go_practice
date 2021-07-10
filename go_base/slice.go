package main

import (
	"fmt"
	"sync"
)

//func main() {
//	foo := make([]int, 5)
//	foo[3] = 42
//	foo[4] = 100
//	for i := range foo {
//		fmt.Printf("%d\t", foo[i])
//	}
//	fmt.Println()
//	fmt.Println("------------------------------")
//	bar := foo[1:4]
//	for j := range bar {
//		fmt.Printf("%d\t", bar[j])
//	}
//	fmt.Println()
//	fmt.Println("------------------------------")
//	bar[1] = 99
//
//	for j := range bar {
//		fmt.Printf("%d\t", bar[j])
//	}
//	fmt.Println()
//	fmt.Println("------------------------------")
//
//	for i := range foo {
//		fmt.Printf("%d\t", foo[i])
//	}
//	fmt.Println()
//	fmt.Println("------------------------------")
//
//	path := []byte("AAAA/BBBBBBBBB")
//	sepIndex := bytes.IndexByte(path, byte('/'))
//	dir1 := path[:sepIndex]
//	dir2 := path[sepIndex+1:]
//	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
//	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB
//	dir1 = append(dir1, "suffix"...)
//	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
//	fmt.Println("dir2 =>", string(dir2))
//}

func main() {
	testMN()
}

// 生产者：消费者=M:N
func testMN() {
	chanInt := make(chan int)
	wg := sync.WaitGroup{}
	wgProducer := sync.WaitGroup{}
	//生产者2个
	wgProducer.Add(2)
	wg.Add(5)

	//生产者1
	go func(ci chan int) {
		defer wg.Done()
		defer wgProducer.Done()

		for i := 0; i < 10; i++ {
			ci <- i
		}
	}(chanInt)
	//生产者2
	go func(ci chan int) {
		defer wg.Done()
		defer wgProducer.Done()

		for i := 10; i < 20; i++ {
			ci <- i
		}
	}(chanInt)

	//消费者1
	for i := 0; i < 2; i++ {
		go func(ci chan int, id int) {
			defer wg.Done()

			for v := range ci {
				fmt.Printf("receive from consumer %d, %d\n", id, v)
			}
		}(chanInt, i)
	}
	//消费者2
	go func() {
		defer wg.Done()
		wgProducer.Wait()
		close(chanInt)
	}()

	wg.Wait()
}
