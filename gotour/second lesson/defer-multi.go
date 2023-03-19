package main

import "fmt"

// Накопление отложенных вызовов
// Отложенные вызовы функций сохраняются в стеке. Когда происходит возврат из функции, ее отложенные вызовы выполняются в порядке "последним-пришел-первым-вышел" (LIFO, last-in-first-out).
func main() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
