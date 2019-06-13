package hello

import "fmt"

// Hello - Hello, World
func Hello(name string) string {
	return "Hello, " + name
}

func hello() {
	fmt.Println(Hello("world"))
}
