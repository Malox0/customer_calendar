package main

import (
	"fmt"

	"github.com/Malox0/customer_calendar/greetings"
)

func main() {

	var message string
	message = greetings.Hello("Julian")
	fmt.Println(message)
}
