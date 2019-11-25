package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	p.first = "The"
	p.last =  "Terminator"

	(*p).last = "John"
	(*p). first = "Rambo"
}


func main() {
	agent := person{
		first: "Demolition",
		last:  "Man",
	}
	fmt.Println(agent)
	changeMe(&agent)
}
