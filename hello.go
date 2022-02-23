package hello

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

// Goodbye returns a farewell for the named person, with an added bonus
func Goodbye(name string) string {
	message := fmt.Sprintf("Goodbye, %v. Be always kind and true...", name)
	return message
}
