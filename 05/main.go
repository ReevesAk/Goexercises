package main

import "fmt"

func main()  {
	yearBorn := 1994

	for yearsIHaveLived := yearBorn; yearsIHaveLived <= 2019; yearsIHaveLived++ {
		fmt.Println(yearsIHaveLived)
	}
}