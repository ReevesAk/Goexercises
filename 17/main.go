package main

import "fmt"

func main() {
	defer def()
	fmt.Println("I love you")
}

func def() {
	defer func() {
		fmt.Println("Let's do this")
	}()

	fmt.Println("Def was defered")
}
