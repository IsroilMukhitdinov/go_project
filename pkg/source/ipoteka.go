package source

import (
	"net/http"

	"github.com/IsroilMukhitdinov/go_project/data"
	"github.com/IsroilMukhitdinov/go_project/pkg/entities"
	"github.com/IsroilMukhitdinov/go_project/pkg/handler"
	"github.com/IsroilMukhitdinov/go_project/pkg/scraper"
)

func Ipoteka(response http.ResponseWriter, request *http.Request) {

	nameQuery := "div.complete > p"
	valueQuery := "div.purchase > span"
	url := "https://ipotekabank.uz"

	temp1, temp2 := scraper.Scrap(nameQuery, valueQuery, url)

	var temp []entities.Currency

	for i, name := range *temp1 {
		temp = append(temp, entities.Currency{
			CurrencyName:  name,
			CurrencyValue: (*temp2)[i+9],
		})
	}

	handler.Handle(response, entities.PageData{
		BankName: "Ipoteka Bank",
		List:     temp,
	})

	data.OpenAndWrite("ipoteka", &temp)
}
