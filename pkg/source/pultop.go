package source

import (
	"net/http"

	"github.com/IsroilMukhitdinov/go_project/pkg/entities"
	"github.com/IsroilMukhitdinov/go_project/pkg/handler"
	"github.com/IsroilMukhitdinov/go_project/pkg/scraper"
)

func Pultop(response http.ResponseWriter, request *http.Request) {

	nameQuery := "div.official-row > div:first-child"
	valueQuery := "div.official-row > div > span.currency-value"
	url := "https://pultop.uz"

	temp1, temp2 := scraper.Scrap(nameQuery, valueQuery, url)

	var temp []entities.Currency

	for i, name := range *temp1 {
		temp = append(temp, entities.Currency{
			CurrencyName:  name,
			CurrencyValue: (*temp2)[i],
		})
	}

	handler.Handle(response, entities.PageData{
		BankName: "National Bank",
		List:     temp,
	})
}
