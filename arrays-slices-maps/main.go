package main

import "fmt"

type floatMap map[string]float64

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Jackson"
	userNames[1] = "Josha"

	userNames = append(userNames, "Jacob")
	userNames = append(userNames, "Daryl")

	courseRatings := make(floatMap, 5)

	courseRatings["go"] = 4.5
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	for _, rating := range courseRatings {
		fmt.Println(rating)
	}
}
