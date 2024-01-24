package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
    fmt.Println(Hello("Chris"))
}

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
    return englishHelloPrefix + name
}
