package main

import (
	"github.com/0xTheFr34k/go-net-http/router"
	"net/http"
)

func main() {
	appMux := http.NewServeMux()

	appMux.Handle("/v1/", http.StripPrefix("/v1", router.V1Router()))

	http.ListenAndServe(":8000", appMux)
}
