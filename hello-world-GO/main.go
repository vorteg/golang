package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
)

// Simple version of a loop

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	for {
		userInput, _ := reader.ReadString('\n')

		fmt.Println("User commented: " + userInput)
	}

}
