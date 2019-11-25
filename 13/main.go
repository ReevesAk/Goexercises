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

	fmt.Println(Hanan)
	for flavours, list := range Hanan.favIceCreamFlavours {
		fmt.Printf("\t These are Hanan's flavours: %v\t %v\n ", flavours, list)
	}

	fmt.Println(Wisdom)
	for flavours, list := range Wisdom.favIceCreamFlavours {
		fmt.Printf("\t These are Wisdom's flavours: %v\t %v\n ", flavours, list)
	}
}
