package main

import (
	// Standard library packages
	//"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	//"github.com/gorilla/mux"
)

// compile all templates and cache them
var templates = template.Must(template.ParseGlob("Event/templates/*"))

func main() {
	// Instantiate a new router
	bindAddr := "192.168.1.82:8080"
	r := httprouter.New()
	r.NotFound = http.FileServer(http.Dir("static/"))
	r.GET("/", indexHandler)
	fmt.Println("Server running on", bindAddr)
	log.Fatal(http.ListenAndServe(bindAddr, r))
}

func indexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

    // you access the cached templates with the defined name, not the filename
    err := templates.ExecuteTemplate(w, "indexPage", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}


