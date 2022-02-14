package handler

import (
	"encoding/json"
	"net/http"

	"github.com/minhhoang1110/go-basic-auth/db"
	"github.com/minhhoang1110/go-basic-auth/models"
)

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	var newUser models.User
	if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"}, nil)
		return
	}
	db, err := db.GetConnect()
	if err != nil {
		responseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "Connect error"}, nil)
		return
	}
	result := db.Create(&newUser)
	if result.Error != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Create error"}, nil)
		return
	}
	responseWithJson(writer, http.StatusOK, newUser, "user")
}
