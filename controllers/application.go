package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PageData - represents a struct of all the data delivered to the template
type PageData struct {
	Title string
}

// HandleError - handles api errors
func HandleError(w http.ResponseWriter, errors []error, status int) {
	fmt.Println(errors)
	json.NewEncoder(w).Encode(errors)
	w.WriteHeader(status)
}
