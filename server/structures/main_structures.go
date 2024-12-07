package structures

import "github.com/golang-jwt/jwt/v5"

// User represents a user in the database
type User struct {
	ID              int    `json:"id"`
	UserUUID        string `json:"user_uuid"`
	PasswordHash    string `json:"password_hash"`
	DatetimeCreated string `json:"datetime_created"`
}

type UserPage struct {
	Username   string
	FirstName  string
	MiddleName string
	LastName   string
	City       string
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type RegisterRequest struct {
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	FirstName  string  `json:"first_name"`
	MiddleName string  `json:"middle_name"`
	LastName   string  `json:"last_name"`
	Password   string  `json:"password"`
	City       string  `json:"user_city"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type Claims struct {
	UserID string `json:"userID"`
	// Add other fields as needed
	jwt.RegisteredClaims
}

// Struct to hold the mappings
type PreferenceMapping struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

var mappings struct {
	Food  []PreferenceMapping `json:"food"`
	Hobby []PreferenceMapping `json:"hobby"`
	Music []PreferenceMapping `json:"music"`
}
