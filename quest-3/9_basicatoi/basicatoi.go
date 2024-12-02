package main

import "fmt"

func basicatoi(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		num := int(s[i] - '0')   // la convertion : i = '3' = 51 et '0' = 48 ; 51 - 48 = 3 la valeur numerique
		result = result*10 + num // l'accumulation
	}
	return result
}

func main() {
	fmt.Println(basicatoi("123"))
	fmt.Println(basicatoi("000012345"))
	fmt.Println(basicatoi("000"))
	fmt.Println(basicatoi("456"))
	a := basicatoi("10")
	b := basicatoi("5")
	fmt.Println(a + b)

}
