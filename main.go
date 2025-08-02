package main

import (
	"net/http"

	"github.com/0xTheFr34k/go-net-http/handler/auth"
	// "github.com/0xTheFr34k/go-net-http/handler/user"
	"github.com/0xTheFr34k/go-net-http/middleware"
)

func main() {

	authMux := http.NewServeMux()

	stack := middleware.CreateStack(
		middleware.AuthMiddleware,
		middleware.Logging,
	)
	authMux.HandleFunc("/login", auth.Login)
	http.ListenAndServe(":8000", stack(authMux))
}
