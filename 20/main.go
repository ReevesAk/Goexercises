package main

import "fmt"

func wisdom() func() string {
	return func() string {
		return "Nama"
	}
}

func main() {
	programmer := wisdom()
	fmt.Println(programmer())
}
