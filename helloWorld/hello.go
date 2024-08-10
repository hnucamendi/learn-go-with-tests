package main

import "fmt"

const englishHelloWorldPrefix string = "Hello "

func Hello(s string) string {
	if s == "" {
		return englishHelloWorldPrefix + "World!"
	}
	return englishHelloWorldPrefix + s + "!"
}

func main() {
	fmt.Println(Hello("David"))
}
