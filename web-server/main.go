package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	Id    string `json::"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	users = make(map[string]User)
	mu    sync.Mutex
)

// create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	users[user.Id] = user
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// get user
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/"):]

	mu.Lock()
	user, ok := users[id]
	mu.Unlock()

	if !ok {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// UpdateUser
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/"):]
	var userUpdate User
	if err := json.NewDecoder(r.Body).Decode(&userUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	user, ok := users[id]
	if !ok {
		mu.Unlock()
		http.NotFound(w, r)
		return
	}

	user.Name = userUpdate.Name
	user.Email = userUpdate.Email
	users[id] = user
	mu.Unlock()

	json.NewEncoder(w).Encode(user)
}

// DeleteUser
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/"):]

	mu.Lock()
	_, ok := users[id]
	if !ok {
		mu.Unlock()
		http.NotFound(w, r)
		return
	}

	delete(users, id)
	mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/users/", GetUser)           // GET
	http.HandleFunc("/users/new", CreateUser)     // POST
	http.HandleFunc("/users/update/", UpdateUser) // PUT
	http.HandleFunc("/users/delete/", DeleteUser) // DELETE

	fmt.Println("Server is running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
