package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	// fmt.Print("->")

	// userInput, _ := reader.ReadString('\n')

	// userInput = strings.Replace(userInput, "\n", "", -1)

	// if userInput == "quit" {

	// } else {
	// 	fmt.Println(userInput)
	// }

	err := keyboard.Open()

	// nil is a keyword used to check the variable/object is NULL
	if err != nil {
		// log is default package in Go
		log.Fatal(err)
	}

	// defer keyword is a built-in keyword and it runs only when the main() is done processing.
	// It's like a finally keyword from Java
	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("====")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	// rune is a single character
	for {
		// char, key, err := keyboard.GetSingleKey()
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		t := fmt.Sprintf("You chose %s", coffees[i])
		fmt.Println(t)

		// if key != 0 {
		// 	fmt.Println("You pressed", char, key)
		// } else {
		// 	fmt.Println("You pressed", char)
		// }

		// if key == keyboard.KeyEsc {
		// 	break
		// }
	}

	fmt.Println("Program exiting.")
}
