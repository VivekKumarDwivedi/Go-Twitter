package middlewares

type ctxKey string

const (
	UserIDKey  ctxKey = "userID"
	EmailKey   ctxKey = "email"
	PayloadKey ctxKey = "payload"
)
