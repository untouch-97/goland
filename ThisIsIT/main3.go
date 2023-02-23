package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var num1 uint8 = 1
	var num2 uint64 = 1

	fmt.Println(unsafe.Sizeof(num2))
	fmt.Println(unsafe.Sizeof(num1))
}
