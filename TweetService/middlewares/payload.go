// Add to middlewares/payload.go
package middlewares

import(
	"context"
	"encoding/json"
	"strings"
	"net/http"
	"TweetService/dto"
	"TweetService/utils"
)
func PayloadMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" || r.Method == "PUT" {
            var payload interface{}
            
        if r.Method == "POST" && r.URL.Path == "/tweet" {
                 var dto dto.CreateTweetRequestDTO
   		 if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
       			 utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid JSON", err)
      		  return
   		 }
  			  payload = dto
  	     } else if r.Method == "PUT" && strings.HasPrefix(r.URL.Path, "/tweet/") {
   			 var dto dto.UpdateTweetRequestDTO
   			 if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
       			 utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid JSON", err)
      			  return
   			 }
    payload = dto
}
            
            ctx := context.WithValue(r.Context(), PayloadKey, payload)
            next.ServeHTTP(w, r.WithContext(ctx))
        } else {
            next.ServeHTTP(w, r)
        }
    })
}