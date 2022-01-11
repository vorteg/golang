package main

import "fmt"

func main() {
	// index start at 0
	//                       1         2         3
	//             01234567890123456789012345678901234

	courseName := "Learn Go for Beginners Crash Course"
	fmt.Println(string(courseName[0]))

	for i := 0; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}
}
