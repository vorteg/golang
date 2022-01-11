package main

import "fmt"

func main() {
	second := 31
	minute := 1
	// the parentesis makes the code more readable but is redundant and not necessary
	if (minute < 59) && ((second + 1) > 59) {
		minute++
	}

	fmt.Println(minute)

}
