package main

import "fmt"

func printnbr(n int) int {
	if n < 0 {
		fmt.Print("-")
		// ila kan negative n'convertew l positive number -(-n) = +n
		n = -n
	}

	if n == 0 {
		fmt.Println("0")
		return 0
	}
	// ila kan positive n'printiweh kima rah
	fmt.Println(n)
	return n
}

func main() {
	a := 5
	b := 0
	printnbr(-123)
	printnbr(456)
	printnbr(0)
	res := printnbr(a) + printnbr(b)
	fmt.Println(res)

}
