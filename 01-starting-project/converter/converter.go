package converter

import (
	"errors"
	"strconv"
)

type NumberSlice interface {
	[]float64 | []int
}

func ConvertStringFloat(strs []string) ([]float64, error) {
	var values []float64

	for _, str := range strs {
		floatVal, err := strconv.ParseFloat(str, 64)

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}
		values = append(values, floatVal)
	}
	return values, nil
}
