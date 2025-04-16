package prices

import (
	"fmt"
	"strconv"

	"example.com/price-calculator/converter"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"taxRate"`
	Prices            []float64           `json:"prices"`
	TaxIncludedPrices PriceTaxedPriceMap  `json:"taxIncludedPrices"`
	IOManager         iomanager.IOManager `json:"-"`
}

type PriceTaxedPriceMap map[string]float64

func (t *TaxIncludedPriceJob) LoadData() error {
	lines, err := t.IOManager.ReadLines()

	if err != nil {
		return err
	}

	values, err := converter.ConvertStringFloat(lines)

	if err != nil {
		return err
	}

	t.Prices = values
	return nil
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:   taxRate,
		IOManager: iom,
	}
}

func (t *TaxIncludedPriceJob) Process() error {
	err := t.LoadData()

	if err != nil {
		return err
	}

	t.TaxIncludedPrices = CalculateTaxedPrice(t.Prices, t.TaxRate)
	t.IOManager.WriteJSON(t)
	return nil
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
