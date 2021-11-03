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

type User struct {
	UserName  string
	Age       int
	FavNumber float64
	OwnsADog  bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User
	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavNumber = readFloat("What is your favourite number?")
	user.OwnsADog = readBool("Do you own a dog? (y/n)")
	//fmt.Println("Your name is", userName, ". You are", age, "years old.")
	//fmt.Println("Your name is "+user.UserName+". You are", user.Age, "years old.")

	// fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old.", userName, age))
	fmt.Printf("Your name is %s. You are %d years old. Your favourite number is %.2f. Owns a dog: %t\n",
		user.UserName, user.Age, user.FavNumber, user.OwnsADog)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)

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
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}

func readBool(s string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println(s)
	for {
		userInput, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if userInput == 'n' || userInput == 'N' {
			return false
		} else if userInput == 'y' || userInput == 'Y' {
			return true
		} else {
			fmt.Println("Please (y/n)")
		}
	}
}
