package main

import "fmt"

func main() {
	fmt.Println("Проверка совпадения двух или более чисел.")

	fmt.Println("Введите первое число:")
	var numOne int
	fmt.Scan(&numOne)

	fmt.Println("Введите второе число:")
	var numTwo int
	fmt.Scan(&numTwo)

	fmt.Println("Введите третье число:")
	var numThree int
	fmt.Scan(&numThree)

	if numOne == numTwo || numOne == numThree || numTwo == numThree {
		fmt.Println("Два числа (или более) из ваших трёх чисел совпадают.")
	} else {
		fmt.Println("Все ваши числа являются разными")
	}
}
