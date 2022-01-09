package source

import (
	"net/http"

	"github.com/IsroilMukhitdinov/go_project/data"
	"github.com/IsroilMukhitdinov/go_project/pkg/entities"
	"github.com/IsroilMukhitdinov/go_project/pkg/handler"
	"github.com/IsroilMukhitdinov/go_project/pkg/scraper"
)

func Agrobank(response http.ResponseWriter, request *http.Request) {

	nameQuery := "tr > td"
	valueQuery := "tr > td"
	url := "https://agrobank.uz/uz/exchange_rates"

	temp1, temp2 := scraper.Scrap(nameQuery, valueQuery, url)

	var temp []entities.Currency

	var names []string
	for i, name := range *temp1 {
		if i%4 == 1 {
			names = append(names, name)
		}
	}

	var values []string
	for i, value := range *temp2 {
		if i%4 == 2 {
			values = append(values, value)
		}
	}

	for i, name := range names {
		temp = append(temp, entities.Currency{
			CurrencyName:  name,
			CurrencyValue: values[i],
		})
	}

	handler.Handle(response, entities.PageData{
		BankName: "AgroBank",
		List:     temp,
	})

	data.OpenAndWrite("agro", &temp)

}
