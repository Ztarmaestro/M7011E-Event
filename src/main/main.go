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
//templates = template.Must(template.ParseGlob("Event/templates/*"))

func main() {
	// Instantiate a new router

	//Real address for server, change back before pushing to git 
	//bindAddr := "192.168.1.82:8080"

	//Address for testing server on LAN
	bindAddr := "127.0.0.1:8080"

	r := httprouter.New()
	r.NotFound = http.FileServer(http.Dir("Event/"))
	r.GET("/", indexHandler)
	r.GET("/events", eventHandler)
	fmt.Println("Server running on", bindAddr)
	log.Fatal(http.ListenAndServe(bindAddr, r))
}

/*
func indexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

    // you access the cached templates with the defined name, not the filename
    err := templates.ExecuteTemplate(w, "main", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
*/

func indexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

    // you access the cached templates with the defined name, not the filename

    pagePath := "Event/templates/main.html"
	pageFooter := "Event/templates/footer.html"
	pageNavbar := "Event/templates/navbar.html"

	pageTemplate := "Event/templates/startpage.html"

	if t, err := template.ParseFiles(pagePath, pageFooter, pageNavbar, pageTemplate); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}

}

func eventHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

    // you access the cached templates with the defined name, not the filename

    pagePath := "Event/templates/main.html"
	pageFooter := "Event/templates/footer.html"
	pageNavbar := "Event/templates/navbar.html"

	pageTemplate := "Event/templates/overview_events.html"

	if t, err := template.ParseFiles(pagePath, pageFooter, pageNavbar, pageTemplate); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}

}


