package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var people []Person
var users []user

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"City,omitempty"`
	State string `json:"state,omitempty"`
}

type user struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	json.NewEncoder(w).Encode(people)
}

func GetUsersEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	db, err := sql.Open("mysql", "root:krasniqi01@tcp(localhost:3306)/test")
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	model := user{}
	for rows.Next() {
		if err := rows.Scan(&model.ID, &model.Name, &model.Email); err != nil {
			log.Fatal(err)
		}
		users = append(users, model)
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Ardit", Lastname: "Krasniqi", Address: &Address{City: "Damanek", State: "Kosova"}})
	router.HandleFunc("/people", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/users", GetUsersEndPoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}
