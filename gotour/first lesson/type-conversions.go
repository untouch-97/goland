package main

import (
	"fmt"
	"math"
)

// var i int = 42
// var f float64 = float64(i)
// var u uint = uint(f)
// Или проще:
// i := 42
// f := float64(i)
// u := uint(f)
func main() {
	var x, y int = 3, 4
	//Выражение T(v) приводит значение v к типу T.
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
