package main

import "fmt"

func strrev(s string) string {
	str := ""
	for i := len(s) - 1; i >= 0; i-- {
		str += string(s[i])

	}
	return str
}

func main() {
	mystr := "hi"
	a := strrev(mystr)
	fmt.Println(a)

}
