package main

import "fmt"

func pointone(n *int) {
	// operation of dereferencing
	*n = 1
}

func main() {
	a := 5
	fmt.Println(a)
	pointone(&a)
	fmt.Println(a)

}

/*
A pointer in Go is a variable that stores the memory address of another variable.
A pointer allows you to modify the original value directly,
as opposed to working with a copy of the value.
*/
