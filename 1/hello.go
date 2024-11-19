package main

import "fmt"

const (
	englishGreetingPrefix = "Hello "
	spanishGreetingPrefix = "Hola "
	frenchGreetingPrefix  = "Bonjour "

	french  = "French"
	spanish = "Spanish"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishGreetingPrefix
	case french:
		prefix = frenchGreetingPrefix
	default:
		prefix = englishGreetingPrefix
	}
	return
}

func main() {
	fmt.Println("Hello world")
}
