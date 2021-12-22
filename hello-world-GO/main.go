package main

import (
	"fmt"
	"myapp/doctor"
)

// first run "go mod init myapp" on terminal to create own package
// Second build doctor structure for this case, can be any outher name

func main() {
	var whatToSay string

	whatToSay = doctor.Intro()
	fmt.Println(whatToSay)
}
