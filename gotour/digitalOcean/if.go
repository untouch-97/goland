package main

import "fmt"

// https://www.digitalocean.com/community/tutorials/how-to-write-conditional-statements-in-go-ru
// Начнем с выражений if, которые оценивают истинность утверждения и запускают код, только если оно истинно.

func main() {
	//grade := 70
	//if grade >= 65 {
	//	fmt.Println("passing grade")

	//balance := 0
	//if balance < 0 {
	//	fmt.Println("Balance is below 0, add funds now or you will be charged a penalty.")
	//} else if balance == 0 {
	//	fmt.Println("Balance is equal to 0, add funds soon.")
	//} else {
	//	fmt.Println("Your balance is 0 or above.")

	//grade := 66
	//if grade >= 90 {
	//	fmt.Println("A grade")
	//} else if grade >= 80 {
	//	fmt.Println("B grade")
	//} else if grade >= 70 {
	//	fmt.Println("C grade")
	//} else if grade >= 65 {
	//	fmt.Println("D grade")
	//} else {
	//	fmt.Println("Failing grade")
	//}
	//}

	if statement1 { // outer if
		fmt.Println("hello world")
		if nested_statement1 { // first nested if
			fmt.Println("yes")
		} else if nested_statement2 { // first nested else if
			fmt.Println("maybe")
		} else { // first nested else
			fmt.Println("no")
		}
	} else if statement2 { // outer else if
		fmt.Println("hello galaxy")
		if nested_statement3 { // second nested if
			fmt.Println("yes")
		} else if nested_statement4 { // second nested else if
			fmt.Println("maybe")
		} else { // second nested else
			fmt.Println("no")
		}
	} else { // outer else
		statement("hello universe")
	}

	//grade := 92
	//if grade >= 65 {
	//		fmt.Print("Passing grade of: ")
	//
	//		if grade >= 90 {
	//			fmt.Println("A")
	//
	//		} else if grade >= 80 {
	//			fmt.Println("B")
	//
	//		} else if grade >= 70 {
	//			fmt.Println("C")
	//
	//		} else if grade >= 65 {
	//			fmt.Println("D")
	//		}
	//
	//	} else {
	//		fmt.Println("Failing grade")

	//	grade := 95
	//	if grade >= 65 {
	//		fmt.Print("Passing grade of: ")
	//	if grade >= 90 {
	//		if grade > 96 {
	//			fmt.Println("A+")
	//		} else if grade > 93 && grade <= 96 {
	//			fmt.Println("A")
	//		} else {
	//			fmt.Println("A-")
	//		}
	//	}
	//	}

}
