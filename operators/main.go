package main

import (
	"fmt"
	"math"
)

func main() {
	// multiplication
	// area = nr2
	var radius = 12.0
	area := math.Pi * radius * radius
	fmt.Println("Area is ", area)

	// integer division
	half := 1 / 2
	fmt.Println("half with integer division", half)

	// float division
	halfFloat := 1.0 / 2.0
	fmt.Println("half float", halfFloat)

	//squering (raising to the power)
	badThreeSquared := 3 ^ 2 // this in reality a bit operator
	fmt.Println("badThreeSquared ", badThreeSquared)
	goodThreeSquered := math.Pow(3.0, 2.0)
	fmt.Println("goodThreeSquared", goodThreeSquered)

	// modulus
	remainder := 50 % 3
	fmt.Println("remainder", remainder)

	// unary operators
	x := 3
	x++
	fmt.Println("x is now", x)

	x--
	x--
	fmt.Println("x is now ", x)
}
