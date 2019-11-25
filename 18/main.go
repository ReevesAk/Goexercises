package main

import "fmt"

type person struct {
	first       string
	last 		string
	age     	int
}

func (p  *person) speak()  {
	fmt.Println("My name is ", p.first, p.last, "and I am ", p.age, "years old")
}

func main()  {
	human := person{
		first:		"Reeves",
		last:		"Akwa",
		age:		25,
	}
	 human.speak()
}