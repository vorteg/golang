package main

import "fmt"

//Here an example of  how to use go rutines

func main() {
	go doSomthing("Hello, world")
	fmt.Println("This is another message")
	for {
		// do nothing
	}
}

func doSomthing(s string) {
	until := 0
	for {
		fmt.Println("s is ", s)
		until = until + 1
		if until >= 5 {
			break
		}
	}
}
