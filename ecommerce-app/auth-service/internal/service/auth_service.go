package service

import (
	"auth-service/internal/config"
	"auth-service/internal/model"
)

func CreateUser(user *model.User) error {
	// set default values
	if user.Status == "" {
		user.Status = "active"
	}
	if user.Role == "" {
		user.Role = "user"
	}
	user.DeleteFlag = false

	// save data to db
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByEmail find user by email
func GetUserByEmail(email string) (model.User, error) {
	var user model.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
