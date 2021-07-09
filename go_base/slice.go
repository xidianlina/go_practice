package main

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
