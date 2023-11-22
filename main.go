package main

import (
	"fmt"
	"strconv"
	"strings"

	"scrapp.com/mod/lib"
	"scrapp.com/mod/types"
)

func main() {
	//pegar lista de dados do IFIX do b3
	taxaIpcaPorcentage := 5.51
	riskRatePorcentage := 2.5
	descountTax := (taxaIpcaPorcentage + riskRatePorcentage) / 100 //0,0826
	ifixList := lib.GetB3Ifixdata()
	var data []types.Fund
	for _, ifix := range ifixList {
		code := ifix.Code
		//get dividend data from status invest
		dvidend, price := lib.GetDividendByCode(code)

		dvidend = strings.ReplaceAll(dvidend, "R$ ", "")
		dvidend = strings.ReplaceAll(dvidend, ",", ".")

		price = strings.ReplaceAll(price, "R$ ", "")
		price = strings.ReplaceAll(price, ",", ".")
		price = strings.TrimSpace(price)

		dvidendNumber, err := strconv.ParseFloat(dvidend, 64)
		if err != nil {
			fmt.Println("Error occurred while parsing:", err)
			return
		}
		priceNumber, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Error occurred while parsing:", err)
			return
		}

		maxPrice := dvidendNumber / descountTax
		dif := maxPrice - priceNumber
		if dif > 5 {
			fmt.Println("dif", code, dvidendNumber, price, maxPrice, maxPrice-priceNumber)
		}
		fund := types.Fund{code, dvidend, priceNumber, maxPrice}
		data = append(data, fund)
	}
	filename := "funds.csv"
	err := lib.CreateCSVFromFunds(data, filename)
	if err != nil {
		println("Error creating CSV:", err)
		return
	}

	println("CSV file created successfully:", filename)
	// Read the response body

}
