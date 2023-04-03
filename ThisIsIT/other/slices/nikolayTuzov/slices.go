package main

import (
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4, 5}

	fmt.Printf("len:%v cap:%v %v \n", len(list), cap(list), list)
	double(list)
	fmt.Printf("len:%v cap:%v %v \n", len(list), cap(list), list)
	//fmt.Printf("len:%v cap:%v %v \n", len(double(list)), cap(double(list)), double(list))

	fmt.Printf("len:%v cap:%v %v \n", len(double1(list)), cap(double1(list)), double1(list))

	fmt.Printf("len:%v cap:%v %v \n", len(double2(list)), cap(double2(list)), double2(list))

}

// не оптимальный вариант
func double(nums []int) []int {
	var res []int
	for _, num := range nums {
		res = append(res, num*2) //cap(res) = 0, каждый цикл  проверяет cap и в случаи необходимости увеличивает на 2
	}
	return res
}

// обтимальный вариант№1
func double1(nums []int) []int {
	res1 := make([]int, 0, len(nums))
	for _, num := range nums {
		res1 = append(res1, num*2)
	}
	return res1
}

// обтимальный вариант№2
func double2(nums []int) []int {
	res2 := make([]int, len(nums))
	for i, num := range nums {
		res2[i] = num * 2
	}
	return res2
}
