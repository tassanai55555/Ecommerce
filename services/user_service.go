package services

import (
	"ecommerce-backend/database"
	"ecommerce-backend/dto"
	"ecommerce-backend/models"
)

type UserService struct{}

func (s *UserService) CreateUser(data dto.CreateUserRequest) (*dto.UserResponse, error) {
	user := models.User{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password, // later: hash this
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}