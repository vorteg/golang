package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
)

// first run "go mod init myapp" on terminal to create own package
// Second build doctor structure for this case, can be any outher name

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	userInput, _ := reader.ReadString('\n')

	fmt.Println(userInput)
}
