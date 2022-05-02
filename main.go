package main

import (
	"fmt"
	"go-booking-app/basic"
)

var firstName string

func main() {
	basic.Greet()
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	basic.GreetUser(firstName)
}
