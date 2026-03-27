package services

import (
	"fmt"
	env"userservice/config/env"
	db "userservice/db/repositories"
	"userservice/dto"
	"userservice/models"
	"userservice/utils"
	jwt "github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById(id string) (*models.User, error)
	CreateUser(payload *dto.CreateUserDTO) (*models.User, error)
	LoginUser(payload *dto.LoginUserRequestDTO) (string, error)
	DeleteUser(id string) error
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
func (r *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDTO) (string, error) {
	fmt.Println("Logging in user in UserService")
	
	// TODO: Implement login logic
	// 1. Find user by email
	// 2. Verify password
	// 3. Generate JWT token
	// 4. Return token

	email := payload.Email
	password := payload.Password

	user, err := r.userRepository.GetUserByEmail(email)

	if err != nil {
		fmt.Println("Error fetching user by email:",err)
		return "", err
	}

	if user == nil {
		fmt.Println("No user found with given email")
		return "",fmt.Errorf("no user found with given email: %s", email)
	}
	
	isPasswordVaild := utils.CheckPasswordHash(password, user.Password)

	if !isPasswordVaild {
		fmt.Println("Password does not match")
		return "",nil
	}

	jwtPayload := jwt.MapClaims{
		"email":user.Email,
		"id":user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)
	
	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET","TOKEN")))

	if err != nil {
		fmt.Println("Error generating token:", err)
		return "", err
	}
	fmt.Println("Token generated:", tokenString)
	
	return tokenString, nil
}

func (r *UserServiceImpl) DeleteUser(id string) error {
	fmt.Println("Deleting user in UserService")
		
	return r.userRepository.DeleteUser(id)
	
}
