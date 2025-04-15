package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	double := createTransformer(2)
	triple := createTransformer(3)
	quadruple := createTransformer(4)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	quadrupled := transformNumbers(&numbers, quadruple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
	fmt.Println(quadrupled)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	transformed := []int{}
	for _, val := range *numbers {
		transformed = append(transformed, transform(val))
	}
	return transformed
}

func createTransformer(factor int) transformFn {
	return func(number int) int {
		return number * factor
	}
}
