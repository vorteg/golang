package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
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

	clearScreen()

	fmt.Print("Please enter rock, paper, or scissors ->")
	playerChoice, _ = reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	if playerChoice == "rock" {
		playerValue = ROCK
		fmt.Println("Player chose", playerValue)
	} else if playerChoice == "paper" {
		playerValue = PAPER
		fmt.Println("Player chose", playerValue)
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
		fmt.Println("Player chose", playerValue)
	}

	switch computerValue {
	case ROCK:
		fmt.Println("Computer chose Rock")
		break
	case PAPER:
		fmt.Println("Computer chose Paper")
		break
	case SCISSORS:
		fmt.Println("Computer chose Scissors")
		break
	default:
	}

	if playerValue == computerValue {
		fmt.Println("It's a draw")
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				fmt.Println("Computer wins!")
			} else {
				fmt.Println("Player wins!")
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				fmt.Println("Computer wins!")
			} else {
				fmt.Println("Player wins!")
			}

			break
		case SCISSORS:
			if computerValue == ROCK {
				fmt.Println("Computer wins!")
			} else {
				fmt.Println("Player wins!")
			}

			break
		default:
			fmt.Println("Invalid choise")
		}
	}

}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "c/", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
