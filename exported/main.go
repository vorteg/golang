package main

//This is an example to explain how to create a package and  export variables and functions into main package or other site.

import (
	"log"
	"myapp/staff"
)

//Here a way to initialized a variable outside package where was created.
var employees = []staff.Employee{
	{FirstName: "Jhon", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Sally", LastName: "Jhons", Salary: 50000, FullTime: true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 60000, FullTime: true},
	{FirstName: "Allan", LastName: "Anderson", Salary: 15000, FullTime: false},
	{FirstName: "Margaret", LastName: "Carter", Salary: 100000, FullTime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	// myStaff.ALL()

	// fmt.Println(myStaff.All())
	//staff.OverPaidLimit = 30000
	log.Println("Overpaid staff", myStaff.Overpaid())
	log.Println("Underpaid staff", myStaff.Underpaid())
	//myStaff.notVisible() is not exported
}
