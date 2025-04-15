package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	sumup := sum(numbers...)
	fmt.Println(sumup)
}

func sum(args ...int) int {
	sum := 0

	for _, num := range args {
		sum += num
	}
	return sum
}
