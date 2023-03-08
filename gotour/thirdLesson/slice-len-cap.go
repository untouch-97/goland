package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	prntSlice(s)

	s = s[:0]
	prntSlice(s)

	s = s[:4]
	prntSlice(s)

	s = s[2:]
	prntSlice(s)
}
func prntSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
