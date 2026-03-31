package utils

import(
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Status string  `json:"status"`
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func WriteJsonErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(JsonResponse{
		Status: "error",
		Code: statusCode,
		Message: message,
		Error: err.Error(),
	})
}

func WriteJsonSuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(JsonResponse{
		Status: "success",
		Code: statusCode,
		Message: message,
		Data: data,
	})
}
