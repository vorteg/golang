package main

import (
	"fmt"
	"myapp/packageone"
)

var one = "one"

func main() {
	//Allways avoid this, since override the variable
	var one = "This is a block level variable"

	fmt.Println(one)
	myFunc()

	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)

	//The function is available since the name was defined first letter of name in capital
	packageone.Exported()

}

func myFunc() {
	fmt.Println(one)
}
