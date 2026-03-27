package middlewares

import (
	"context"
	"net/http"
	"userservice/dto"
	"userservice/utils"
)

func CreateUserRequestVaildator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.CreateUserDTO

		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invaild request body", err)
			return
		}
		ctx := context.WithValue(r.Context(), PayloadKey, payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
