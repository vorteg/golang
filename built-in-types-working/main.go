package main

import "log"

// basic types (numbers, strings, booleans)
var myInt int
var myInt16 int16
var myInt32 int32
var myInt64 int64

var myUnit uint // unsigned value

var myFloat float32
var myFLoat64 float64

// aggregate types (array, struct)

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main() {
	myInt = 10
	myUnit = 20

	myFloat = 10.1
	myFLoat64 = 100.1

	log.Println(myInt, myUnit, myFloat, myFLoat64)

	myString := "Trevor"

	log.Println(myString)

	myString = "John" // We can chang the string but not the type since is inmutable
	log.Println(myString)

	var myBool = true
	log.Println(myBool)
}
