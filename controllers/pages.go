package controllers

import (
	"html/template"
	"net/http"
)

// PagesHome - "/" index route handler for index
func PagesHome(w http.ResponseWriter, h *http.Request) {
	d := PageData{Title: "531 Trainer"}

	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, d)
}
