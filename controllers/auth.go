package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/timkellogg/531/server/models"
)

// AuthLogin - creates a new user cookie for signed in users
func AuthLogin(w http.ResponseWriter, r *http.Request) {
	var user models.User

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	found, record := user.FindUser()

	if found {
		session, err := models.Encrypt(record.Username + os.Getenv("SALT"))

		cookie := http.Cookie{Name: "fivethreeonesession", Value: session, Expires: time.Now().Add(time.Hour), HttpOnly: true, MaxAge: 50000, Path: "/"}
		if err != nil {
			panic(err)
		}

		http.SetCookie(w, &cookie)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&record)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(nil)
	}
}

// AuthLogout - deletes a cookie for signed in users
func AuthLogout(w http.ResponseWriter, r *http.Request) {
	// set ttl to expired date (ie, yesterday)
	// set the cookie contents to be rubbish
}
