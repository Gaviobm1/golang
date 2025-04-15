package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"example.com/price-calculator/converter"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	Prices            []float64
	TaxIncludedPrices PriceTaxedPriceMap
}

type PriceTaxedPriceMap map[string]float64

func (t *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Error occurred")
		return
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Read file content failed")
		fmt.Println(err.Error())
		file.Close()
		return
	}

	values, err := converter.ConvertStringFloat(lines)

	if err != nil {
		fmt.Println(err.Error())
		file.Close()
		return
	}

	file.Close()
	t.Prices = values
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

func (t *TaxIncludedPriceJob) Process() {
	t.LoadData()
	t.TaxIncludedPrices = CalculateTaxedPrice(t.Prices, t.TaxRate)
	fmt.Println(t.TaxIncludedPrices)
}

func CalculateTaxedPrice(prices []float64, taxRate float64) PriceTaxedPriceMap {
	result := make(PriceTaxedPriceMap)
	for _, price := range prices {
		taxIncludedPrice := price * (1 + taxRate)
		taxIncludedFloat, err := strconv.ParseFloat(fmt.Sprintf("%.2f", taxIncludedPrice), 64)

		if err != nil {
			fmt.Println("Error converting string to float")
			fmt.Println(err.Error())
			return nil
		}

		result[fmt.Sprintf("%.2f", price)] = taxIncludedFloat
	}
	return result
}
