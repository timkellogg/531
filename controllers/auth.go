package controllers

import (
	"encoding/json"
	"net/http"

	"os"

	"github.com/timkellogg/531/server/models"
)

// AuthLogin - creates a new user cookie for signed in users
func AuthLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	// create a cookie
	found, record := user.FindUser()

	if found {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&record)
	}

	// TODO: fixup compile error
	session := models.Encrypt(record.Username + os.Getenv("SALT"))
	cookie := http.Cookie{Name: "531Session", Value: session}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(nil)
}

// AuthLogout - deletes a cookie for signed in users
func AuthLogout(w http.ResponseWriter, r *http.Request) {
	// set ttl to expired date (ie, yesterday)
	// set the cookie contents to be rubbish
}
