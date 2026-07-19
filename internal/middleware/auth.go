package middleware

import (
	"context"
	"net/http"
)

type contextKey string

type Session struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
}

const SessionIDKey contextKey = "session_id"
const UserKey contextKey = "user"

func CookieMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("session")
		if err == nil {
			ctx := context.WithValue(
				r.Context(),
				SessionIDKey,
				cookie.Value,
			)

			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}
