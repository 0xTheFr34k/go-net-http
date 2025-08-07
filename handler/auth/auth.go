package auth

import (
	"fmt"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/v1/auth/register", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 5)
	fmt.Fprintf(w, "Register complete\n")
}
