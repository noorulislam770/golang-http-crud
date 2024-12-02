package handlers

import (
	"encoding/json"
	"golang-http-crud/models"
	"golang-http-crud/repository"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
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

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repository.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
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
	// Get user ID from query parameters
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Retrieve all users
	users, err := repository.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	// Find the user by ID
	var existingUser *models.User
	for i := range users {
		if users[i].ID == userID {
			existingUser = &users[i]
			break
		}
	}
	if existingUser == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode partial updates
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Apply updates dynamically
	for key, value := range updates {
		switch key {
		case "name":
			if v, ok := value.(string); ok {
				existingUser.Name = v
			}
		case "email":
			if v, ok := value.(string); ok {
				existingUser.Email = v
			}
		case "hobbies":
			if v, ok := value.([]interface{}); ok {
				var hobbies []string
				for _, hobby := range v {
					if hobbyStr, ok := hobby.(string); ok {
						hobbies = append(hobbies, hobbyStr)
					}
				}
				existingUser.Hobbies = hobbies
			}
		default:
			// Ignore unknown or unsupported fields
		}
	}

	// Save updated users back to the repository
	if err := repository.SaveUsers(users); err != nil {
		http.Error(w, "Failed to save updates", http.StatusInternalServerError)
		return
	}

	// Respond with the updated user
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingUser)
}

func patchUser(w http.ResponseWriter, r *http.Request) {
	// Get user ID from URL query or request parameters
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Retrieve all users
	users, err := repository.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	// Find the user by ID
	var existingUser *models.User
	for i := range users {
		if users[i].ID == userID {
			existingUser = &users[i]
			break
		}
	}
	if existingUser == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode partial updates
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Apply updates dynamically
	for key, value := range updates {
		switch key {
		case "name":
			if v, ok := value.(string); ok {
				existingUser.Name = v
			}
		case "email":
			if v, ok := value.(string); ok {
				existingUser.Email = v
			}
		case "hobbies":
			if v, ok := value.([]interface{}); ok {
				var hobbies []string
				for _, hobby := range v {
					if hobbyStr, ok := hobby.(string); ok {
						hobbies = append(hobbies, hobbyStr)
					}
				}
				existingUser.Hobbies = hobbies
			}
		default:
			// Ignore unknown fields
		}
	}

	// Save updated users
	if err := repository.SaveUsers(users); err != nil {
		http.Error(w, "Failed to save updates", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingUser)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// Get user ID from URL query or request parameters
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Retrieve all users
	users, err := repository.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	// Find and remove the user
	var updatedUsers []models.User
	var userFound bool
	for _, user := range users {
		if user.ID == userID {
			userFound = true
			continue // Skip adding the user to the updated list
		}
		updatedUsers = append(updatedUsers, user)
	}

	if !userFound {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Save updated users
	if err := repository.SaveUsers(updatedUsers); err != nil {
		http.Error(w, "Failed to save updates", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	searchType := r.URL.Query().Get("type")
	log.Println(query, searchType)
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
		case "id":
			log.Println("entered id")
			if strings.Contains(strings.ToLower(user.ID), strings.ToLower(query)) {
				results = append(results, user)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
