package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mcgtrt/azure-graphql/store"
)

func JWTAuthenticate(ctx context.Context, store store.EmployeeStorer, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, ok := r.Header["X-Api-Token"]
		if !ok {
			writeJSONUnauthorised(w)
			return
		}

		claims := getClaims(token[0])
		if claims == nil {
			writeJSONUnauthorised(w)
			return
		}

		var (
			expFloat = claims["expires"].(float64)
			expInt   = int64(expFloat)
		)
		if time.Now().Unix() > expInt {
			writeJSONTokenExpired(w)
			return
		}

		employeeID := claims["employeeID"].(string)
		employee, err := store.GetEmployeeByID(ctx, employeeID)
		if err != nil {
			writeJSONUnauthorised(w)
			return
		}
		if employee.Email != claims["email"].(string) {
			writeJSONUnauthorised(w)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func getClaims(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unauthorised")
		}
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "superstrongsecretpassword"
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	}
	return nil
}
