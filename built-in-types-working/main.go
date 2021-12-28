package main

import (
	"fmt"
	"sort"
)

// reference types (pointers, slices, maps, functions, channels)

// interface type

//slices
func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)

	//slice and maps iteration

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println("Element 0 is", animals[0])
	fmt.Println("First two elements are ", animals[0:2])
	fmt.Println("The slice is ", len(animals), "elements long")
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)

}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}
