package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

// Simple version of a loop

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	for {
		fmt.Print("->")
		userInput, _ := reader.ReadString('\n')
		println(userInput)
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		println(userInput)
		userInput = strings.Replace(userInput, "\n", "", -1)
		println(userInput)

		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}

	}

}
