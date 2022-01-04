package main

import "fmt"

//This an example how to implement composition with OOP migration to GO

type Vechicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vechicle
}

func (v Vechicle) showDetails() {
	fmt.Println("Number of passengers:", v.NumberOfPassengers)
	fmt.Println("Number of wheels:", v.NumberOfWheels)
}

func (c Car) show() {
	fmt.Println("Make:", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("Year:", c.Year)
	fmt.Println("Is Electric:", c.IsElectric)
	fmt.Println("Is Hybrid:", c.IsHybrid)
	c.Vehicle.showDetails()
}

func main() {
	suv := Vechicle{
		NumberOfWheels:     4,
		NumberOfPassengers: 6,
	}

	volvoXC90 := Car{
		Make:       "Volvo",
		Model:      "XC90 T8",
		Year:       2021,
		IsElectric: false,
		IsHybrid:   true,
		Vehicle:    suv,
	}

	volvoXC90.show()
	fmt.Println()

	teslaModelX := Car{
		Make:       "Tesla",
		Model:      "Model X",
		Year:       2021,
		IsElectric: true,
		IsHybrid:   false,
		Vehicle:    suv,
	}

	teslaModelX.Vehicle.NumberOfPassengers = 7

	teslaModelX.show()
}
