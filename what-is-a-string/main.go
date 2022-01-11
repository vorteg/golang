package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "This is a clear example of why we search in one case only"

	if strings.Contains(myString, "this") {

		if strings.Contains(myString, "this") {
			fmt.Println("Found it!")

		} else {
			fmt.Println("DId not find it!")
		}
	}

	// other case functions
	fmt.Println(strings.ToLower(myString))
	fmt.Println(strings.ToUpper(myString))
	fmt.Println(strings.Title(myString))
}
