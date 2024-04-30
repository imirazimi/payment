package parser

import (
	"encoding/csv"
	"fmt"
	"os"
	entity "payment/entity/purchase"
	"strconv"
)

func (p parser) ParseItems(path string) ([]entity.Item, error) {
	file, oErr := os.Open(path)
	if oErr != nil {
		return []entity.Item{}, fmt.Errorf("can not open file")
	}
	defer file.Close()
	rows, rErr := csv.NewReader(file).ReadAll()
	if rErr != nil {
		return []entity.Item{}, fmt.Errorf("can not parse file")
	}
	items := make([]entity.Item, 0)
	for _, row := range rows {
		id64, cErr := strconv.ParseUint(row[0], 10, 64)
		if cErr != nil {
			return []entity.Item{}, fmt.Errorf("can not parse file")
		}
		id := uint(id64)
		price, cErr := strconv.ParseFloat(row[1], 64)
		if cErr != nil {
			return []entity.Item{}, fmt.Errorf("can not parse file")
		}
		items = append(items, entity.Item{
			ID:    id,
			Price: price,
		})
	}
	return items, nil
}
