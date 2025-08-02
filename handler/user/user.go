package user

import (
	"encoding/json"
	"github.com/0xTheFr34k/go-net-http/dto"
	"net/http"
	"strconv"
)

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	_id := r.PathValue("id")

	id, err := strconv.Atoi(_id)

	w.Header().Set("Content-Type", "application/json")

	if err != nil || id < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			UserDto.StatusResponse{
				Message:    "Error",
				CodeStatus: http.StatusBadRequest,
			})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		UserDto.StatusResponse{
			Message:    "Success",
			CodeStatus: http.StatusOK,
		})
}

func Panic(w http.ResponseWriter, r *http.Request) {
	var tmp *int
	*tmp += 1
}
