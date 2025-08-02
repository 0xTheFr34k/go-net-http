package user

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	_id := r.PathValue("id")

	id, err := strconv.Atoi(_id)
	if err != nil || id < 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Your ID is %s\n", _id)
}

func Panic(w http.ResponseWriter, r *http.Request) {
	var tmp *int
	*tmp += 1
}
