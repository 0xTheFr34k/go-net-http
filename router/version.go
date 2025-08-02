package router

import "net/http"

func V1Router() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/auth/", http.StripPrefix("/auth", AuthRouter()))
	mux.Handle("/user/", http.StripPrefix("/user", UserRouter()))
	return mux
}
