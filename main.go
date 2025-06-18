package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = make(map[int]User)
var idCounter = 1

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/", userHandler)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handle /users route (GET all, POST new)
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAllUsers(w, r)
	case http.MethodPost:
		createUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handle /users/{id} route
func userHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getUser(w, r, id)
	case http.MethodPut:
		updateUser(w, r, id)
	case http.MethodDelete:
		deleteUser(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// CRUD operations

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var allUsers []User
	for _, user := range users {
		allUsers = append(allUsers, user)
	}
	json.NewEncoder(w).Encode(allUsers)
}

func getUser(w http.ResponseWriter, r *http.Request, id int) {
	user, exists := users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	user.ID = idCounter
	idCounter++
	users[user.ID] = user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request, id int) {
	_, exists := users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	user.ID = id
	users[id] = user
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	_, exists := users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	delete(users, id)
	w.WriteHeader(http.StatusNoContent)
}
