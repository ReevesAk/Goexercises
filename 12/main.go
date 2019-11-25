package main

import "fmt"

func main() {
	mapping := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	//
	//fmt.Println(mapping)
	//
	//delete(mapping, `bond_james`)

	for index, ranging := range mapping {
		fmt.Println("Numbering index record: ", index)
		for indexOfRanging, value := range ranging {
			fmt.Printf("\t The record for activities : %v\n \t with value: %v\n", indexOfRanging, value)
		}
	}
}
