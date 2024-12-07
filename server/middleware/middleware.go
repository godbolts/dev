package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// jwtSecretKey holds the JWT secret key
var jwtSecretKey []byte

// LoadJWTSecretKey loads the JWT secret key from environment variables or .env file
func LoadJWTSecretKey() []byte {
	// Load environment variables from .env file (optional)
	err := godotenv.Load("../server/config.env")
	if err != nil {
		log.Printf("Error loading config.env file: %v", err)
	}

	// Get JWT_SECRET_KEY from environment variables
	jwtSecretKeyEnv := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKeyEnv == "" {
		log.Fatal("JWT_SECRET_KEY is not set in environment variables")
	}

	// Set and return the secret key
	jwtSecretKey = []byte(jwtSecretKeyEnv)
	return jwtSecretKey
}

// GetJWTSecretKey provides access to the loaded JWT secret key
func GetJWTSecretKey() []byte {
	if jwtSecretKey == nil {
		LoadJWTSecretKey()
	}
	return jwtSecretKey
}

// ValidateToken validates the JWT token from the Authorization header.
func ValidateToken(r *http.Request) (*jwt.Token, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("missing authorization header")
	}

	// Extract the token from the Authorization header: "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, fmt.Errorf("invalid token format")
	}

	tokenString := parts[1]

	// Parse the token and validate it with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
