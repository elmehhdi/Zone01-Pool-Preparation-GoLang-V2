package main

import "fmt"

func sqrt(number int) int {
	if number <= 0 {
		return 0
	}

	i := 1
	for {
		if i*i == number {
			return i
		}
		if i > number/i {
			break
		}
		i++
	}

	return 0
}

func main() {
	a := sqrt(25)
	fmt.Println(a)
	b := sqrt(26)
	fmt.Println(b)
}
