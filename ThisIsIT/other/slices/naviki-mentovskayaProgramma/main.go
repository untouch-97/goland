package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	intSlice := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		intSlice = append(intSlice, i)
	}
	fmt.Println("---------------------------")
	fmt.Printf("| intSlice %v, len: %d, cap: %d \n", intSlice, len(intSlice), cap(intSlice))

	fmt.Println("---------------------------")
	cut := intSlice[2:4]
	fmt.Printf("| cut %v, len: %d, cap: %d \n", cut, len(cut), cap(cut))

	fmt.Println("---------------------------")
	cut = cut[:cap(cut)]
	_ = cut[2]
	fmt.Printf("| cut %v, len: %d, cap: %d \n", cut, len(cut), cap(cut))

	fmt.Println("---------------------------")
	fmt.Printf(
		"| intSlice  %d, cut: %d |\n"+
			"| intSlice %v, cut: %v |\n",
		reflect.ValueOf(intSlice).Pointer(),
		reflect.ValueOf(cut).Pointer(),
		*(*reflect.SliceHeader)(unsafe.Pointer(&intSlice)),
		*(*reflect.SliceHeader)(unsafe.Pointer(&cut)),
	)
	fmt.Println("---------------------------")
	cut[0] = 1 << 32
	fmt.Printf("| intSlice %v, len: %d, cap: %d \n", intSlice, len(intSlice), cap(intSlice))
	fmt.Printf("| cut %v, len: %d, cap: %d \n", cut, len(cut), cap(cut))
	fmt.Println("---------------------------")

	cut = append(cut, (1<<32 + 1))
	cut[0] = 1 << 10
	fmt.Printf("| intSlice %v, len: %d, cap: %d \n", intSlice, len(intSlice), cap(intSlice))
	fmt.Printf("| intSlice %v, len: %d, cap: %d \n", cut, len(cut), cap(cut))
	fmt.Println("---------------------------")

}
