package main

import "fmt"

//interface type

// Dog is the type for dogs
type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

// Cat is the type for cats
type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func main() {
	// ask a riddle
	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}
	Riddle(dog)
	var cat Cat
	cat.Name = "cat"
	cat.NumberOfLegs = 4
	cat.Sound = "Meow"

	Riddle(cat)
}

func Riddle(d Dog) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is it?`, d.Sound, d.NumberOfLegs)
	fmt.Println(riddle)
}
