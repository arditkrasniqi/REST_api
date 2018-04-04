package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var people []Person
var users []user
var db, dberror = sql.Open("mysql", "root:krasniqi01@tcp(localhost:3306)/test")

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
	json.NewEncoder(w).Encode(people)
}

func GetUsersEndpoint(w http.ResponseWriter, req *http.Request) {
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
	defer rows.Close()
	json.NewEncoder(w).Encode(users)
}

func InsertUserEndpoint(w http.ResponseWriter, req *http.Request) {
	_, err := db.Query("insert into users(name, email) values('" + req.FormValue("name") + "','" + req.FormValue("email") + "')")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, "Inserted")
}

func DeleteUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	_, err := db.Query("delete from users where id = " + params["id"])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, "Deleted")
}

func UpdateUserEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	_, err := db.Query("update users set name = '" + req.FormValue("name") + "', email = '" + req.FormValue("email") + "' where id = " + params["id"])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, "Updated")
}

func main() {
	dberror = db.Ping()
	if dberror != nil {
		log.Fatal(dberror)
	}
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Ardit", Lastname: "Krasniqi", Address: &Address{City: "Damanek", State: "Kosova"}})
	router.HandleFunc("/people", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/users", GetUsersEndpoint).Methods("GET")
	router.HandleFunc("/new-user", InsertUserEndpoint).Methods("POST")
	router.HandleFunc("/delete-user/{id}", DeleteUserEndpoint).Methods("DELETE")
	router.HandleFunc("/update-user/{id}", UpdateUserEndpoint).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":12345", router))
}
