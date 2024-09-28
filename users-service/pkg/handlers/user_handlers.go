package handlers

import (
    "net/http"
)

// GetUser handles GET requests for retrieving a user
func GetUser(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Here is the user"))
}

// Signup handles POST requests for creating a new user
func Signup(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("A new user has been created"))
}

// Login handles POST requests for logging in a user
func Login(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("A user has been logged in"))
}

// UpdateUser handles PUT requests for updating a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("A user has been updated"))
}

// DeleteUser handles DELETE requests for deleting a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("A user has been deleted"))
}
