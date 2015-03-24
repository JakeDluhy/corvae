package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"reflect"
	"corvae/app/models"
	"corvae/db"
)

func PatientsIndexHandler(res http.ResponseWriter, req *http.Request) {
	database := db.GetDB()
	patients := database.Find(&models.Patient{})

  js, err := json.Marshal(patients)
  if err != nil {
    http.Error(res, err.Error(), http.StatusInternalServerError)
    return
  }

  res.Header().Set("Content-Type", "application/json")
  res.Write(js)
}

func PatientsShowHandler(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	database := db.GetDB()
	patient := database.Find(&models.Patient{}, id)

	js, err := json.Marshal(patient)
	fmt.Println(reflect.TypeOf(js))
	fmt.Println(js)
  if err != nil {
    http.Error(res, err.Error(), http.StatusInternalServerError)
    return
  }

  res.Header().Set("Content-Type", "application/json")
  res.Write(js)
}