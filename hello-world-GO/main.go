package main // this a convention name

import "fmt" // this one main package of Golang

//This a name function name convention
func main() {
	sayHelloWorld("Hello, world again from sayHelloWorld")
}

//How to create another function and called from main
func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
