package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Database struct {
	DB *sql.DB
}

type User struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

var (
	db Database
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	//(w).Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	log.Println("Got a request")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	var databaseUsername string
	var databasePassword string

	result := db.DB.QueryRow("SELECT username, password FROM user WHERE username=?", user.Username)
	err = result.Scan(&databaseUsername, &databasePassword)
	if databasePassword == user.Password {
		fmt.Fprintf(w, "login successful")
		fmt.Println("Login succesful")
	} else {
		fmt.Println("Login unsuccessful")

	}
	// err = db.DB.QueryRow("SELECT username, password FROM user WHERE username=?", user.Username).Scan(&databaseUsername, &databasePassword)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(user.Password))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// w.Write([]byte("Hello" + databaseUsername))
	// fmt.Printf("Login")
}

func (database *Database) CloseDB() {
	database.DB.Close()
}

func (database *Database) InitializeDB() {
	var err error
	//DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	database.DB, err = sql.Open("mysql", "comp:pass@tcp(db:3306)/comp440")
	if err != nil {
		log.Fatal(err)
	}

	// database.DB, err = gorm.Open(Dbdriver, DBURL)
	//fmt.Println(DBURL)
	//database.DB, err = sql.Open("mysql", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to database")
		log.Fatal("This is the error: ", err)
	} else {
		fmt.Println("Server: connected to the database")
	}
}

func main() {
	db.InitializeDB()
	fmt.Println("Started Go Server")
	defer db.CloseDB()
	router := mux.NewRouter()
	router.HandleFunc("/register", registerUser)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
