package lib

import (
	"fmt"

	"github.com/gocolly/colly"
)

// this function get the dividend by code and return the dividend and the price
func GetDividendByCode(code string) (string, string) {
	var dividendText string
	var price string
	c := colly.NewCollector()

	c.OnHTML("div.headerTicker__content__price p", func(h *colly.HTMLElement) {
		price = h.Text
	})
	c.OnHTML("div.indicators.historic div.indicators__box:nth-of-type(3) p:nth-of-type(2)", func(h *colly.HTMLElement) {
		dividendText = h.Text
	})
	url := fmt.Sprintf("https://www.fundsexplorer.com.br/funds/%s", code)
	c.Visit(url)
	return dividendText, price
}
