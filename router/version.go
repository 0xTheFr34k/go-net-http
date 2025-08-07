package router

import (
	"net/http"

	"github.com/0xTheFr34k/go-net-http/middleware"
)

func V1Router() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/auth/", http.StripPrefix("/auth", AuthRouter()))
	mux.Handle("/user/", http.StripPrefix("/user", UserRouter()))

	stack := middleware.CreateStack(
		middleware.WithWaitGroup,
	)

	return stack(mux)
}
