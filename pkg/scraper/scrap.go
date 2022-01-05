package scraper

import (
	"github.com/IsroilMukhitdinov/go_project/pkg/entities"
	"github.com/gocolly/colly"
)

func Scrap(nameQuery string, valueQuery string, url string) *[]entities.Currency {

	var names []string
	var values []string

	var list []entities.Currency

	c := colly.NewCollector(colly.AllowedDomains("pultop.uz"))

	c.OnHTML(nameQuery, func(h *colly.HTMLElement) {
		names = append(names, h.Text)
	})

	c.OnHTML(valueQuery, func(h *colly.HTMLElement) {
		values = append(values, h.Text)
	})

	defer func() {
		for i, name := range names {
			list = append(list, entities.Currency{
				CurrencyName:  name,
				CurrencyValue: values[i],
			})
		}
	}()

	c.Visit(url)

	return &list
}
