package source

import (
	"net/http"

	"github.com/IsroilMukhitdinov/go_project/data"
	"github.com/IsroilMukhitdinov/go_project/pkg/entities"
	"github.com/IsroilMukhitdinov/go_project/pkg/handler"
	"github.com/IsroilMukhitdinov/go_project/pkg/scraper"
)

func Ravnaqbank(response http.ResponseWriter, request *http.Request) {

	nameQuery := "div.header__exchange_item > strong"
	valueQuery := "div.header__exchange_item > span"
	url := "https://www.ravnaqbank.uz/en/services/exchange-rates/"

	temp1, temp2 := scraper.Scrap(nameQuery, valueQuery, url)

	var temp []entities.Currency

	for i, name := range *temp1 {
		temp = append(temp, entities.Currency{
			CurrencyName:  name,
			CurrencyValue: (*temp2)[i],
		})
	}

	handler.Handle(response, entities.PageData{
		BankName: "Ravnaq Bank",
		List:     temp,
	})

	data.OpenAndWrite("ravnaq", &temp)

}
