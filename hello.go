package hello

import (
	"fmt"

	"github.com/microsoft/dev-tunnels/go/tunnels"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func Hi(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func Create(name string) string {
	// Return a greeting that embeds the name in a message.
	tunnel := tunnels.Tunnel{
		Name: name,
	}
	return tunnel.Name
}
