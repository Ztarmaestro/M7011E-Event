func main() {
	// Instantiate a new router

	//Real address for server, change back before pushing to git 
	bindAddr := "130.240.170.56:8080"

	//Address for testing server on LAN
	//bindAddr := "127.0.0.1:8080"

	r := httprouter.New()
	r.NotFound = http.FileServer(http.Dir("Event/"))

	//Handlers for differnt pages
	r.GET("/", indexHandler)
	r.POST("/events", eventHandler)
    r.GET("/profile", profileHandler)
    r.GET("/create_event", createHandler)
    r.GET("/about", aboutHandler)
    r.GET("/search_result", searchHandler)

	fmt.Println("Server running on", bindAddr)
	log.Fatal(http.ListenAndServe(bindAddr, r))
}