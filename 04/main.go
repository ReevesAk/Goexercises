package main

import (
	"fmt"
)

//nested loop printing every rune code point of the
// uppercase alphabet three times. Your output should look like this:
// 65
//  U+0041 'A'
//  U+0041 'A'
//  U+0041 'A'



func main()  {
	for index := 65; index <= 90; index++ {
		fmt.Println(index)
		for indexPoundType := 0; indexPoundType <= 3; indexPoundType++  {
			fmt.Printf("\t%U\n", index)
		}
	}
}
