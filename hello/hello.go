package hello

import "fmt"

// Hello - Greet
func Hello(name, language string) string {
	var greetPrefix = "Hello, "
	if name == "" {
		name = "World"
	}
	switch language {
	case "Spanish":
		greetPrefix = "Hola, "
	case "French":
		greetPrefix = "Bonjour, "
	}
	return greetPrefix + name
}

func hello() {
	fmt.Println(Hello("", ""))
}
