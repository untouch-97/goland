package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "Uzbekistan"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
}

//Константы объявляются как переменные,
//но с помощью ключевого слова const.
//Константы могут иметь символьные, строковые,
//булевые или числовые значения.
//Константы не могут быть объявлены с помощью синтаксиса :=
