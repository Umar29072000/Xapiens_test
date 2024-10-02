package main

import (
	"log"
	"net/http"
	"Xapiens_test/handler"
	"Xapiens_test/middleware"
)

func main() {
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/users", handler.UsersHandler)
	http.HandleFunc("/users-protected", middleware.JWTAuth(handler.UsersHandler))
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
