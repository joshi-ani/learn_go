package main

import "fmt"

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

	// Delete the key from map
	delete(intMap, "four")
	// for key, value := range intMap {
	// 	fmt.Println(key, value)
	// }

	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in the map")
	} else {
		fmt.Println(el, "is not in the map")
	}

	intMap["two"] = 4
	fmt.Println(intMap)
}
