package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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

func main(){
	router := mux.NewRouter()
	http.ListenAndServe(":5000", router)
	fmt.Println("Starting server...")
}