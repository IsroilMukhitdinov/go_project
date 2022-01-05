package scraper

import (
	"github.com/gocolly/colly"
)

func Scrap(nameQuery string, valueQuery string, url string) (*[]string, *[]string) {

	var names []string
	var values []string

	// var list []entities.Currency

	c := colly.NewCollector()

	c.OnHTML(nameQuery, func(h *colly.HTMLElement) {
		names = append(names, h.Text)
	})

	c.OnHTML(valueQuery, func(h *colly.HTMLElement) {
		values = append(values, h.Text)
	})

	// defer func() {
	// 	for i, name := range names {
	// 		list = append(list, entities.Currency{
	// 			CurrencyName:  name,
	// 			CurrencyValue: values[i],
	// 		})
	// 	}
	// }()

	c.Visit(url)

	return &names, &values
}
