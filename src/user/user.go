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
	UserID    uint64 `json:"UserID"`
	IdToken   string `json:"IdToken"`
	Username  string `json:"Username"`
}

/*
	Get a user from the db
	!DONE FOR TESTING!
*/
func GetUser(w http.ResponseWriter, r *http.Request) (interface{}, *HandlerError) {
	//mux.Vars(r)["id"] grabs variables from the path
	param := mux.Vars(r)["id"]
	fmt.Println(param)
	con, err := sql.Open("mymysql", "tcp:130.240.170.56:3306*mydb/dbadmin/eventdb")
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
		var IdToken string
		var UserID uint64
		var Username string

		fmt.Println(row)

		if err := row.Scan(&IdToken, &UserID, &Username); err != nil {
			log.Fatal(err)
		}
		//user.IdToken = id
		//user.UserID = userid
		//user.Username = name
		
	}

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
	con, err := sql.Open("mymysql", "tcp:130.240.170.56:3306*mydb/dbadmin/eventdb")
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

	_, err = con.Exec("insert into User( UserID, idToken, Username) values(?,?,?,?)", payload.UserID, payload.IdToken, payload.Username)

	if err != nil {
		fmt.Println("Kunde inte lägga till :/")
		return nil, &HandlerError{err, "Error adding to DB", http.StatusInternalServerError}
	}

	return payload, nil
	//row, err := con.Query("select * from users where uid =?", param)
}
