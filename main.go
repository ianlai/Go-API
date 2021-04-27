package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("3 + 5 = ", sum(3, 5))
	fmt.Println("3 * 5 = ", multiply(3, 5))
}

func sum(n1, n2 int) int {
	return n1 + n2
}

func multiply(n1, n2 int) int {
	return n1 * n2
}
