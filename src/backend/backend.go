package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/ziutek/mymysql/godrv"
	//  "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/native" // Native engine
    // _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
	"io/ioutil"
	"log"
	//"math"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	//"time"

	// this is test for photos
	"bytes"
	"encoding/base64"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"strings"

)

// error response struct
/*
A error if composed of
	Error
	Message	of what error it is
	Code	errorcode
*/
type handlerError struct {
	Error   error
	Message string
	Code    int
}

// user struct

/*
A  User is composed of a
	UserID witch we set
	IdToken to that persons facebook/google
	Name of that user
	and a Photo witch is a url to that persons fb profile pic
*/
type User_table struct {
	UserID    	uint64 `json:"UserID"`
	IdToken   	string `json:"IdToken"`
	Name 		string `json:"Username"`
	//LastName  string `json:"last_name"`
	//Attending	uint64 `json:"Attending"`
	//Photo     string `json:"photo"`
}

//Event struct
/*
A Event is composed of
	Id			witch we set
	Date 	 	of the event
	Address		of the event
	User		the user whom have added this event
	Photo 		the acctual picture
	Preview 	a shrunken down photo
	Description	of the event
	User		the user whom have added this stair

*/
type Event_table struct {
	Event_ID    uint64  `json:"Event_ID"`
	Date 		string 	`json:"Date"`
	Address     string  `json:"Address"`
	Zipcode     string  `json:"Zipcode"`
	Name        string  `json:"Event_name"`
	Photo       string  `json:"Photo"`
	Preview 	string  `json:"Preview"`
	Description string  `json:"Info"`
	User        uint64  `json:"User"`

	//Attending	uint64  `json:"Attending"`
}

//Picture struct
/*
A Picture is composed of
	PhotoId 	a unique id for that photo
	EventId 	the id for the event
	UserId  	the id for the user
	Picture 	the acctual picture
	Preview 	a shrunken down picture

type Picture struct {
	PhotoId uint64 `json:"Photo_ID"`
	EventId uint64 `json:"Event_ID"`
	Picture string `json:"Photo"`
	Preview string `json:"Preview"`
	//UserId  uint64 `json:"userID"`
}
*/

// a custom type that we can use for handling errors and formatting responses
type handler func(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError)

// attach the standard ServeHTTP method to our handler so the http library can call it
func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// here we could do some prep work before calling the handler if we wanted to

	// call the actual handler
	response, err := fn(w, r)

	// check for errors
	if err != nil {
		log.Printf("ERROR: %v\n", err.Error)
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Message), err.Code)
		return
	}

	if response == nil {
		log.Printf("ERROR: response from method is nil\n")
		http.Error(w, "Internal server error. Check the logs.", http.StatusInternalServerError)
		return
	}

	// turn the response into JSON
	bytes, e := json.Marshal(response)
	if e != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	// send the response and log
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(bytes)
	log.Printf("%s %s %s %d", r.RemoteAddr, r.Method, r.URL, 200)
}

/*
	Add event to DB
	Function to add a event to the db
*/
func addEvent(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	data, e := ioutil.ReadAll(req.Body)

	if e != nil {

		return nil, &handlerError{e, "Can't read request", http.StatusBadRequest}
	}
	var payload Event_table
	e = json.Unmarshal(data, &payload)

	if e != nil {

		return Event_table{}, &handlerError{e, "Could'nt parse JSON", http.StatusInternalServerError}
	}
	//handle photos
	//l, _ := base64.StdEncoding.DecodeString(payload.Photo)
	//reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(payload.Photo))
	a := strings.Split(payload.Photo, ",")
	reader, err := base64.StdEncoding.DecodeString(a[1])
	if err != nil {
		return err, &handlerError{err, "Internal", http.StatusInternalServerError}
	}
	s := string(reader[:])
	photo, _, err := image.Decode(strings.NewReader(s))
	if err != nil {

		return Event_table{}, &handlerError{e, "Could'nt fix this image", http.StatusInternalServerError}
	}

	// resize photo
	newphoto := resize.Resize(215, 0, photo, resize.Lanczos3)

	//creates a buffer to save the encoded file to
	buf := new(bytes.Buffer)

	//encodes the image again and saves it to buf
	err = jpeg.Encode(buf, newphoto, nil)
	if err != nil {
		return Event_table{}, &handlerError{err, "Could'nt fix this image", http.StatusInternalServerError}
	}

	//encodes the photo to base64 agian
	Preview := base64.StdEncoding.EncodeToString(buf.Bytes())

	// adds the header from the website again
	payload.Photo = a[0] + "," + payload.Photo
	db, err := sql.Open("mymysql", "tcp:130.240.170.56:3306*mydb/dbadmin/krnhw4twf")
		//"dbadmin:krnhw4twf@tcp(130.240.170.56:3306)/mydb")
	if err != nil {

		return nil, &handlerError{err, "Internal server error", http.StatusInternalServerError}
	}
	defer db.Close()

	//inputs the event to the db
	_, err = db.Exec("insert into Event_table(Event_ID, Date, Address, Zipcode, Event_name, Info, Photo, Preview, User) values(?,?,?,?,?,?,?,?,?)", payload.Event_ID, payload.Date, payload.Address, payload.Zipcode, payload.Name, payload.Description, payload.Photo, Preview, payload.User)

	if err != nil {

		return nil, &handlerError{err, "Error adding to DB", http.StatusInternalServerError}
	}

	return payload, nil
}

/*
	Get specific event from DB
	Grabs a event from the db.
*/

func getEvent(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(req)["id"]
	db, err := sql.Open("mymysql", "dbadmin:krnhw4twf@tcp(130.240.170.56:3306)/mydb")
	if err != nil {
		return nil, &handlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer db.Close()

	row, err := db.Query("select Event_ID, Date, Address, Zipcode, Event_name, Info, User from Event_table where Event_ID =?", param)
	if err == sql.ErrNoRows {
		return nil, &handlerError{err, "Error event not found", http.StatusBadRequest}

	}

	if err != nil {
		return nil, &handlerError{err, "Internal Error when req DB", http.StatusInternalServerError}
		//panic(err)
	}

	var result []Event_table // create an array of events
	var Address, Name, Date, Zipcode, Description string
	var ID uint64
	
	for row.Next() {
		event := new(Event_table)
		if err := row.Scan(&ID, &Address, &Name, &Zipcode, &Date, &Description); err != nil {
			return nil, &handlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
			//log.Fatal(err)
		}

		event.Event_ID = ID
		event.Name = Name
		event.Address = Address
		event.Zipcode = Zipcode
		event.Date = Date
		event.Description = Description

	}

	return result, nil
}

/*
	Get all events from DB
	Returns all the events in the db
*/
func getAllEvent(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	db, err := sql.Open("mymysql", "dbadmin:krnhw4twf@tcp(130.240.170.56:3306)/mydb")
	if err != nil {
		return nil, &handlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select Event_ID, Date, Address, Zipcode, Event_name, Info, User from Event_table")
	if err != nil {
		return nil, &handlerError{err, "Error in DB", http.StatusInternalServerError}
		//log.Printf("No user with that ID")
	}

	var result []Event_table // create an array of Event
	var Address, Name, Date, Zipcode, Description string
	var ID uint64

	for rows.Next() {
		event := new(Event_table)
		err := rows.Scan(&ID, &Address, &Name, &Zipcode, &Date, &Description); //err != nil 
		if err != nil {
			return result, &handlerError{err, "Error in DB", http.StatusInternalServerError}
		}

		event.Event_ID = ID
		event.Name = Name
		event.Address = Address
		event.Zipcode = Zipcode
		event.Date = Date
		event.Description = Description

		result = append(result, *event)
	}

 	return result, nil
}

/*
	Get a user events from DB
	Grabs an event from the db.
*/

func getUserEvent(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(req)["id"]
	db, err := sql.Open("mymysql", "dbadmin:krnhw4twf@tcp(130.240.170.56:3306)/mydb")
	if err != nil {
		return nil, &handlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer db.Close()

	row, err := db.Query("select Date, Address, Zipcode, Event_name, Info, User from Event_table where User =?", param)
	if err == sql.ErrNoRows {
		return nil, &handlerError{err, "Error no event found", http.StatusBadRequest}
		//log.Printf("No user with that ID")
	}

	if err != nil {
		return nil, &handlerError{err, "Internal Error when req DB", http.StatusInternalServerError}
		//panic(err)
	}
	var result []Event_table // create an array of events
	var Address, Name, Date, Zipcode, Description string
	var ID uint64
	for row.Next() {

		event := new(Event_table)
		if err := row.Scan(&ID, &Address, &Name, &Zipcode, &Date, &Description); err != nil {
			return nil, &handlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
			//log.Fatal(err)
		}

		event.Event_ID = ID
		event.Name = Name
		event.Address = Address
		event.Zipcode = Zipcode
		event.Date = Date
		event.Description = Description

		result = append(result, *event)

	}

	return result, nil

}

/*
	Remove Event from DB
	Function not yet implemented
*/

func removeEvent(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(r)["id"]
	id, e := strconv.Atoi(param)
	if e != nil {
		return nil, &handlerError{e, "Id should be an integer", http.StatusBadRequest}
	}
	id = id

	returnable := string("removeEvent")
	return returnable, nil
}

/*
	ADD USER TO DB
	Function to add a new user to the db.
	This also checks to se that this user isnt already in the db
*/
func addUser(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {

	data, e := ioutil.ReadAll(r.Body)

	if e != nil {
		return nil, &handlerError{e, "Can't read request", http.StatusBadRequest}
	}

	// create new user called payload
	var payload User_table
	e = json.Unmarshal(data, &payload)

	if e != nil {
		return User_table{}, &handlerError{e, "Could'nt parse JSON", http.StatusInternalServerError}
	}
	db, err := sql.Open("mymysql", "dbadmin:krnhw4twf@tcp(130.240.170.56:3306)/mydb")
	if err != nil {
		return nil, &handlerError{err, "Internal server error", http.StatusInternalServerError}
	}
	defer db.Close()
	row, _ := db.Query("select count(*) from User_table where idToken=?", payload.IdToken)
	var count int
	for row.Next() {
		row.Scan(&count)
	}

	if count == 1 {
		return nil, &handlerError{nil, "User already exists", http.StatusFound}

	}

	_, err = db.Exec("insert into User_table(Username, IdToken, UserID) values(?,?,?)", payload.Name, payload.IdToken, payload.UserID)

	if err != nil {
		return nil, &handlerError{err, "Error adding to DB", http.StatusInternalServerError}
	}

	return payload, nil

}

/*
	List all users in the db
	This function lists all users from the db
*/
func listAllUsers(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	db, err := sql.Open("mymysql", "dbadmin:krnhw4twf@tcp(130.240.170.56:3306)/mydb")
	if err != nil {
		return nil, &handlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select Username, IdToken from User_table")
	if err != nil {
		return nil, &handlerError{err, "Error in DB", http.StatusInternalServerError}
		//log.Printf("No user with that ID")
	}

	var result []User_table // create an array of users
	var uid uint64
	var name string

	for rows.Next() {
		user := new(User_table)
		err = rows.Scan(&name, &uid)
		if err != nil {
			return result, &handlerError{err, "Error in DB", http.StatusInternalServerError}
		}
		user.Name = name
		user.UserID = uid
		result = append(result, *user)
	}

	return result, nil
}

/*
	Get a user from the db
	This function gets a specific user from the db.
	we sort out whom they whant by searching the incomming data for id
*/
func getUser(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	//mux.Vars(r)["id"] grabs variables from the path
	param := mux.Vars(r)["id"]
	db, err := sql.Open("mysql", "dbadmin:krnhw4twf@tcp(130.240.170.56:3306)/mydb")
	
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	row, err := db.Query("select * from User_table where IdToken =?", param)
	if err == sql.ErrNoRows {
		log.Printf("No user with that ID")
	}

	if err != nil {
		panic(err)
	}

	user := new(User_table)
	for row.Next() {
		var idToken string
		var uid uint64
		var name string

		if err := row.Scan(&uid, &name, &idToken); err != nil {
			log.Fatal(err)
		}
		user.IdToken = idToken
		user.UserID = uid
		user.Name = name
	
	}

	return user, nil
}

/*
	Remove user from DB
	Function not yet implemented
*/

func removeUser(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(r)["id"]
	id, e := strconv.Atoi(param)
	if e != nil {
		return nil, &handlerError{e, "Id should be an integer", http.StatusBadRequest}
	}
	id = id

	returnable := string("removeUser")
	return returnable, nil
}

/*
	Retrive a event photo
*/
func retriveEventPhoto(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(req)["id"]
	db, err := sql.Open("mysql", "dbadmin:krnhw4twf@tcp(130.240.170.56:3306)/mydb")
	if err != nil {
		return nil, &handlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer db.Close()

	row, err := db.Query("select Photo from User_table where Event_ID =?", param)
	if err == sql.ErrNoRows {
		return nil, &handlerError{err, "Error event not found", http.StatusBadRequest}
		//log.Printf("No user with that ID")
	}

	if err != nil {
		return nil, &handlerError{err, "Internal Error when req DB", http.StatusInternalServerError}
		//panic(err)
	}

	var result []Event_table
	var Photo string

	for row.Next() {
		picture := new(Event_table)

		if err := row.Scan(&Photo); err != nil {
			return nil, &handlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
			//log.Fatal(err)
		}

		picture.Photo = Photo
		result = append(result, *picture)

	}

	return result, nil
}

/*
	Retrive a events preview pictures from the db
*/
func retriveEventPreview(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(req)["id"]
	db, err := sql.Open("mysql", "dbadmin:krnhw4twf@tcp(130.240.170.56:3306)/mydb")
	if err != nil {
		return nil, &handlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer db.Close()

	row, err := db.Query("select Preview from Event_table where Event_ID =?", param)
	if err == sql.ErrNoRows {
		return nil, &handlerError{err, "Error commenting on Event", http.StatusBadRequest}

	}

	if err != nil {
		return nil, &handlerError{err, "Internal Error when req DB", http.StatusInternalServerError}
	}
	var result []Event_table
	var Preview string

	for row.Next() {
		picture := new(Event_table)

		if err := row.Scan(&Preview); err != nil {
			return nil, &handlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
		}

		picture.Preview = Preview
		result = append(result, *picture)

	}

	return result, nil
}

func main() {
	// command line flags
	port := flag.Int("port", 8000, "port to serve on")
	dir := flag.String("directory", "web/", "directory of web files")
	flag.Parse()

	// handle all requests by serving a file of the same name
	fs := http.Dir(*dir)
	fileHandler := http.FileServer(fs)

	// setup routes
	router := mux.NewRouter()
	router.Handle("/", http.RedirectHandler("/static/", 302))

	// Handlers for Users
	router.Handle("/users", handler(listAllUsers)).Methods("GET")
	router.Handle("/users", handler(addUser)).Methods("POST")
	router.Handle("/users/{id}", handler(getUser)).Methods("GET")
	router.Handle("/users/{id}", handler(removeUser)).Methods("DELETE")

	// Handlers for event
	router.Handle("/event", handler(addEvent)).Methods("POST")
	router.Handle("/event/{id}", handler(getEvent)).Methods("GET")
	router.Handle("/event", handler(getAllEvent)).Methods("GET")
	router.Handle("/event/{id}", handler(removeEvent)).Methods("DELETE")

	// Get all event a user have added
	router.Handle("/users/event/{id}", handler(getUserEvent)).Methods("GET")
	//Get all pictures for a event
	router.Handle("/event/picture/{id}", handler(retriveEventPhoto)).Methods("GET")
	//Get all preview pictures for a event
	router.Handle("/event/picture/preview/{id}", handler(retriveEventPreview)).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	http.Handle("/", router)

	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf("130.240.170.56:%d", *port)
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
