package router

import (
	"userservice/controllers"
	"userservice/middlewares"
	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	UserController controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) Router {
	return &UserRouter{
		UserController: *_userController,
	}
}

func (ur *UserRouter) Register(r chi.Router) {
	r.With(middlewares.UserLoginRequestVaildator).Post("/login", ur.UserController.LoginUser)
	r.With(middlewares.CreateUserRequestVaildator).Post("/signup", ur.UserController.CreateUser)
	r.With(middlewares.JWTAuthMiddleware).Get("/user/{id}", ur.UserController.GetUserByID)
	r.Delete("/user/{id}", ur.UserController.DeleteUser)
}
