package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	controller "golang-crud/controller"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getEmployee", controller.AllEmployee).Methods("GET")
	router.HandleFunc("/insertEmployee", controller.InsertEmployee).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Connected to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
