package routes

import (
	"database/sql"
	"encoding/json"
	"log"
	databaseSetup "match_me_module/database"
	middleware "match_me_module/middleware"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func AboutYou(w http.ResponseWriter, r *http.Request) {
	// Validate the JWT token from the Authorization header
	token, err := middleware.ValidateToken(r)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		log.Printf("Error in authorizing: %v", err)
		return
	}

	// Extract the user ID from the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		log.Println("Invalid token")
		return
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		http.Error(w, "Invalid user_id in token", http.StatusUnauthorized)
		log.Println("Missing or invalid user_id in token")
		return
	}

	// Parse the request body to get the "About You" field
	var requestBody struct {
		AboutYou string `json:"about_you"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("Failed to decode request body for user_id %s: %v", userID, err)
		return
	}

	if requestBody.AboutYou == "" {
		http.Error(w, "About You field cannot be empty", http.StatusBadRequest)
		log.Println("Attempted to set an empty About You field for user_id", userID)
		return
	}

	// Use UPSERT query to either insert or update the "about_you" field
	query := `
		INSERT INTO profile_info (user_uuid, about_me)
		VALUES ($1, $2)
		ON CONFLICT (user_uuid) DO UPDATE 
		SET about_me = EXCLUDED.about_me
	`
	_, err = db.Exec(query, userID, requestBody.AboutYou)
	if err != nil {
		http.Error(w, "Failed to update About You field", http.StatusInternalServerError)
		log.Printf("Error upserting About You field for user_id %s: %v", userID, err)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "About You updated successfully",
	})
}

func AboutYouGet(w http.ResponseWriter, r *http.Request) {
	// Validate the JWT token from the Authorization header
	token, err := middleware.ValidateToken(r)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		log.Printf("Error in authorizing: %v", err)
		return
	}

	// Extract the user ID from the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		log.Println("Invalid token")
		return
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		http.Error(w, "Invalid user_id in token", http.StatusUnauthorized)
		log.Println("Missing or invalid user_id in token")
		return
	}

	// Connect to the database
	db := databaseSetup.GetDB() // Assume GetDB() returns *sql.DB
	defer db.Close()

	// Query the "About You" field from the database
	var aboutYou string
	query := "SELECT about_me FROM profile_info WHERE user_uuid = $1"
	err = db.QueryRow(query, userID).Scan(&aboutYou)
	if err != nil {
		if err == sql.ErrNoRows {
			// If the user does not have an "About You" field, return an empty string
			aboutYou = ""
		} else {
			// Handle any other database errors
			http.Error(w, "Failed to fetch About You field", http.StatusInternalServerError)
			log.Printf("Error fetching About You field for user_id %s: %v", userID, err)
			return
		}
	}

	// Send the "About You" field as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"about_you": aboutYou,
	})
}

func Birthday(w http.ResponseWriter, r *http.Request) {
	// Validate the JWT token from the Authorization header
	token, err := middleware.ValidateToken(r)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		log.Printf("Error in authorizing: %v", err)
		return
	}

	// Extract the user ID from the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		log.Println("Invalid token")
		return
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		http.Error(w, "Invalid user_id in token", http.StatusUnauthorized)
		log.Println("Missing or invalid user_id in token")
		return
	}

	// Parse the request body to get the "birthday" field
	var requestBody struct {
		Birthday string `json:"birthday"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("Failed to decode request body for user_id %s: %v", userID, err)
		return
	}

	if requestBody.Birthday == "" {
		http.Error(w, "Birthday field cannot be empty", http.StatusBadRequest)
		log.Println("Attempted to set an empty birthday field for user_id", userID)
		return
	}

	// Use UPSERT query to either insert or update the "birthday" field
	query := `
		INSERT INTO user_info (user_uuid, birthday)
		VALUES (?, ?)
		ON DUPLICATE KEY UPDATE birthday = VALUES(birthday)
	`
	_, err = db.Exec(query, userID, requestBody.Birthday)
	if err != nil {
		http.Error(w, "Failed to update Birthday field", http.StatusInternalServerError)
		log.Printf("Error upserting Birthday field for user_id %s: %v", userID, err)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Birthday updated successfully",
	})
}

func BirthdayGet(w http.ResponseWriter, r *http.Request) {
	// Validate the JWT token from the Authorization header
	token, err := middleware.ValidateToken(r)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		log.Printf("Error in authorizing: %v", err)
		return
	}

	// Extract the user ID from the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		log.Println("Invalid token")
		return
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		http.Error(w, "Invalid user_id in token", http.StatusUnauthorized)
		log.Println("Missing or invalid user_id in token")
		return
	}

	// Connect to the database
	db := databaseSetup.GetDB() // Assume GetDB() returns *sql.DB
	defer db.Close()

	// Query the "birthday" field from the database
	var birthday string
	query := "SELECT birthday FROM user_info WHERE user_uuid = $1"
	err = db.QueryRow(query, userID).Scan(&birthday)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Birthday not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch Birthday", http.StatusInternalServerError)
			log.Printf("Error fetching Birthday for user_id %s: %v", userID, err)
		}
		return
	}

	// Calculate age from the birthday (assuming birthday is in YYYY-MM-DD format)
	var age int
	birthdayTime, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		http.Error(w, "Invalid birthday format", http.StatusInternalServerError)
		log.Printf("Error parsing birthday for user_id %s: %v", userID, err)
		return
	}

	// Calculate the age by comparing the birthday to today's date
	currentYear := time.Now().Year()
	age = currentYear - birthdayTime.Year()
	if birthdayTime.After(time.Now().AddDate(-age, 0, 0)) {
		age-- // Adjust age if birthday hasn't occurred yet this year
	}

	// Send the birthday and age as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"birthday": birthday,
		"age":      age,
	})
}
