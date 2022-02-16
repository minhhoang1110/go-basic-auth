package handler

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/minhhoang1110/go-basic-auth/db"
	"github.com/minhhoang1110/go-basic-auth/models"
)

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	var newUser models.User
	if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"}, "error")
		return
	}
	if govalidator.IsNull(newUser.Email) || govalidator.IsNull(newUser.Password) {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Data can not empty"}, "error")
		return
	}

	if !govalidator.IsEmail(newUser.Email) {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Email is invalid"}, "error")
		return
	}
	db, err := db.GetConnect()
	if err != nil {
		responseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "Connect error"}, "error")
		return
	}
	newUser.Password, err = newUser.Hash(newUser.Password)
	if err != nil {
		responseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "Can not Encode password"}, "error")
		return
	}
	result := db.Create(&newUser)
	if result.Error != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": result.Error.Error()}, "error")
		return
	}
	responseWithJson(writer, http.StatusOK, newUser, "user")
}
