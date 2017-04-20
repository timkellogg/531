package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/timkellogg/531/server/config"
	"github.com/timkellogg/531/server/models"
)

// UsersCreate - creates a new user
func UsersCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	errors := user.ValidateUser()
	if len(errors) > 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(errors)
	}

	config.Database.DB.Create(&user)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}
