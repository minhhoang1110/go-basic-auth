package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minhhoang1110/go-basic-auth/apis"
)

func main() {
	r := mux.NewRouter()
	// dsn := "host=localhost user=root password=123 dbname=quanlynhanvien port=5432 sslmode=disable TimeZone=Asia/Saigon"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("Fail connect database !!!!!!!!!")
	// 	return
	// }
	apis.UserApis(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
