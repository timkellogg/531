package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/timkellogg/531/server/models"
)

// UsersCreate - creates a new user
func UsersCreate(w http.ResponseWriter, r *http.Request) {
	var user models.User

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	fmt.Println("TOP")
	// fmt.Println(json.Unmarshal(body))
	fmt.Println([]byte(body), &user)
	fmt.Println("BOTTOM")

	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
}
