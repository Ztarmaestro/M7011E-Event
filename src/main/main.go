package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	//"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

    // you access the cached templates with the defined name, not the filename

    pagePath := "Event/templates/main.html"

	pageTemplate := "Event/templates/startpage.html"

	if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}}

func eventHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

    // you access the cached templates with the defined name, not the filename

    pagePath := "Event/templates/main.html"

	pageTemplate := "Event/templates/overview_events.html"
	pageSidemeny := "Event/templates/sidemeny.html"
	pageEventbutton := "Event/templates/create_event_form.html"
	
	if t, err := template.ParseFiles(pagePath, pageTemplate, pageSidemeny, pageEventbutton); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}}

func profileHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

    // you access the cached templates with the defined name, not the filename

    pagePath := "Event/templates/main.html"

	pageTemplate := "Event/templates/profile.html"

	if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}}

func showroomHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

    // you access the cached templates with the defined name, not the filename

    pagePath := "Event/templates/main.html"

	pageSidemeny := "Event/templates/sidemeny.html"
	pageTemplate := "Event/templates/view_event.html"
	//pageMapbutton := "Event/templates/map.html"
	

	if t, err := template.ParseFiles(pagePath, pageMapbutton, pageSidemeny, pageTemplate); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}}

func aboutHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

    // you access the cached templates with the defined name, not the filename

    pagePath := "Event/templates/main.html"

	pageTemplate := "Event/templates/about.html"

	if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}}

func searchHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

    // you access the cached templates with the defined name, not the filename

    pagePath := "Event/templates/main.html"

	pageTemplate := "Event/templates/search_result.html"

	if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}}

func eventHandler_nologin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

    // you access the cached templates with the defined name, not the filename

    pagePath := "Event/templates/main.html"

	pageTemplate := "Event/templates/overview_events_nologin.html"
	pageSidemeny := "Event/templates/sidemeny_nologin.html"

	
	if t, err := template.ParseFiles(pagePath, pageTemplate, pageSidemeny); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}}

func main() {
	// Instantiate a new router

	//Real address for server, change back before pushing to git 
	bindAddr := "130.240.170.56:8080"

	//Address for testing server on LAN
	//bindAddr := "127.0.0.1:8050"

	r := httprouter.New()
	r.NotFound = http.FileServer(http.Dir("Event/"))

	//Handlers for differnt pages
	r.GET("/", indexHandler)
	r.GET("/events", eventHandler)
    //r.GET("/profile", profileHandler)
    r.GET("/show_event", showroomHandler)
    r.GET("/about", aboutHandler)
    //r.GET("/search_result", searchHandler)
    r.GET("/events_nologin", eventHandler_nologin)

	fmt.Println("Server running on", bindAddr)
	log.Fatal(http.ListenAndServe(bindAddr, r)) }



