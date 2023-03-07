package main

import (
	"net/http"
	"os"

	"github.com/emilgibi/orders-microservices/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	var handlerObj handlers.Handler

	godotenv.Load()

	HOST := os.Getenv("DATABASE_HOST")
	USER := os.Getenv("DATABASE_USER")
	PASS := os.Getenv("DATABASE_PASS")
	NAME := os.Getenv("DATABASE_NAME")
	PORT := os.Getenv("DATABASE_PORT")

	handlerObj.Connect(HOST, USER, PASS, NAME, PORT)

	dbinstance, _ := handlerObj.DB.DB()
	defer dbinstance.Close()

	router := mux.NewRouter()

	router.HandleFunc("/order", handlerObj.GetOrder).Methods("GET")
	router.HandleFunc("/order", handlerObj.AddOrder).Methods("POST")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8084", router)
}
