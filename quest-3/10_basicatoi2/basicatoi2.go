package main

import "fmt"

func basicatoi2(s string) int {

	result := 0
	for i := 0; i < len(s); i++ {

		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		num := int(s[i] - '0')
		result = result*10 + num
	}
	return result

}

func main() {
	a := basicatoi2("hh123")
	fmt.Println(a)
	b := basicatoi2("10")
	fmt.Println(b)
	c := basicatoi2("0055")
	fmt.Println(c)

}
