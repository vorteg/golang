package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	name := "Hello world"
	fmt.Println("String:", name)
	fmt.Println()

	fmt.Println("Bytes")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x \n", name[i])
	}

	tab_char(name)
	name = "Γειά σου Κόσμε"
	tab_char(name)

	// Three ways to concatenate strings
	fmt.Println("Three ways to concatenate strings")
	h := "Hello, "
	w := "world."

	//using + not very efficient
	myString := h + w
	fmt.Println(myString)
	fmt.Println()

	//using fmt - more efficient
	myString = fmt.Sprintf("%s%s", h, w)
	fmt.Println(myString)
	fmt.Println()

	//using stringBuilder - very efficient
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String())

	fmt.Println()
	name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println("Getting a substring")
	fmt.Println(name[0:13])
}

func tab_char(name string) {
	fmt.Println("Index\tRune\tString")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}
	fmt.Println()
}
