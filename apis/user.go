package apis

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minhhoang1110/go-basic-auth/handler"
)

func UserApis(r *mux.Router) {
	r.HandleFunc("/api/user", handler.CreateUser).Methods(http.MethodPost)
}
