package main

import "fmt"

type person struct {
	firstName           string
	lastName            string
	favIceCreamFlavours []string
}

func main() {
	Hanan := person{
		firstName: "Jehohanan",
		lastName:  "Ogieva",
		favIceCreamFlavours: []string{
			"Chocolate",
			"Vanilla",
			"Martini",
		},
	}

	Wisdom := person{
		firstName: "Wisdom",
		lastName:  "Olofua",
		favIceCreamFlavours: []string{
			"baby shit",
			"drunkard flavour",
			"bull crap taste",
		},
	}

	mapping := map[string]person{
		Hanan.lastName:  Hanan,
		Wisdom.lastName: Wisdom,
	}

	for index, values := range mapping {
		fmt.Printf("\t S/N: %v \t\n Names: %v ", index, values)
		for indexOfMapping, flavours := range values.favIceCreamFlavours {
			fmt.Printf("\nNo: %v\t These are the selected flavours: %v\n ", indexOfMapping, flavours)
		}
	}

}
