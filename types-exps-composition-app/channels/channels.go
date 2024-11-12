package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// Example of a goroutine

var keyPressChan chan rune

func main() {

	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit")
	_ = keyboard.Open()

	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()

		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char
	}

	// go doSomething("Hello, world")

	// fmt.Println("This is another message.")
	// for {
	// 	// do nothing
	// }
}

func listenForKeyPress() {

}

// func doSomething(s string) {
// until := 0
// for {
// 	fmt.Println("s is", s)
// 	until = until + 1
// 	if until >= 5 {
// 		break
// 	}
// }
// }
