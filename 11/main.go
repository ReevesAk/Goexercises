package main

import (
	"fmt"
)

func main() {
	jB := []string{"James", "Bond", "Shaken, not stirred."}
	mP := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(jB)
	fmt.Println(mP)

	double07 := [][]string{jB, mP}
	fmt.Println(double07)

	for IndexOf007, ran := range double07 {
		fmt.Println("Records: ", IndexOf007)
		for IndexOfIndexOf007, value := range ran {
			fmt.Printf("\t position : %d \t value: %s\n", IndexOfIndexOf007, value)
		}
	}

}
