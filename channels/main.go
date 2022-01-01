package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// Here a basic channels feature example
var keyPressChan chan rune

func main() {
	keyPressChan = make(chan rune)

	go listenForKeyPres()
	fmt.Println("Press any key, or q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char
	}
}

func listenForKeyPres() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}
}
