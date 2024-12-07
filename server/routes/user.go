package routes

import (
	"database/sql"
	"encoding/json"
	"log"
	databaseSetup "match_me_module/database"
	middleware "match_me_module/middleware"
	"match_me_module/structures"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var db = databaseSetup.GetDB()

// Define your JWT secret key (store securely, e.g., in environment variables)
var jwtSecretKey = middleware.GetJWTSecretKey()

// Retrieves users from the database and sends them as JSON.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, user_uuid, password_hash, datetime_created FROM user_table")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []structures.User

	for rows.Next() {
		var user structures.User
		if err := rows.Scan(&user.ID, &user.UserUUID, &user.PasswordHash, &user.DatetimeCreated); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Login function in the API.
func Login(w http.ResponseWriter, r *http.Request) {
	var loginReq structures.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error in decoding: %v", err)
		return
	}

	if loginReq.Username == "" || loginReq.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	var storedHash string

	// First, check if the user exists
	err := db.QueryRow("SELECT COUNT(1) FROM user_info WHERE username=$1", loginReq.Username).Scan(&storedHash)
	if err != nil {
		// Handle database error
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Printf("Error checking user existence: %v", err)
		return
	}

	// If user doesn't exist, return an unauthorized error
	if storedHash == "0" {
		http.Error(w, "User not found", http.StatusUnauthorized)
		log.Printf("User not found: %s", loginReq.Username)
		return
	}

	// User exists, now try to get the password hash
	err = db.QueryRow("SELECT password_hash FROM user_table a JOIN user_info b ON a.user_uuid = b.user_uuid WHERE b.username=$1", loginReq.Username).Scan(&storedHash)
	if err != nil {
		// Handle errors retrieving the password hash
		if err == sql.ErrNoRows {
			http.Error(w, "Password not found", http.StatusUnauthorized)
			log.Printf("Password not found for user: %s", loginReq.Username)
		} else {
			http.Error(w, "Server error", http.StatusInternalServerError)
			log.Printf("Error retrieving password hash: %v", err)
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(loginReq.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		log.Printf("Error, wrong hash password: %v", err)
		return
	}
	var user_id string

	err = db.QueryRow("SELECT user_uuid FROM user_info WHERE username=$1", loginReq.Username).Scan(&user_id)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found for username: %s", loginReq.Username)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		log.Printf("Query error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	token, err := GenerateJWT(user_id)
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(structures.LoginResponse{
		Status:  "success",
		Message: "Login successful",
		Token:   token,
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	var registerReq structures.RegisterRequest

	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&registerReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error in decoding: %v", err)
		return
	}

	// Check if required fields are not empty
	if registerReq.Username == "" || registerReq.Email == "" || registerReq.FirstName == "" || registerReq.MiddleName == "" ||
		registerReq.LastName == "" || registerReq.Password == "" || registerReq.City == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		log.Printf("Missing required fields in registration request")
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		log.Printf("Error hashing password: %v", err)
		return
	}

	// Generate UUID for the new user
	userUUID := uuid.New()

	// Create a timestamp for when the user registers
	datetimeCreated := time.Now()

	// Insert data into the `user_table`
	_, err = db.Exec("INSERT INTO user_table (user_uuid, password_hash, datetime_created) VALUES ($1, $2, $3)", userUUID, hashedPassword, datetimeCreated)
	if err != nil {
		http.Error(w, "Error saving user data", http.StatusInternalServerError)
		log.Printf("Error saving user data: %v", err)
		return
	}

	// Insert data into the `user_info` table
	_, err = db.Exec("INSERT INTO user_info (user_uuid, username, email, first_name, middle_name, last_name) VALUES ($1, $2, $3, $4, $5, $6)", userUUID, registerReq.Username, registerReq.Email, registerReq.FirstName, registerReq.MiddleName, registerReq.LastName)
	if err != nil {
		http.Error(w, "Error saving user info", http.StatusInternalServerError)
		log.Printf("Error saving user info: %v", err)
		return
	}

	// Insert data into the `user_data` table with latitude and longitude (home city)
	_, err = db.Exec("INSERT INTO user_data (user_uuid, user_city, register_location) VALUES ($1, $2, ST_SetSRID(ST_MakePoint($3, $4), 4326))", userUUID, registerReq.City, registerReq.Longitude, registerReq.Latitude)
	if err != nil {
		http.Error(w, "Error saving user data location", http.StatusInternalServerError)
		log.Printf("Error saving user data location: %v", err)
		return
	}

	// Insert default weights into the `weights` table
	_, err = db.Exec("INSERT INTO weights (user_uuid, weigh_distance, weigh_age, weigh_food, weigh_hobbies, weigh_music) VALUES ($1, $2, $3, $4, $5, $6)",
		userUUID, 1, 1, 1, 1, 1)
	if err != nil {
		http.Error(w, "Error saving user weights", http.StatusInternalServerError)
		log.Printf("Error saving user weights: %v", err)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "User registered successfully"}`))
}

func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}
