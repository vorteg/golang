package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//const never change
const prompt = "and don't type your number in, just press ENTER when ready."

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// rand generates a number between 0 and watever is passed as a parameter
	// we add 2 to it because we want the number used to be at least 2, and no
	// greater than 10 (multiplying by 1 makes the game a bit silly)

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction

	playTheGame(firstNumber, secondNumber, subtraction, answer)
}

//Implenting a function "One function, one prpose"
func playTheGame(firstNumber, secondNumber, subtraction, answer int) {
	// create our reader variable, wich reads input from standard in (the keyboard)
	reader := bufio.NewReader((os.Stdin))

	// display a welcome/instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number betwen 1 to 10", prompt)
	reader.ReadString('\n')

	// take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	fmt.Println("The answer is", answer)
}
