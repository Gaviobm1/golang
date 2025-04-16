package cmdmanager

import "fmt"

type CMDManager struct{}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm each price with ENTER")

	var prices []string

	for {
		var price string
		fmt.Println("Enter a price:")
		fmt.Scan(&price)
		prices = append(prices, price)

		if price == "0" {
			break
		}
	}

	return prices, nil
}

func (cmd CMDManager) WriteJSON(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
