package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 12
	b := 6

	c, err := divideTwoNumbers(a, b)

	if err != nil {
		fmt.Print(err)
	} else {
		if c == 2 {
			fmt.Println("we found a two")
		}
	}

}

func divideTwoNumbers(x, y int) (int, error) {

	if y == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return x / y, nil
}
