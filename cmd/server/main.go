// Server instance for routing API endpoints.
package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/ping", _)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
