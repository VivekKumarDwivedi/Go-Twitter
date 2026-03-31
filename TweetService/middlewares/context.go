package middlewares

type ctxKey string

const (
	UserIDKey  ctxKey = "userID"
	PayloadKey ctxKey = "payload"
)
