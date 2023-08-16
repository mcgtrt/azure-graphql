package api

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mcgtrt/azure-graphql/store"
	"github.com/mcgtrt/azure-graphql/types"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	store store.EmployeeStorer
}

func NewAuthHandler(store store.EmployeeStorer) *AuthHandler {
	return &AuthHandler{
		store: store,
	}
}

func (h *AuthHandler) HandleAuth(w http.ResponseWriter, r *http.Request) {
	var authParams types.AuthParams
	if err := json.NewDecoder(r.Body).Decode(&authParams); err != nil {
		writeJSONBadRequest(w)
		return
	}
	employee, err := h.store.GetEmployeeByEmail(context.Background(), authParams.Email)
	if err != nil {
		writeJSONUnauthorised(w)
		return
	}
	if !isPasswordValid(employee.EncryptedPassword, authParams.Password) {
		writeJSONInvalidCredentials(w)
		return
	}
	resp := types.AuthResponse{
		Email: employee.Email,
		Token: CreateTokenFromAuthEmployee(employee),
	}
	writeJSON(w, http.StatusOK, resp)
}

func CreateTokenFromAuthEmployee(employee *types.AuthEmployee) string {
	expires := time.Now().Add(time.Hour * 4).Unix()
	claims := jwt.MapClaims{
		"employeeID": employee.ID,
		"email":      employee.Email,
		"expires":    expires,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return ""
	}
	return tokenString
}

func isPasswordValid(encpw, pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encpw), []byte(pw)) == nil
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeJSONUnauthorised(w http.ResponseWriter) {
	writeJSON(w, http.StatusUnauthorized, "unauthorised")
}

func writeJSONTokenExpired(w http.ResponseWriter) {
	writeJSON(w, http.StatusUnauthorized, "token expired")
}

func writeJSONBadRequest(w http.ResponseWriter) {
	writeJSON(w, http.StatusBadRequest, "bad request")
}

func writeJSONInvalidCredentials(w http.ResponseWriter) {
	writeJSON(w, http.StatusBadRequest, "invalid credentials")
}
