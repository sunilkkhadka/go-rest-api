package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var profiles []Profile = []Profile{}

type User struct{
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
}

type Profile struct {
	Department string `json:"department"`
	Designation string `json:"designation"`
	Employee User `json:"employee"`
}

func addItem(q http.ResponseWriter, r *http.Request){
	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)

	q.Header().Set("Content-Type", "application/json")

	profiles = append(profiles, newProfile)

	json.NewEncoder(q).Encode(profiles)
}


func main(){
	router := mux.NewRouter()

	router.HandleFunc("/profiles", addItem).Methods("POST")

	fmt.Println("Server started...")
	http.ListenAndServe(":5000", router)
}

