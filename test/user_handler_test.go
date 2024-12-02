package tests

import (
	"bytes"
	"golang-http-crud/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCreateUser tests the creation of a new user.
func TestCreateUser(t *testing.T) {
	// Create a new user payload
	payload := []byte(`{"name":"Alice testing 3","email":"alicesecond22@example.com","dob":"1995-01-01","gender":"female","hobbies":["reading","swimming"]}`)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()

	// Call the UserHandler to simulate the POST request
	handlers.UserHandler(w, req)

	// Check if the status code is 201 Created
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code 201, got %v", w.Code)
	}
}

// TestGetUsers tests fetching all users.
func TestGetUsers(t *testing.T) {
	// Create a test user to ensure there is at least one user in the system

	// Send a GET request to fetch users
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	// Call the UserHandler to simulate the GET request
	handlers.UserHandler(w, req)

	// Check if the status code is 200 OK
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", w.Code)
	}

	// Verify the response body contains user data

}

// TestUpdateUser tests updating an existing user.
func TestUpdateUser(t *testing.T) {
	// Set initial users

	// Prepare the updated user data
	payload := []byte(`{"name":"Alice Smith","email":"alice.smith2wa@example.com","dob":"1995-01-01","gender":"female","hobbies":["reading","swimming","cycling"]}`)
	req := httptest.NewRequest(http.MethodPut, "/users?id=1", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()

	// Call the UserHandler to simulate the PUT request
	handlers.UserHandler(w, req)

	// Check if the response status is 200 OK
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", w.Code)
	}

	// Verify the updated user data in the response body

}

// TestDeleteUser tests deleting a user.
func TestDeleteUser(t *testing.T) {
	// Set initial users

	// Send a DELETE request to delete the user with ID=1
	req := httptest.NewRequest(http.MethodDelete, "/users?id=1", nil)
	w := httptest.NewRecorder()

	// Call the UserHandler to simulate the DELETE request
	handlers.UserHandler(w, req)

	// Check if the status code is 200 OK (assuming successful deletion)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", w.Code)
	}

}
