package main

import (
	"github.com/drone/routes"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type User struct{
	Email string `json:"email"`
	Zip string `json:"zip"`
	Country string `json:"country"`
	Profession string `json:"profession"`
	Favorite_Color string `json:"favorite_color"`
	Is_Smoking string `json:"is_smoking"`
	Favorite_Sport string `json:"favorite_sport"`
	Food `json:"food"`
	Music `json:"music"`
	Movie `json:"movie"`
	Travel `json:"travel"`
}

type Food struct {
	Type string `json:"type"`
	Drink_Alcohol string `json:"drink_alcohol"`
}
type Music struct{
	Spotify_User_Id string `json:"spotify_user_id"`
}
type Movie struct{
	TV_Shows []string `json:"tv_shows"`
	Movies []string `json:"movies"`
}
type Travel struct {
	Flight `json:"flight"`
}
type Flight struct{
	Seat string `json:"seat"`
}


var users map[string]User

func main() {
	mux := routes.New()
	mux.Post("/profile",CreateProfile)
	mux.Get("/profile/:email", GetProfile)
	mux.Del("/profile/:email", DeleteProfile)
	mux.Put("/profile/:email",UpdateProfile)

	http.Handle("/", mux)
	users = make(map[string]User)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func CreateProfile(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Something is wrong ***")
	}
	var newUser User
	err = json.Unmarshal([]byte(body), &newUser)
	if err != nil {
		log.Println("Something is wrong ***")
	}
	users[newUser.Email] = newUser
	w.WriteHeader(http.StatusCreated)
	
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	email := params.Get(":email")
	user, ok := users[email]
	if ok == true {
		outgoingJSON, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(outgoingJSON))
	}else {
		w.Write([]byte("No record exists with email id"+ email))
	}
}

func DeleteProfile(w http.ResponseWriter, r *http.Request){
	params := r.URL.Query()
	email := params.Get(":email")
	_, ok := users[email]
	if ok == true {
		delete(users, email)
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Deleted User:" + email))

	}else {
		w.Write([]byte("No record exists with email id"+ email))
	}

}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	email := params.Get(":email")
	user, ok := users[email]
	if ok == true {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Something is wrong ***")
		}
		
		err = json.Unmarshal([]byte(body), &user)
		if err != nil {
			log.Println("Something is wrong ***")
		}
		users[user.Email] = user
		w.WriteHeader(http.StatusNoContent)
	}else {
		fmt.Fprintf(w, "No record exists with email id %s", email)
	}
}


