package main

import "fmt"

func main() {
	sliceOfStrings := make([]string, 10)

	sliceOfStrings[0] = "Wisdom"
	sliceOfStrings[1] = "Reeves"
	sliceOfStrings[2] = "Elsie"
	sliceOfStrings[3] = "Victor"
	sliceOfStrings[4] = "Bassey"
	sliceOfStrings[5] = "Yvonne"
	sliceOfStrings[6] = "Diana"
	sliceOfStrings[7] = "Christopher"
	sliceOfStrings[8] = "Raphael"
	sliceOfStrings[9] = "Prisca"

	for i, printSlice := range sliceOfStrings {
		fmt.Println(i, printSlice)
	}
	fmt.Printf("The underlying type of the slice above is: %T\n ", sliceOfStrings)
}
