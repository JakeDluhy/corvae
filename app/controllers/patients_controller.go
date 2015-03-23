package controllers

import (
	"encoding/json"
	"net/http"
	// "fmt"
	"corvae/app/models"
	"corvae/db"
)

func PatientsIndexHandler(res http.ResponseWriter, req *http.Request)  {
	database := db.GetDB()
	pat := database.First(&models.Patient{})

  js, err := json.Marshal(pat)
  if err != nil {
    http.Error(res, err.Error(), http.StatusInternalServerError)
    return
  }

  res.Header().Set("Content-Type", "application/json")
  res.Write(js)
}