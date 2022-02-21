package main

import (
	"bookstore_case/middleware"
	routes "bookstore_case/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/user/login", routes.UserLogin).Methods("POST")
	router.HandleFunc("/api/user/registration", routes.SignUp).Methods("POST")

	router.HandleFunc("/api/stock/add", middleware.IsAuthorized(routes.AddStock)).Methods("POST")
	router.HandleFunc("/api/stock/update", middleware.IsAuthorized(routes.UpdateStock)).Methods("POST")
	router.HandleFunc("/api/stock/delete", middleware.IsAuthorized(routes.DeleteStock)).Methods("DELETE")
	router.HandleFunc("/api/order/buy", middleware.IsAuthorized(routes.OrderBuy)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
