package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"personalities_api/database"
	"personalities_api/models"

	"github.com/gorilla/mux"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the home page!")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {

	var personality []models.Personality

	db := database.GetDbConnection()
	db.Find(&personality)

	json.NewEncoder(w).Encode(personality)
}

func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var personality models.Personality

	db := database.GetDbConnection()
	db.First(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality

	json.NewDecoder(r.Body).Decode(&personality)

	db := database.GetDbConnection()
	db.Create(&personality)

	json.NewEncoder(w).Encode(&personality)
}
func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var personality models.Personality

	db := database.GetDbConnection()
	db.Delete(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var personality models.Personality

	db := database.GetDbConnection()
	db.First(&personality, id)

	json.NewDecoder(r.Body).Decode(&personality)

	db.Save(&personality)

	json.NewEncoder(w).Encode(personality)
}
