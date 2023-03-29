package main

import "fmt"

func main() {

	for i := 0; i <= 20; i++ {
		if i%2 == 1 {
			continue
		}
		//		fmt.Println(i)
	}

first:
	for i := 1; i <= 20; i++ {
		for k := 1; k <= 10; k++ {
			//			fmt.Println("I: ", i, "K: ", k)
			if k >= 5 {
				continue first
			}
		}
	}

	for i := 0; i <= 20; i++ {
		if i >= 10 {
			break
		}
		//		fmt.Println(i)
	}
second:
	for i := 1; i <= 20; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println("I: ", i, "J: ", j)
			if i >= 10 {
				break second
			}
		}

	}
}
