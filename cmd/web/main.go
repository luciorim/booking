package main

import (
	"log"
	"net/http"

	"github.com/luciorim/booking/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Fatal(http.ListenAndServe(":4000", nil))

}
