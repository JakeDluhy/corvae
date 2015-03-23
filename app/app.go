package main

import (
	"fmt"
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
	// patient := models.Patient{Name: "Jake", Email: "jake.dluhy@gmail.com", Age: 22}
	// database.Create(&patient)

	http.HandleFunc("/", controllers.RootHandler)
	http.HandleFunc("/patients", controllers.PatientsIndexHandler)
	http.ListenAndServe(":8080", nil)
}