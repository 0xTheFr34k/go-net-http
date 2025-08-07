package router

import (
	"github.com/0xTheFr34k/go-net-http/handler/auth"
	"github.com/0xTheFr34k/go-net-http/middleware"
	"net/http"
)

func AuthRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", auth.Login)
	mux.HandleFunc("/register", auth.Register)

	stack := middleware.CreateStack(
		middleware.AuthMiddleware,
		middleware.Logging,
	)

	return stack(mux)
}
