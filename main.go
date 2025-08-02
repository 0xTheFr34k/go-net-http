package main

import (
	"net/http"

	"github.com/0xTheFr34k/go-net-http/handler/auth"
	"github.com/0xTheFr34k/go-net-http/handler/user"
)

func main() {
	mainMux := http.NewServeMux()

	authMux := http.NewServeMux()
	authMux.HandleFunc("/login", auth.Login)

	userMux := http.NewServeMux()
	userMux.HandleFunc("/user/{id}", user.GetUserByID)

	authMux.Handle("/", userMux)

	mainMux.Handle("/", authMux)

	http.ListenAndServe(":8000", mainMux)
}
