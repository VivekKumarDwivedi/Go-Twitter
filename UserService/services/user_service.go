package services

import (
	"fmt"
	db "userservice/db/repositories"
	"userservice/dto"
	"userservice/models"
	"userservice/utils"
)

type UserService interface {
	GetUserById(id string) (*models.User, error)
	CreateUser(payload *dto.CreateUserDTO) (*models.User, error)
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (r *UserServiceImpl) GetUserById(id string) (*models.User, error) {
	fmt.Println("fetching user in userservice")
	user, err := r.userRepository.GetUserByID(id)

	if err != nil {
		fmt.Println("error fetching user:", err)
		return nil, err
	}
	return user, nil

}

func (r *UserServiceImpl) CreateUser(payload *dto.CreateUserDTO) (*models.User, error) {

	fmt.Println("Creating user in UserService")

	// Hash the password using utils.hashpassword

	hashedPassword, err := utils.HashPassword(payload.Password)

	if err != nil {
		fmt.Println("error hashing password", err)
		return nil, err
	}

	user, err := r.userRepository.CreateUser(payload.Username, payload.Email, hashedPassword)

	if err != nil {
		fmt.Println("Error creating user:", err)
		return nil, err
	}

	return user, nil

}
