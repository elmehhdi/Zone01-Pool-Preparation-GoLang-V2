package main

import "fmt"

func atoi(s string) int {
	result := 0

	// 1 = positive et -1 = negative .. f lowel nkheliwha positive
	sign := 1

	// nkheliwha false w ghadi tweli true ila l9ina + ou -
	sign_found := false

	for i := 0; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' {
			if sign_found || i > 0 {
				return 0
			}
			sign_found = true

			// ila kan - ndiro num negative
			if s[i] == '-' {
				sign = -1
			}
			continue
		}

		if s[i] >= '0' && s[i] <= '9' {
			num := int(s[i] - '0')   // n'convertew
			result = result*10 + num // l'accumulation
		} else {
			// ila machi num rah invalide w n'returniw 0 3la 7sab task
			return 0
		}
	}
	return result * sign
}

func main() {
	fmt.Println(atoi("12"))
	fmt.Println(atoi("-123"))
	fmt.Println(atoi("+12"))
	fmt.Println(atoi("01234"))
	fmt.Println(atoi("1"))
	fmt.Println(atoi("12"))
	fmt.Println(atoi("97"))
	fmt.Println(atoi("43"))
	fmt.Println(atoi("102"))
	fmt.Println(atoi("61"))
	fmt.Println(atoi("108"))
	fmt.Println(atoi("1234m5"))
	fmt.Println(atoi("++123"))
	fmt.Println(atoi("--123"))
	fmt.Println(atoi("1ok234"))
	fmt.Println(atoi("500"))
	fmt.Println(atoi("-4"))

}
