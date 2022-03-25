package api

import (
	"errors"
	"technical-test-atmatech/database"
	"technical-test-atmatech/models"
)

func GetUserById(id uint) (models.User, error) {
	var user models.User
	var err error

	database.DB.Find(&user, "id = ?", id)
	if user.Username == "" {
		err = errors.New("user does not exist")
	}
	return user, err
}
