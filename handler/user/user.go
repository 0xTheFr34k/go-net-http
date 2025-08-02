package user

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Response struct {
	Message    string `json:"message"`
	CodeStatus int    `json:"code"`
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	_id := r.PathValue("id")

	id, err := strconv.Atoi(_id)

	w.Header().Set("Content-Type", "application/json")

	if err != nil || id < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			Response{
				Message:    "Error",
				CodeStatus: http.StatusBadRequest,
			})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Message:    "Success",
		CodeStatus: http.StatusOK,
	})
}

func Panic(w http.ResponseWriter, r *http.Request) {
	var tmp *int
	*tmp += 1
}
