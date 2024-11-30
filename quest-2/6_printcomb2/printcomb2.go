package main

import "fmt"

// aa < bb
func printComb2() {
	for a := 0; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {
			fmt.Printf("%02d %02d", a, b)

			if a != 98 || b != 99 {
				fmt.Print(", ")
			}
		}
	}
	fmt.Println()
}

func main() {
	printComb2()
}
