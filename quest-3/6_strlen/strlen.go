package main

import "fmt"

func strlen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count

}

func main() {
	str := "hi"
	fmt.Println(strlen(str))

}
