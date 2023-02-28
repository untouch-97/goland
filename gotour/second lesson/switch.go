package main

import (
	"fmt"
	"runtime"
)

// Условия блоков case проверяются сверху вниз до тех пор, пока одно из них не будет успешным. Например:
// switch i {
// case 0:
// case f():
// }
// В этом примере f не вызывается если i==0.
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
