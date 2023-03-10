package controller

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

func (server *Server) Initialize() {
	server.DB = Connect()
	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func Connect() *sql.DB {
	var err error
	var db *sql.DB

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbMaxOpenConns, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	db.SetMaxOpenConns(dbMaxOpenConns)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected to Db "+dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	}
	return db
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port " + addr)
	handler := cors.AllowAll().Handler(server.Router)
	log.Fatal(http.ListenAndServe(addr, handler))
}
