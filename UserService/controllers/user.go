package controllers

import (
	"fmt"
	"net/http"
	"userservice/dto"
	"userservice/middlewares"
	"userservice/services"
	"userservice/utils"
	"github.com/go-chi/chi/v5"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUserById called in UserController")

	userId := r.URL.Query().Get("id")

	if userId == "" {
		userId = r.Context().Value(middlewares.UserIDKey).(string)
	}

	fmt.Println("User ID from context or query: ", userId)

	if userId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "User Id is required", fmt.Errorf("missing user ID"))
		return
	}

	// convert string to int
	user, err := uc.UserService.GetUserById(userId)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch user", err)
		return
	}

	if user == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "user not found", fmt.Errorf("user with ID %s not found", userId))
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User fetched successfully", user)
	fmt.Println("user fetched successfully:", user)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Create user called in usercontroller\n")
	payload, ok := r.Context().Value(middlewares.PayloadKey).(dto.CreateUserDTO)

	if !ok {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Println("Payload received:", payload)

	user, err := uc.UserService.CreateUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "User created successfully", user)
	fmt.Println("user created successfully:", user)
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Login user called in usercontroller\n")
	payload := r.Context().Value(middlewares.PayloadKey).(dto.LoginUserRequestDTO)

	fmt.Println("Payload received:", payload)

	jwtToken, err := uc.UserService.LoginUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to login user", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully", jwtToken)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Delete user called in usercontroller\n")

	userId := chi.URLParam(r, "id")

	if userId == ""{
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"User id required",fmt.Errorf("missing user id"))
		return
	}
	err := uc.UserService.DeleteUser(userId)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to delete user", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User deleted successfully", nil)
}
