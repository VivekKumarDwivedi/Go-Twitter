package middlewares

import (
	"fmt"
	"context"
	"net/http"
	"userservice/dto"
	"userservice/utils"
)

func UserLoginRequestVaildator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.LoginUserRequestDTO

		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}

		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation error", err)
			return
		}
		fmt.Println("Payload recived for login",payload)
		cx := context.WithValue(r.Context(), PayloadKey, payload)
		next.ServeHTTP(w, r.WithContext(cx))
	})
}
func CreateUserRequestVaildator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.CreateUserDTO

		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invaild request body", err)
			return
		}

		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation error", err)
			return
		}

		ctx := context.WithValue(r.Context(), PayloadKey, payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UpdateUserRequestVaildator(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var payload dto.UpdateUserRequestDTO

		if err := utils.ReadJsonBody(r, &payload); err != nil{
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}

		if err := utils.Validator.Struct(payload); err != nil{
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation error", err)
			return
		}

		ctx := context.WithValue(r.Context(), PayloadKey, payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}