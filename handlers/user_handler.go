package handlers

import (
	"encoding/json"
	"golang-http-crud/models"
	"golang-http-crud/repository"
	"net/http"
	"strconv"
	"strings"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createUser(w, r)
	case http.MethodPut:
		updateUser(w, r)
	case http.MethodPatch:
		patchUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if user.Email == "" || user.Name == "" {
		http.Error(w, "Name and Email are required", http.StatusBadRequest)
		return
	}

	users, err := repository.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	if repository.FindUserByEmail(users, user.Email) {
		http.Error(w, "User with this email already exists", http.StatusConflict)
		return
	}

	user.ID = strconv.Itoa(len(users) + 1)
	users = append(users, user)
	if err := repository.SaveUsers(users); err != nil {
		http.Error(w, "Failed to save user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	// Full update logic here
}

func patchUser(w http.ResponseWriter, r *http.Request) {
	// Partial update logic here
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// Delete logic here
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	searchType := r.URL.Query().Get("type")

	users, err := repository.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	var results []models.User
	for _, user := range users {
		switch strings.ToLower(searchType) {
		case "name":
			if strings.Contains(strings.ToLower(user.Name), strings.ToLower(query)) {
				results = append(results, user)
			}
		case "email":
			if strings.Contains(strings.ToLower(user.Email), strings.ToLower(query)) {
				results = append(results, user)
			}
		case "hobbies":
			for _, hobby := range user.Hobbies {
				if strings.Contains(strings.ToLower(hobby), strings.ToLower(query)) {
					results = append(results, user)
					break
				}
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
