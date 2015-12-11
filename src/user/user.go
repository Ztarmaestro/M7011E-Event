package api

import (
	"database/sql"
	"encoding/json"
	//	"flag"
	"fmt"
	_ "github.com/ziutek/mymysql/godrv"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	//	"time"

	"github.com/gorilla/mux"
)

type User_tabel struct {
	UserID    uint64 `json:"userID"`
	IdToken   string `json:"id"`
	Username  string `json:"Username"`
}

/*
	List all users in the db
	!READY FOR TESTING!
*/
func ListAllUsers(w http.ResponseWriter, r *http.Request) (interface{}, *HandlerError) {
	con, err := sql.Open("mymysql", "tcp:130.240.170.56:3306*mydb/dbadmin/eventdb")
	if err != nil {
		return nil, &HandlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	rows, err := con.Query("select Username, userID from Users")
	if err != nil {
		return nil, &HandlerError{err, "Error in DB", http.StatusInternalServerError}
		//log.Printf("No user with that ID")
	}

	var result []User // create an array of stairs
	var userID uint64
	var Username string

	for rows.Next() {
		user := new(User)
		err = rows.Scan(&name, &uid)
		if err != nil {
			return result, &HandlerError{err, "Error in DB", http.StatusInternalServerError}
		}
		user.Username = name
		user.userID = uid
		result = append(result, *user)
	}

	return result, nil
}

/*
	Get a user from the db
	!DONE FOR TESTING!
*/
func GetUser(w http.ResponseWriter, r *http.Request) (interface{}, *HandlerError) {
	//mux.Vars(r)["id"] grabs variables from the path
	param := mux.Vars(r)["id"]
	fmt.Println(param)
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select * from Users where idToken =?", param)
	if err == sql.ErrNoRows {
		log.Printf("No user with that ID")
	}

	if err != nil {
		panic(err)
	}

	user := new(User)
	for row.Next() {
		var idToken string
		var uid uint64
		var name, lastname, photo string

		fmt.Println(row)

		if err := row.Scan(&uid, &name, &lastname, &idToken, &photo); err != nil {
			log.Fatal(err)
		}
		user.IdToken = idToken
		user.UserID = uid
		user.FirstName = name
		user.LastName = lastname
		user.Photo = photo
	}
	//fmt.Println(row)
	//	user.UserID = row[0]
	//	user.FirstName = row[1]
	//	user.LastName = row[2]

	//returnable := json.Marshal(user)
	//returnable := json.Marshal(user)

	return user, nil
}

/*
	ADD USER TO DB
	!DONE for TESTING!
*/

func AddUser(w http.ResponseWriter, r *http.Request) (interface{}, *HandlerError) {

	data, e := ioutil.ReadAll(r.Body)

	fmt.Println("BEFORE UNMARSHAL" + string(data))
	if e != nil {
		fmt.Println("AJAJAJ 1111")
		fmt.Println(string(data))
		return nil, &HandlerError{e, "Can't read request", http.StatusBadRequest}
	}

	// create new user called payload
	var payload User
	e = json.Unmarshal(data, &payload)

	if e != nil {
		fmt.Println("SATAN")
		fmt.Println(e)
		fmt.Println("kunde inte unmarshla detta:")
		fmt.Println(payload)
		return Stair{}, &HandlerError{e, "Could'nt parse JSON", http.StatusInternalServerError}
	}
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		fmt.Println("Kunde inte öppna DB")
		return nil, &HandlerError{err, "Internal server error", http.StatusInternalServerError}
	}
	defer con.Close()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Detta är token:")
	fmt.Println(payload.IdToken)
	fmt.Println("")
	fmt.Println("")
	row, _ := con.Query("select count(*) from Users where idToken=?", payload.IdToken)
	var count int
	for row.Next() {
		row.Scan(&count)
	}

	if count == 1 {
		return nil, &HandlerError{nil, "User already exists", http.StatusFound}

	}

	_, err = con.Exec("insert into Users( UserID, idToken, Username) values(?,?,?,?)", payload.UserID, payload.IdToken, payload.Username)

	if err != nil {
		fmt.Println("Kunde inte lägga till :/")
		return nil, &HandlerError{err, "Error adding to DB", http.StatusInternalServerError}
	}

	return payload, nil
	//row, err := con.Query("select * from users where uid =?", param)
}

/*
	Remove user from DB
*/

func RemoveUser(w http.ResponseWriter, r *http.Request) (interface{}, *HandlerError) {
	param := mux.Vars(r)["id"]
	id, e := strconv.Atoi(param)
	if e != nil {
		return nil, &HandlerError{e, "Id should be an integer", http.StatusBadRequest}
	}
	fmt.Println(id)
	// this is jsut to check to see if the book exists

	returnable := string("removeUser")
	return returnable, nil
}