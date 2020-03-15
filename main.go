package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Database struct {
	DB *sql.DB
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	db Database
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got some request")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

	}
	defer r.Body.Close()
	var user User
	err = json.Unmarshal(body, &user)
	// saveUserToDB(user)
	fmt.Println(user.Username, user.Password)
}

func (database *Database) InitializeDB(DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	// database.DB, err = gorm.Open(Dbdriver, DBURL)
	fmt.Println(DBURL)
	database.DB, err = sql.Open("mysql", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to database")
		log.Fatal("This is the error: ", err)
	} else {
		fmt.Printf("We are connected to the database")
	}
}

func main() {
	//db.InitializeDB()
	//db.InitializeDB(DbUser, DbPassword, DbPort, DbHost, DbName)
	fmt.Println("Started Go Server")
	//defer db.CloseDB()

	router := mux.NewRouter()
	router.HandleFunc("/register", registerUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
