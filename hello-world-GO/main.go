package main // this a convention name

import "fmt" // this one main package of Golang

//This a name function name convention
func main() {
	// golang dosen't allow any unuse variable
	// var whatToSay string = "Hello world, again from variable"

	//This other choice to assaign a variable is to define type variable automaticly
	whatToSay := "Hello world with define type varible"
	sayHelloWorld(whatToSay)
}

//How to create another function and called from main
func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
