package main

import "fmt"

func isnegative(nb int) {
	if nb < 0 {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}

}

func main() {
	a := 15
	b := -3
	isnegative(a)
	isnegative(b)

}
