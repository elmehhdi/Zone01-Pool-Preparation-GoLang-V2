package main

import "fmt"

func printstr(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
	}

}

func main() {

	printstr("hello")
}
