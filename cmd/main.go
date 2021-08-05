package main

import (
	"log"
	"net/http"

	gif "github.com/fabulousginger/randocats/gif"
)

func main() {

	gif.GetGif()
	http.HandleFunc("/", gif.WebHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
