package controllers

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/timkellogg/531/server/models"
)

// UsersCreate - creates a new user
func UsersCreate(w http.ResponseWriter, r *http.Request) {
	var user models.User

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user.Email)
	fmt.Println(user.Username)
}
