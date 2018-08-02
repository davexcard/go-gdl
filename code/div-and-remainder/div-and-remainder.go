package main

import "fmt"

func divAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	div, remainder := divAndRemainder(3, 2)
	fmt.Println(div, remainder)

	div, _ = divAndRemainder(10, 4)
	fmt.Println(div)

	_, remainder = divAndRemainder(100, -100)
	fmt.Println(remainder)

	divAndRemainder(-1, 20)
}
