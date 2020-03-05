package main

import "github.com/gorilla/mux"

//create structs for JSON objects recieved and responses

//touter service struct
type Server struct {
	router *mux.Router
}

//config struct
type Config struct {
	CRUDHost        string
	CRUDPort        string
	MESSAGINGPort string
}
