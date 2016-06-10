package main

import "fmt"

// "By considering the terms in the Fibonacci sequence whose values do not
// exceed four million, find the sum of the even-valued terms"
func main() {
	a := 1
	b := 2
	sum := 0
	for {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
		if b > 4000000 {
			break
		}
	}
	fmt.Println(sum)
}
