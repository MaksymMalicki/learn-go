package main

import "fmt"

const englishPerfix = "Hello, "

func Hello(name string) string {
	if name == "" {
		return englishPerfix + "World"
	}
	return englishPerfix + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
