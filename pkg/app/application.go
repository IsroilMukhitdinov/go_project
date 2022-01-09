package main

import (
	"log"
	"net/http"

	"github.com/IsroilMukhitdinov/go_project/pkg/source"
)

const port = ":7890"

func main() {

	// http.HandleFunc("/home", func(rw http.ResponseWriter, r *http.Request) {
	// 	parsed, err := template.ParseFiles("./templates/home.html")

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	parsed.Execute(rw, nil)
	// })

	http.HandleFunc("/pultop", source.Pultop)
	http.HandleFunc("/ipoteka", source.Ipoteka)
	http.HandleFunc("/agro", source.Agrobank)
	http.HandleFunc("/ravnaq", source.Ravnaqbank)

	log.Println("Started Listening on Port ", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
