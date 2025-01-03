package middleware

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"net/http"
	"strings"
)

type userContextKey int

const (
	userCtxKey = iota
)

func Auth(client *auth.Client) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !strings.HasPrefix(r.Header.Get("Authorization"), "Bearer ") {
				fmt.Printf("No auth header\n")
				next.ServeHTTP(w, r)
				return
			}
			idToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
			token, err := client.VerifyIDToken(r.Context(), idToken)
			if err != nil || token.Claims["email_verified"] != true || token.Firebase.SignInProvider != "google.com" {
				fmt.Printf("Error verifying token: %v\n", err)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), userCtxKey, token)
			//fmt.Printf("Authenticated user object: %+v\n", token)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUser(ctx context.Context) *auth.Token {
	user, ok := ctx.Value(userCtxKey).(*auth.Token)
	if !ok {
		return nil
	}
	return user
}
