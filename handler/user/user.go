package user

import "net/http"

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Your ID is " + id + "\n"))
}
