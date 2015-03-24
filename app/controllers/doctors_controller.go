package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	// "fmt"
	"corvae/app/models"
	"corvae/db"
)

func DoctorsIndexHandler(res http.ResponseWriter, req *http.Request) {
	database := db.GetDB()
	doctors := database.Find(&models.Doctor{})

  js, err := json.Marshal(doctors)
  if err != nil {
    http.Error(res, err.Error(), http.StatusInternalServerError)
    return
  }

  res.Header().Set("Content-Type", "application/json")
  res.Write(js)
}

func DoctorsShowHandler(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	database := db.GetDB()
	doctor := database.Find(&models.Doctor{}, id)

	js, err := json.Marshal(doctor)
  if err != nil {
    http.Error(res, err.Error(), http.StatusInternalServerError)
    return
  }

  res.Header().Set("Content-Type", "application/json")
  res.Write(js)
}

// func DoctorPatientsHandler(res http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	id := params["id"]
// 	database := db.GetDB()
// 	patients := 
// }