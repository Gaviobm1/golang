package main

import "fmt"

type transformFn func(int) int

func main() {
	number := []int{1, 2, 3, 4, 5, 6}
	doubled := transformNumbers(&number, double)
	tripled := transformNumbers(&number, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	transformed := []int{}
	for _, val := range *numbers {
		transformed = append(transformed, transform(val))
	}
	return transformed
}

func double(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3
}
