package main

import "fmt"

func iterativefactorial(num int) int {
	if num < 0 {
		return 0
	}
	if num == 0 || num == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= num; i++ {
		result *= i
	}
	return result
}

func main() {
	num1 := iterativefactorial(4)
	fmt.Println(num1)
	num2 := iterativefactorial(5)
	fmt.Println(num2)

}
