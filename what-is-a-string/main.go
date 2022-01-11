package main

import (
	"fmt"
	"strings"
)

func main() {

	newString := "Go is a great programing language. Go for it!"

	if strings.Contains(newString, "Go") {

		newString = strings.ReplaceAll(newString, "Go", "Golang")
		//newString = strings.Replace(newString, "Go", "Golang", 1) // by -1 repleace all words
	}

	fmt.Println(newString)

	// string comparison
	if "Alpha" > "Absolute" {
		fmt.Println("A is greater than B")

	} else {
		fmt.Println("A is not greater than B")
	}

	badEMail := " me@here.com "
	//Deleting spaces
	badEMail = strings.TrimSpace(badEMail)
	fmt.Printf("=%s=", badEMail)
	fmt.Println()
}
