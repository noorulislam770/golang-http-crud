package tests

import (
	"bytes"
	"golang-http-crud/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	payload := []byte(`{"name":"Alice","email":"alice@example.com","dob":"1995-01-01","gender":"female","hobbies":["reading","swimming"]}`)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()

	handlers.UserHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code 201, got %v", w.Code)
	}
}
