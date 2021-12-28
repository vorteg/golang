package main

import "fmt"

// aggregate types (array, struct)
//array
// var myStrings [3]string
// myStrings[0] = "cat"
// myStrings[1] = "dog"
// myStrings[2] = "fish"

// fmt.Println("First element in array is", myStrings[0])

// reference types (pointers, slices, maps, functions, channels)

// interface type

//struct
type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {

	myCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XC90",
		Year:          2019,
	}

	fmt.Printf("My car is a %d %s %s \n", myCar.Year, myCar.Make, myCar.Model)

}
