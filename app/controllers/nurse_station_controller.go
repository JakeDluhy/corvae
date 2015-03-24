package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	// "fmt"
	"corvae/app/models"
	"corvae/db"
)

func NurseStationsIndexHandler(res http.ResponseWriter, req *http.Request) {
	database := db.GetDB()
	stations := database.Find(&models.NurseStation{})

  js, err := json.Marshal(stations)
  if err != nil {
    http.Error(res, err.Error(), http.StatusInternalServerError)
    return
  }

  res.Header().Set("Content-Type", "application/json")
  res.Write(js)
}

func NurseStationsShowHandler(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	database := db.GetDB()
	station := database.Find(&models.NurseStation{}, id)

	js, err := json.Marshal(station)
  if err != nil {
    http.Error(res, err.Error(), http.StatusInternalServerError)
    return
  }

  res.Header().Set("Content-Type", "application/json")
  res.Write(js)
}

// func NurseStationPatientsHandler(res http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	id := params["id"]
// 	database := db.GetDB()
// 	patients := 
// }