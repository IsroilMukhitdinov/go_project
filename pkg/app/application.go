package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IsroilMukhitdinov/go_project/pkg/entities"
	"github.com/IsroilMukhitdinov/go_project/pkg/handler"
	"github.com/IsroilMukhitdinov/go_project/pkg/scraper"
)

const port = ":7890"

func main() {
	nameQuery := "div.official-row > div:first-child"
	valueQuery := "div.official-row > div > span.currency-value"
	url := "https://pultop.uz"

	temp := scraper.Scrap(nameQuery, valueQuery, url)

	fmt.Println(*temp)
	log.Println("Started Listening on Port ", port)

	http.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
		handler.Handle(rw, entities.PageData{
			BankName: "National Bank",
			List:     *temp,
		})
	})

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
