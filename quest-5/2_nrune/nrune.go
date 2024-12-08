package main

import "fmt"

func nrune(s string, n int) rune {
	// ila kan n 9el wla fog la longueur nta3 string n'returniw 0
	if n <= 0 {
		return 0
	}
	if n > len(s) {
		return 0
	}
	// si nn n'returniw n-1 l'indexation katbda f zero
	return rune(s[n-1])
}

func main() {
	a := nrune("hi", 2)
	fmt.Printf("%c\n", a)
	b := nrune("mehdi", 3)
	fmt.Printf("%c\n", b)
	c := nrune("hello", 5)
	fmt.Printf("%c\n", c)
}
