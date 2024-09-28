package handlers

import (
    "net/http"
)

// GetUsers handles GET requests for retrieving all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Here are all the users"))
}

// RegisterUser handles POST requests for registering a new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("A new user has been registered"))
}
