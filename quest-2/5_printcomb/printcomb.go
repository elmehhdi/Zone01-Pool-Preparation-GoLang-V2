package main

import "fmt"

// a < b < c
func printcomb() {
	for a := 0; a <= 7; a++ {
		for b := a + 1; b <= 8; b++ {
			for c := b + 1; c <= 9; c++ {
				fmt.Printf("%d%d%d", a, b, c)

				if a != 7 || b != 8 || c != 9 {
					fmt.Print(", ")
				}
			}
		}
	}
	fmt.Println()
}

func main() {
	printcomb()
}
