package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type

//maps
func main() {
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	//delete a key
	//delete(intMap, "four")

	//check if exist a key is in map

	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is in not map")
	}

	// modify a value

	intMap["two"] = 4

}
