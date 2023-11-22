package lib

import (
	"encoding/csv"
	"os"
	"strconv"

	"scrapp.com/mod/types"
)

func CreateCSVFromFunds(funds []types.Fund, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header row
	header := []string{"Code", "Yield", "Price", "MaxPrice"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write data rows
	for _, fund := range funds {
		row := []string{
			fund.Code,
			fund.Yield,
			// Convert float64 to string for Price and MaxPrice
			strconv.FormatFloat(fund.Price, 'f', -1, 64),
			strconv.FormatFloat(fund.MaxPrice, 'f', -1, 64),
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
