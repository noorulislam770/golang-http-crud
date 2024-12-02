package repository

import (
	"encoding/json"
	"golang-http-crud/models"
	"os"
)

const FilePath = "users.json"

func GetAllUsers() ([]models.User, error) {
	file, err := os.OpenFile(FilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []models.User
	if stat, _ := file.Stat(); stat.Size() > 0 {
		err = json.NewDecoder(file).Decode(&users)
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}

func SaveUsers(users []models.User) error {
	file, err := os.OpenFile(FilePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(users)
}

func FindUserByID(users []models.User, id string) (*models.User, int) {
	for i, user := range users {
		if user.ID == id {
			return &user, i
		}
	}
	return nil, -1
}

func FindUserByEmail(users []models.User, email string) bool {
	for _, user := range users {
		if user.Email == email {
			return true
		}
	}
	return false
}

func FindUserByName(users []models.User, name string) bool {
	for _, user := range users {
		if user.Name == name {
			return true
		}
	}
	return false
}
