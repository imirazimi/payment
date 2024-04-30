package parser

import (
	"encoding/csv"
	"fmt"
	"os"
	entity "payment/entity/tax"
	"strconv"
)

func (p parser) ParseTaxes(path string) ([]entity.Tax, error) {
	file, oErr := os.Open(path)
	if oErr != nil {
		return []entity.Tax{}, fmt.Errorf("can not open file")
	}
	defer file.Close()
	rows, rErr := csv.NewReader(file).ReadAll()
	if rErr != nil {
		return []entity.Tax{}, fmt.Errorf("can not parse file")
	}
	taxes := make([]entity.Tax, 0)
	for _, row := range rows {
		country := row[0]
		percentage, cErr := strconv.ParseFloat(row[1], 64)
		if cErr != nil {
			return []entity.Tax{}, fmt.Errorf("can not parse file")
		}
		taxes = append(taxes, entity.Tax{
			Country:    country,
			Percentage: percentage,
		})
	}
	return taxes, nil
}
