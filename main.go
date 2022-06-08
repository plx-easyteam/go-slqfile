package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	r := mux.NewRouter()

	if e != nil {
		// Be sure to remove any comments in the .env file
		// If needed use command "set .env"
		log.Fatalln("::::: Error loading .env file :::::")
	}

	localhost := os.Getenv("DB_Host")
	port := os.Getenv("PORT")

	r.HandleFunc("/", greet).Methods("GET")

	log.Printf("..... Server running on %v:%v .....", localhost, port)
	log.Fatalln(http.ListenAndServe("localhost:"+port, r))
}


func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Greetings World!")
}
