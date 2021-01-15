package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for idx, args := range os.Args {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}
	fmt.Println("\n-------------------------------")

	for idx, args := range os.Args[1:] {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}

	fmt.Println("\n-------------------------------")
	fmt.Println(strings.Join(os.Args[1:], "\n"))
}
