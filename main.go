package main

import (
	"github.com/0xTheFr34k/go-net-http/handler/auth"
	"github.com/0xTheFr34k/go-net-http/handler/user"
	"github.com/0xTheFr34k/go-net-http/middleware"
	"net/http"
)

func main() {

	appMux := http.NewServeMux()

	authMux := http.NewServeMux()
	authMux.HandleFunc("/login", auth.Login)

	authstack := middleware.CreateStack(
		middleware.AuthMiddleware,
		middleware.Logging,
	)

	appMux.Handle("/auth/", http.StripPrefix("/auth", authstack(authMux)))

	userMux := http.NewServeMux()
	userMux.HandleFunc("/{id}", user.GetUserByID)

	userstack := middleware.CreateStack(
		middleware.AuthMiddleware,
		middleware.Logging,
	)

	appMux.Handle("/user/", http.StripPrefix("/user", userstack(userMux)))

	versionMux := http.NewServeMux()

	versionMux.Handle("/v1/", http.StripPrefix("/v1", appMux))

	http.ListenAndServe(":8000", versionMux)
}
