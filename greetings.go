package greetings

import (
	"fmt"
	"rsc.io/quote"
)

func Hello(name string) string {
	// return a string that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}