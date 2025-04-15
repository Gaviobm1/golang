package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)

}

func factorial(val int) int {
	if val == 1 {
		return val
	}
	return val * factorial(val-1)
}
