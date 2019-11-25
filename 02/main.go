package main

import "fmt"

//Create your own type. Have the underlying type be an int.
//create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
//in func main
//print out the value of the variable “x”
//print out the type of the variable “x”
//assign 42 to the VARIABLE “x” using the “=” OPERATOR
//print out the value of the variable “x”

type person int

var x  person
var y  int

func main()  {

	fmt.Println(x)
	fmt.Printf("%T\t", x)
	x = 42
	fmt.Println(x)


	//now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
	//then use the “=” operator to ASSIGN that value to “y”
	//print out the value stored in “y”
	//print out the type of “y”

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)

}