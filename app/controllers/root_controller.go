package controllers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var templates = template.Must(template.ParseFiles(filepath.Join("assets","index.html")))

func RootHandler(res http.ResponseWriter, req *http.Request) {
	err := templates.ExecuteTemplate(res, "index.html", nil)
  if err != nil {
  	http.Error(res, err.Error(), http.StatusInternalServerError)
  }
}