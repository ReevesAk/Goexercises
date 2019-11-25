package main

import "fmt"

func main() {
	arrayContainer := [5]int{25, 27, 31, 33, 50}

	for index, arrayValues := range arrayContainer {
		fmt.Println(index, arrayValues)
	}

	fmt.Printf("The underlying type of the array being: %T\n", arrayContainer)
}
