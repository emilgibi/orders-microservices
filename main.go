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

	HOST := os.Getenv("DB_HOST")
	USER := os.Getenv("USER_NAME")
	PASS := os.Getenv("PASS")

	handlerObj.Connect(HOST, USER, PASS, "postgres", "5432")

	dbinstance, _ := handlerObj.DB.DB()
	defer dbinstance.Close()

	router := mux.NewRouter()

	router.HandleFunc("/order", handlerObj.GetOrder).Methods("GET")
	router.HandleFunc("/order", handlerObj.AddOrder).Methods("POST")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8084", router)
}
