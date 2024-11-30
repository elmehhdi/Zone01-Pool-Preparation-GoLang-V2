package main

import "fmt"

func ultimatepointone(p ***int) {
	***p = 1

}

func main() {

	a := 15
	b := &a
	c := &b
	// 9bel
	fmt.Println(a)
	ultimatepointone(&c)
	// apres
	fmt.Println(a)

}
