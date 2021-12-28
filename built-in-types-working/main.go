package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type

//Pointers allows modify values directly to having access to allocation memory address
// with no parse variable or change scope
func main() {
	var myInt int
	myInt = 10

	myFirstPointer := &myInt // with "&" we can have acces to variable address

	fmt.Println(myInt)
	fmt.Println("myFirstPointer is ", myFirstPointer)

	*myFirstPointer = 15 // with "*" we can change the value directly of a variable

	fmt.Println("myInt is now", myInt)

	changeValueOfPointer(&myInt)
	fmt.Println("After function call, myInt is noe", myInt)
}

func changeValueOfPointer(num *int) {
	*num = 25
}
