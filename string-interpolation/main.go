package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

// User-defined type
type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwnsADog        bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User
	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavouriteNumber = readFloat("What is your favourite number?")
	user.OwnsADog = readBool("Do you own a dog? (y/n)")
	// userName := readString("What is your name?")
	// age := readInt("How old are you?")
	// fmt.Println("Your name is: "+userName+". You are", age, "years old.")  # First way of printing
	// fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old", userName, age)) # Second way of printing

	// fmt.Printf("Your name is %s. You are %d years old", userName, age) // Third way of printing
	if user.OwnsADog {
		fmt.Printf("Your name is %s. You are %d years old. Your favourite number is %.4f. You own a dog.",
			user.UserName,
			user.Age,
			user.FavouriteNumber)
	} else {
		fmt.Printf("Your name is %s. You are %d years old. Your favourite number is %.4f. You don't own a dog.",
			user.UserName,
			user.Age,
			user.FavouriteNumber)
	}

}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = replaceCarriageReturns(userInput)
		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func replaceCarriageReturns(userInput string) string {
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = replaceCarriageReturns(userInput)
		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = replaceCarriageReturns(userInput)
		num, err := strconv.ParseFloat(userInput, 64)

		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}

func readBool(s string) bool {
	// fmt.Println(s)

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		prompt()
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("Please type y or n")
		} else if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'Y' {
			return true
		}
	}
	// userInput, _ := reader.ReadString('\n')
	// userInput = replaceCarriageReturns(userInput)
	// var isTrue bool
	// if strings.ToLower(userInput) == "y" {
	// 	isTrue = true
	// } else {
	// 	isTrue = false
	// }
	// return isTrue
}
