package main

import (
	"fmt"
	"time"
)

// Switch без условия это тоже самое, что и switch true.
// Эта конструкция может быть использована как более ясный способ записи длинной цепочки if-then-else.
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("good afternoon.")
	default:
		fmt.Println("good evening.")
	}
}
