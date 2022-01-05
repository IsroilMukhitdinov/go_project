package main

import (
	"log"
	"net/http"

	"github.com/IsroilMukhitdinov/go_project/pkg/source"
)

const port = ":7890"

func main() {

	http.HandleFunc("/pultop", source.Pultop)
	http.HandleFunc("/ipoteka", source.Ipoteka)

	log.Println("Started Listening on Port ", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
