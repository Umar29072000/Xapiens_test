package handler

import (
	"encoding/json"
	"net/http"
	"Xapiens_test/models"
	"Xapiens_test/services"
)

var users = []models.User{}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "POST":
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Invalid request payload")
			return
		}
		if user.Username == "" || user.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Username and password are required")
			return
		}
		for _, u := range users {
			if u.Username == user.Username {
				w.WriteHeader(http.StatusConflict)
				json.NewEncoder(w).Encode("User already exists")
				return
			}
		}
		users = append(users, user)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)

	case "GET":
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)

	case "PUT":
		var updatedUser models.User
		if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Invalid request payload")
			return
		}
		for i, user := range users {
			if user.ID == updatedUser.ID {
				users[i].Username = updatedUser.Username
				users[i].Password = updatedUser.Password
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(users[i])
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")

	case "DELETE":
		var userToDelete models.User
		if err := json.NewDecoder(r.Body).Decode(&userToDelete); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Invalid request payload")
			return
		}
		for i, user := range users {
			if user.ID == userToDelete.ID {
				users = append(users[:i], users[i+1:]...)
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode("User deleted")
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request payload")
		return
	}
	if user.Username == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Username and password are required")
		return
	}
	for _, u := range users {
		if u.Username == user.Username && u.Password == user.Password {
			token, err := services.GenerateJWT(u.Username)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode("Error generating token")
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(token))
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode("Invalid credentials")
}
