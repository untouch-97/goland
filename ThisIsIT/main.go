package main

import (
	"fmt"
)

func main() {
	name := "vitya"
	//fmt.Println("Please enter your name.")
	//var name string
	//fmt.Scanln(&name)
	hello := fmt.Sprintf("hello %s", name)
	fmt.Println(hello)
	fmt.Printf("nomer1: %T nomer2: %v\n", hello, hello)
}
