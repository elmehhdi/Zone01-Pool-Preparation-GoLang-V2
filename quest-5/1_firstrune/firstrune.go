package main

import "fmt"

func firstrune(s string) rune {
	for _, r := range s {
		return r
	}
	// ila string kan khawi
	return 0
}

func main() {
	fmt.Println(string(firstrune("ok")))
	fmt.Println(string(firstrune("hello")))
	fmt.Println(string(firstrune("hi")))
	fmt.Println(string(firstrune("1234")))

}
