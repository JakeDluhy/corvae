package main

import (
	"fmt"
	"github.com/gorilla/mux"
	// "flag"
 //  "io/ioutil"
  // "log"
  // "net"
  "net/http"
  // "regexp"
  "corvae/db"
  // "corvae/app/models"
  "corvae/app/controllers"
)




func main() {
	fmt.Println("Server up and running at localhost:8080")
	fmt.Println("Initializing Database")
	db.Start()
	fmt.Println("Database up and running")

	// database := db.GetDB()
	// pat := models.Patient{}
	// database.Find(&pat, 2)
	// fmt.Println(pat.LoginPatient("fooba"))

	rtr := mux.NewRouter()
	// Root
	rtr.HandleFunc("/", controllers.RootHandler)

	// Patients
	rtr.HandleFunc("/patients", controllers.PatientsIndexHandler)
	rtr.HandleFunc("/patients/{id:[0-9]+}", controllers.PatientsShowHandler)

	// Doctors
	rtr.HandleFunc("/doctors", controllers.DoctorsIndexHandler)
	rtr.HandleFunc("/doctors/{id:[0-9]+}", controllers.DoctorsShowHandler)

	// Nurse Station
	rtr.HandleFunc("/nurse_stations", controllers.NurseStationsIndexHandler)
	rtr.HandleFunc("/nurse_stations/{id:[0-9]+}", controllers.NurseStationsShowHandler)

	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)
}