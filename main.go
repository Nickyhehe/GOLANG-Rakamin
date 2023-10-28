package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Photo struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photoUrl"`
	UserID   string `json:"userId"`
}

var users []User
var photos []Photo

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	users = append(users, newUser)

	json.NewEncoder(w).Encode(newUser)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUser User
	_ = json.NewDecoder(r.Body).Decode(&loginUser)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

}

func CreatePhoto(w http.ResponseWriter, r *http.Request) {
	var newPhoto Photo
	_ = json.NewDecoder(r.Body).Decode(&newPhoto)

	photos = append(photos, newPhoto)

	json.NewEncoder(w).Encode(newPhoto)
}

func GetPhotos(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(photos)
}

func UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	photoID := vars["photoId"]

}

func DeletePhoto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	photoID := vars["photoId"]

}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/users/register", RegisterUser).Methods("POST")
	router.HandleFunc("/users/login", LoginUser).Methods("POST")
	router.HandleFunc("/users/{userId}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{userId}", DeleteUser).Methods("DELETE")

	router.HandleFunc("/photos", CreatePhoto).Methods("POST")
	router.HandleFunc("/photos", GetPhotos).Methods("GET")
	router.HandleFunc("/photos/{photoId}", UpdatePhoto).Methods("PUT")
	router.HandleFunc("/photos/{photoId}", DeletePhoto).Methods("DELETE")

	fmt.Println("Server started on :8000")
	http.ListenAndServe(":8000", router)
}
