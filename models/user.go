package models

type User struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	DOB     string   `json:"dob"`
	Gender  string   `json:"gender"`
	Hobbies []string `json:"hobbies"`
}
