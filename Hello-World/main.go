package main

import "fmt"

func main() {
	//var whatToSay string
	//whatToSay = "Hi, World!"

	whatToSay := "Hi, World!"

	fmt.Println("Hello, World!")
	sayHelloWorld(whatToSay)
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
