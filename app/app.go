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

	rtr := mux.NewRouter()
	rtr.HandleFunc("/", controllers.RootHandler)
	rtr.HandleFunc("/patients", controllers.PatientsIndexHandler)
	rtr.HandleFunc("/patients/{id:[0-9]+}", controllers.PatientsShowHandler)

	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)
}