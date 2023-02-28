package main

import (
	"fmt"
	"time"
)

// Условия блоков case проверяются сверху вниз до тех пор, пока одно из них не будет успешным. Например:
//
// switch i {
// case 0:
// case f():
// }
// В этом примере f не вызывается если i==0.
func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
