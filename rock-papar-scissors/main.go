package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1

	computerValue := rand.Intn(3)

	reader := bufio.NewReader(os.Stdin)

	//clearScreen()

	fmt.Print("Please enter rock, paper, or scissors ->")
	playerChoice, _ = reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	//clearScreen()
	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	fmt.Println()
	fmt.Println("Player chose", playerChoice, "and value is", playerValue)
	fmt.Println("Computer chose", computerValue)

}

// clearScreen clears the screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	// if strings.Contains(runtime.GOOS, "windows") {
	// 	// windows
	// 	cmd := exec.Command("clear")
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Run()
	// } else {
	// 	// linux or mac
	// 	cmd := exec.Command("clear")
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Run()
	// }
}
