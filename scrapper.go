package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Fund struct {
	Code     string
	Yield    string
	Price    float64
	MaxPrice float64
}

func main() {
	//pegar lista de dados do IFIX do b3
	taxaIpcaPorcentage := 5.51
	riskRatePorcentage := 2.5
	descountTax := (taxaIpcaPorcentage + riskRatePorcentage) / 100 //0,0826
	ifixList := L.GetB3Ifixdata()
	var data []Fund
	for _, ifix := range ifixList {
		code := ifix.Code
		//get dividend data from status invest
		dvidend, price := L.getDividendByCode(code)

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
		fund := Fund{code, dvidend, priceNumber, maxPrice}
		data = append(data, fund)
	}
	// Read the response body

}
