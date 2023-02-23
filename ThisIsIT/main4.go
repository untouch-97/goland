package main

import "fmt"

// commit
/*
   некий текст
*/

func main() {
	//1  Greet()
	PersonGreet("Baha")
}

// 1   func Greet() {
// 1 	fmt.Println("Hello guys")
func PersonGreet(name string) {
	fmt.Printf("zdarova %s\n", name)
}
