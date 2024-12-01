package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp

}

func main() {
	a := 20
	b := 1
	swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
