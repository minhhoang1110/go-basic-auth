package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minhhoang1110/go-basic-auth/apis"
)

func main() {
	r := mux.NewRouter()
	apis.UserApis(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
