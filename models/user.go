package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email" gorm:"primaryKey"`
	Password string `json:"password"`
}

func (u *User) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *User) CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}
