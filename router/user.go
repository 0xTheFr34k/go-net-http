package router

import (
	"github.com/0xTheFr34k/go-net-http/handler/user"
	"github.com/0xTheFr34k/go-net-http/middleware"
	"net/http"
)

func UserRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/{id}", user.GetUserByID)
	mux.HandleFunc("/panic", user.Panic)

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.RecoveryMiddlerware,
	)

	return stack(mux)
}
