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
	// *** added two variables to keep track of score
	playerScore := 0
	computerScore := 0
	drawScore := 0

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	// *** print out some instructions
	fmt.Println("Rock, Paper, & Scissors")
	fmt.Println("-----------------------")
	fmt.Println("Game is played for three rounds, and best of three wins the game. Good luck!")
	fmt.Println()

	// *** added the for loop
	for i := 1; i <= 3; i++ {
		// *** print out round information
		fmt.Println()
		fmt.Println("Round", i)
		fmt.Println("-------")

		fmt.Print("Please enter rock, paper, or scissors ->")
		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)
		playerChoice = strings.Replace(playerChoice, "\n", "", -1)

		// moved computer choice to for loop, so it's reset each time through
		computerValue := rand.Intn(3)
		clearScreen()

		if playerChoice == "rock" {
			playerValue = ROCK
		} else if playerChoice == "paper" {
			playerValue = PAPER
		} else if playerChoice == "scissors" {
			playerValue = SCISSORS
		} else {
			// *** reset playerValue to -1 if  invalid choice
			playerValue = -1
		}

		// *** print out player choice
		fmt.Println()
		fmt.Println("Player chose", strings.ToUpper(playerChoice))

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
		fmt.Println()

		if playerValue == computerValue {
			fmt.Println("It's a draw")
			drawScore++
		} else {
			switch playerValue {
			case ROCK:
				if computerValue == PAPER {
					computerScore = computerWins(computerScore)

				} else {
					playerScore = playerWins(playerScore)
				}
				break
			case PAPER:
				if computerValue == SCISSORS {
					computerScore = computerWins(computerScore)

				} else {
					playerScore = playerWins(playerScore)
				}

				break
			case SCISSORS:
				if computerValue == ROCK {
					computerScore = computerWins(computerScore)

				} else {
					playerScore = playerWins(playerScore)
				}

				break
			default:
				fmt.Println("Invalid choise")
				i--
			}
		}

	}

	clearScreen()
	fmt.Println("Final score")
	fmt.Println("-----------")
	fmt.Printf("Player: %d/3, Computer %d/3, It's a draw %d/3", playerScore, computerScore, drawScore)
	fmt.Println()

	if playerScore > computerScore {
		fmt.Println("Player wins game!")
	} else {
		fmt.Println("Computer wins game!")
	}

}

func computerWins(score int) int {
	fmt.Println("Computer wins!")
	return score + 1
}

func playerWins(score int) int {
	fmt.Println("player wins!")
	return score + 1
}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
