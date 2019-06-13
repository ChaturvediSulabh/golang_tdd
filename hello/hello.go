package hello

import "fmt"

// Hello - Greet
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	greetPrefix := prefix(language)
	return greetPrefix + name
}

func prefix(language string) string {
	var greetPrefix string
	switch language {
	case "Spanish":
		greetPrefix = "Hola, "
	case "French":
		greetPrefix = "Bonjour, "
	default:
		greetPrefix = "Hello, "
	}
	return greetPrefix
}

func hello() {
	fmt.Println(Hello("", ""))
}
