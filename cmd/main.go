package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const amountOfTimeDivisions = 4

func main() {

	router := mux.NewRouter()
	fmt.Println("Server started")

	router.HandleFunc("/service/u/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", Redirect)
	router.HandleFunc("/service/{id:[a-z A-Z][0-9]+}", Redirect)

	http.Handle("/", router)

	go GarbageCollector()

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
