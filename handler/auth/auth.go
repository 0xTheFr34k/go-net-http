package auth

import "net/http"

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to login\n"))
}
