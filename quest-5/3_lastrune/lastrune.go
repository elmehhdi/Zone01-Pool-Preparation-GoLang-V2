package main

import "fmt"

func lastrune(s string) rune {
	if s == "" {
		return 0
	}

	var lastrune rune
	for _, r := range s {
		lastrune = r
	}

	return lastrune
}

func main() {
	fmt.Printf("%c\n", lastrune("hello"))
	fmt.Printf("%c\n", lastrune("ok"))
	fmt.Printf("%c\n", lastrune("fine"))
	fmt.Printf("%c\n", lastrune("123"))
	fmt.Printf("%c\n", lastrune("4455"))
	fmt.Printf("%c\n", lastrune("60"))

}
