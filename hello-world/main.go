package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	// var whatToSay string
	// whatToSay = "Hello-World, again !"
	// whatToSay := "Hello-World, again!"
	// sayHelloWorld(whatToSay)

	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()
	// sayHelloWorld(whatToSay)
	fmt.Println(whatToSay)

	// for {} will run always. it's the same as while(true) in Java
	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n') // interpreter will wait to complete the program until the user presses Enter command.

		// fmt.Println(userInput)
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}

	}

}

// func sayHelloWorld(whatToSay string) {
// 	fmt.Println(whatToSay)
// }
