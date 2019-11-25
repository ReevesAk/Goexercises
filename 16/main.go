package main

import (
	"fmt"
)

func foo(num ...int) int {
	total := 0
	for _, number := range num {
		total += number
	}
	return total
}

func bar(numb []int) int {
	total := 0
	for _, nx := range numb {
		total += nx
	}
	return total
}

func main() {
	fooNUmber := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(fooNUmber)
	barNumber := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fooN := foo(fooNUmber...)
	fmt.Println(fooN)
	barN := bar(barNumber)
	fmt.Println(barN)
}
