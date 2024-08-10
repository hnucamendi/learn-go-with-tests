package main

import "fmt"

func Hello(s string) string {
	return fmt.Sprintf("Hello %s!", s)
}

func main() {
	fmt.Println(Hello("David"))
}
