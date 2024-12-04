package main

import "fmt"

func isprime(num int) bool {

	if num <= 1 {
		return false
	}

	// n'tchekew men 2 7ta l square dyal num
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isprime(5))
	fmt.Println(isprime(4))
	fmt.Println(isprime(11))
	fmt.Println(isprime(1))
	fmt.Println(isprime(0))
	fmt.Println(isprime(100))
	fmt.Println(isprime(12))
	fmt.Println(isprime(4))
	fmt.Println(isprime(3))

}
