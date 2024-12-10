package main

import "fmt"

func index(s string, char_to_find string) int {
	if char_to_find == "" {
		return 0
	}
	for i := 0; i <= len(s)-len(char_to_find); i++ {
		if s[i:i+len(char_to_find)] == char_to_find {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(index("hello", "l"))
	fmt.Println(index("salam", "m"))
	fmt.Println(index("ok", "k"))
	fmt.Println(index("im working ", "o"))
	fmt.Println(index("123456789", "8"))
	fmt.Println(index("4455", "3"))
	fmt.Println(index("test", "e"))

}
