package main

import "fmt"

// reference types (functions)

// interface type

// this a curios way to define a function assigned a name for return
// but is not recomendable use it
// func main() {
// 	z := addTwoNumber(2, 4)
// 	fmt.Println(z)
// }

// func addTwoNumber(x, y int) (sum int) {
// 	sum = x + y
// 	return
// }
type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

//This a way to add functions  to own variables
func (a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println()
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
	fmt.Println()
}

func main() {
	// myTotal := sumMany(2, 3, 4, 5, 88, 7, -5)
	// fmt.Println(myTotal)

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "Woof"
	dog.NumberOfLegs = 4
	dog.Says()
	dog.HowManyLegs()

	cat := Animal{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}

	cat.HowManyLegs()

}

func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}

	return total
}
