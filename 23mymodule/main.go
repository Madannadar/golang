package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("video on module")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":400", r)) // log error
}

func greeter() {
	fmt.Println("Hey there nod users")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang serier on youtube</h1>"))
}
