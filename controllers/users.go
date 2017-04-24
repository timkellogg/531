package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

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
		fmt.Println(errors)
		json.NewEncoder(w).Encode(errors)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	saved := user.SaveUser()
	if saved != nil {
		fmt.Println(saved)
		json.NewEncoder(w).Encode(saved)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}
