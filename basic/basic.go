package basic

import "fmt"

func Greet() {
	fmt.Println("Hello World")
}

func GreetUser(username string) {
	fmt.Println("Welcome to booking app", username, "!")
}
