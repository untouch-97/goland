package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4}
	fmt.Printf("Before len:%v cap:%v %v \n", len(list), cap(list), list)
	handle(list[:1])
	fmt.Printf("After len:%v cap:%v %v \n", len(list), cap(list), list)
}

func handle(list []int) {
	//list[1] = 123
	list = append(list, 5)
	fmt.Printf("Append len:%v cap:%v %v \n", len(list), cap(list), list)
}
