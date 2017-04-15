package controllers

import (
	"html/template"
	"net/http"
)

// PagesHome - "/" index route handler for application
func PagesHome(w http.ResponseWriter, h *http.Request) {
	d := PageData{Title: "531 Trainer"}

	t, _ := template.ParseFiles("views/application.html")
	t.Execute(w, d)
}
