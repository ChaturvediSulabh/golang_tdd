package hello

import "fmt"

const greetPrefix = "Hello, "

// Hello - Hello, World
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return greetPrefix + name
}

func hello() {
	fmt.Println(Hello("world"))
}
