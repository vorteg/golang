package main

import (
	"myapp/packageone"
)

var myVar = "myvar from package level"

func main() {
	// variables
	var blockVar = "test"
	packageone.PrintMe(myVar)
	packageone.PrintMe(blockVar)
	packageone.PrintMe(packageone.PackageVar)

}
