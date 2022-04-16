package hello

import "fmt"

const french = "French"
const spanish = "Spanish"
const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s, %s", greetingPrefix(lang), name)
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Peter", ""))
}
