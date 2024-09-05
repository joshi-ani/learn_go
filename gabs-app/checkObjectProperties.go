package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

func main() {
	// Sample JSON object
	jsonObj := `{
		"name": "John Doe",
		"age": 30,
		"email": "johndoe@example.com",
		"address": {
			"city": "New York",
			"zipcode": "10001"
		}
	}`

	// Parse the JSON object into a Gabs container
	parsedJSON, err := gabs.ParseJSON([]byte(jsonObj))
	if err != nil {
		panic(err)
	}

	// Iterate over all keys in the object and print key-value pairs
	for key, value := range parsedJSON.ChildrenMap() {
		fmt.Printf("Key: %s, Value: %v\n", key, value.Data())
	}

	// If the value is also a JSON object, you can further iterate over it
	if address, ok := parsedJSON.Path("address").Data().(map[string]interface{}); ok {
		for key, value := range address {
			fmt.Printf("Key: address.%s, Value: %v\n", key, value)
		}
	}

	for key, child := range parsedJSON.ChildrenMap() {
		value := child.Data()
		if _, isMap := value.(map[string]interface{}); isMap {
			fmt.Printf("Key: %s is a map (JSON object)\n", key)
		} else {
			fmt.Printf("Key: %s is not a map, value: %v\n", key, value)
		}
	}
}
