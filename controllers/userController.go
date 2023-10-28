package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/username/projectname/app"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "username:password@tcp(localhost:3306)/dbname")
	if err != nil {
		panic(err.Error())
	}
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser app.User
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	insertQuery := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	_, err := db.Exec(insertQuery, newUser.Username, newUser.Email, newUser.Password)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	var user app.User
	query := "SELECT id, username, email FROM users WHERE id = ?"
	err := db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUser app.User
	_ = json.NewDecoder(r.Body).Decode(&loginUser)

	var user app.User
	query := "SELECT id, username, email FROM users WHERE email = ? AND password = ?"
	err := db.QueryRow(query, loginUser.Email, loginUser.Password).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(user)
}
