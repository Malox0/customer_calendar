package greetings

import "fmt"

func Hello(name string) string {
	//Return a greeting to  given name

	message := fmt.Sprintf("Hello, dear %v.", name)
	return message
}
