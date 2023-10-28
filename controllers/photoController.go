package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/username/projectname/app"
)

func GetPhoto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	photoID := params["id"]

	var photo app.Photo
	query := "SELECT id, title, caption, photo_url, user_id FROM photos WHERE id = ?"
	err := db.QueryRow(query, photoID).Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoURL, &photo.UserID)
	if err != nil {
		http.Error(w, "Photo not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(photo)
}

func DeletePhoto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	photoID := params["id"]

	deleteQuery := "DELETE FROM photos WHERE id = ?"
	_, err := db.Exec(deleteQuery, photoID)
	if err != nil {
		http.Error(w, "Failed to delete photo", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Photo deleted successfully"))
}

func UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	photoID := params["id"]

	var updatedPhoto app.Photo
	_ = json.NewDecoder(r.Body).Decode(&updatedPhoto)

	updateQuery := "UPDATE photos SET title=?, caption=?, photo_url=?, user_id=? WHERE id=?"
	_, err := db.Exec(updateQuery, updatedPhoto.Title, updatedPhoto.Caption, updatedPhoto.PhotoURL, updatedPhoto.UserID, photoID)
	if err != nil {
		http.Error(w, "Failed to update photo", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Photo updated successfully"))
}
