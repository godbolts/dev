package routes

import (
	"encoding/json"
	"log"
	middleware "match_me_module/middleware"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func EditUsername(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("Invalid token: %v", err)
		return
	}

	// Extract the user_id (user_uuid) from the token's claims
	userID, ok := claims["user_id"].(string)
	if !ok {
		http.Error(w, "Invalid user_id in token", http.StatusUnauthorized)
		log.Println("Missing or invalid user_id in token")
		return
	}

	// Parse the request body to get the new username
	var requestBody struct {
		Username string `json:"username"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("Failed to decode request body: %v", err)
		return
	}

	if requestBody.Username == "" {
		http.Error(w, "Username cannot be empty", http.StatusBadRequest)
		log.Println("Attempted to set an empty username")
		return
	}

	// Update the username in the database
	query := `UPDATE user_info SET username = ? WHERE user_uuid = ?`
	_, err = db.Exec(query, requestBody.Username, userID)
	if err != nil {
		http.Error(w, "Failed to update username", http.StatusInternalServerError)
		log.Printf("Error updating username for user_id %s: %v", userID, err)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Username updated successfully",
	})
}

func EditEmail(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("Invalid token: %v", err)
		return
	}

	// Extract the user_id (user_uuid) from the token's claims
	userID, ok := claims["user_id"].(string)
	if !ok {
		http.Error(w, "Invalid user_id in token", http.StatusUnauthorized)
		log.Println("Missing or invalid user_id in token")
		return
	}

	// Parse the request body to get the new email
	var requestBody struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("Failed to decode request body: %v", err)
		return
	}

	if requestBody.Email == "" {
		http.Error(w, "Email cannot be empty", http.StatusBadRequest)
		log.Println("Attempted to set an empty email")
		return
	}

	// Update the email in the database
	query := `UPDATE user_info SET email = ? WHERE user_uuid = ?`
	_, err = db.Exec(query, requestBody.Email, userID)
	if err != nil {
		http.Error(w, "Failed to update email", http.StatusInternalServerError)
		log.Printf("Error updating email for user_id %s: %v", userID, err)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Email updated successfully",
	})
}

func EditFirst(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("Invalid token: %v", err)
		return
	}

	// Extract the user_id (user_uuid) from the token's claims
	userID, ok := claims["user_id"].(string)
	if !ok {
		http.Error(w, "Invalid user_id in token", http.StatusUnauthorized)
		log.Println("Missing or invalid user_id in token")
		return
	}

	// Parse the request body to get the new FirstName
	var requestBody struct {
		FirstName string `json:"first_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("Failed to decode request body: %v", err)
		return
	}

	if requestBody.FirstName == "" {
		http.Error(w, "first_name cannot be empty", http.StatusBadRequest)
		log.Println("Attempted to set an empty first_name")
		return
	}

	// Update the FirstName in the database
	query := `UPDATE user_info SET first_name = ? WHERE user_uuid = ?`
	_, err = db.Exec(query, requestBody.FirstName, userID)
	if err != nil {
		http.Error(w, "Failed to update first_name", http.StatusInternalServerError)
		log.Printf("Error updating first_name for user_id %s: %v", userID, err)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "FirstName updated successfully",
	})
}

func EditMiddle(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("Invalid token: %v", err)
		return
	}

	// Extract the user_id (user_uuid) from the token's claims
	userID, ok := claims["user_id"].(string)
	if !ok {
		http.Error(w, "Invalid user_id in token", http.StatusUnauthorized)
		log.Println("Missing or invalid user_id in token")
		return
	}

	// Parse the request body to get the new MiddleName
	var requestBody struct {
		MiddleName string `json:"middle_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("Failed to decode request body: %v", err)
		return
	}

	if requestBody.MiddleName == "" {
		http.Error(w, "middle_name cannot be empty", http.StatusBadRequest)
		log.Println("Attempted to set an empty middle_name")
		return
	}

	// Update the Middle Name in the database
	query := `UPDATE user_info SET middle_name = ? WHERE user_uuid = ?`
	_, err = db.Exec(query, requestBody.MiddleName, userID)
	if err != nil {
		http.Error(w, "Failed to update middle_name", http.StatusInternalServerError)
		log.Printf("Error updating middle_name for user_id %s: %v", userID, err)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Middle name updated successfully",
	})
}

func EditLast(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("Invalid token: %v", err)
		return
	}

	// Extract the user_id (user_uuid) from the token's claims
	userID, ok := claims["user_id"].(string)
	if !ok {
		http.Error(w, "Invalid user_id in token", http.StatusUnauthorized)
		log.Println("Missing or invalid user_id in token")
		return
	}

	// Parse the request body to get the new Last Name
	var requestBody struct {
		LastName string `json:"last_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("Failed to decode request body: %v", err)
		return
	}

	if requestBody.LastName == "" {
		http.Error(w, "last_name cannot be empty", http.StatusBadRequest)
		log.Println("Attempted to set an empty last_name")
		return
	}

	// Update the Last Name in the database
	query := `UPDATE user_info SET last_name = ? WHERE user_uuid = ?`
	_, err = db.Exec(query, requestBody.LastName, userID)
	if err != nil {
		http.Error(w, "Failed to update last_name", http.StatusInternalServerError)
		log.Printf("Error updating last_name for user_id %s: %v", userID, err)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Last name updated successfully",
	})
}

func EditPassword(w http.ResponseWriter, r *http.Request) {
	// Validate the JWT token from the Authorization header
	token, err := middleware.ValidateToken(r)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		log.Printf("Authorization error: %v", err)
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

	// Parse the request body to get the new password
	var requestBody struct {
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("Failed to decode request body for user_id %s: %v", userID, err)
		return
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		log.Printf("Error hashing password for user_id %s: %v", userID, err)
		return
	}

	// Update the password in the database
	query := `UPDATE user_table SET password_hash = ? WHERE user_uuid = ?`
	_, err = db.Exec(query, hashedPassword, userID)
	if err != nil {
		http.Error(w, "Error updating password", http.StatusInternalServerError)
		log.Printf("Error updating password for user_id %s: %v", userID, err)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Password updated successfully",
	})
}

func EditCity(w http.ResponseWriter, r *http.Request) {
	// Validate the JWT token from the Authorization header
	token, err := middleware.ValidateToken(r)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		log.Printf("Authorization error: %v", err)
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

	// Parse the request body for city and coordinates
	var requestBody struct {
		City      string  `json:"city"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("Failed to decode request body for user_id %s: %v", userID, err)
		return
	}

	// Update the user's city and location in the database
	query := `
		UPDATE user_data 
		SET user_city = $1, 
		    register_location = ST_SetSRID(ST_MakePoint($2, $3), 4326)
		WHERE user_uuid = $4
	`
	_, err = db.Exec(query, requestBody.City, requestBody.Longitude, requestBody.Latitude, userID)
	if err != nil {
		http.Error(w, "Error updating user location", http.StatusInternalServerError)
		log.Printf("Error updating user location for user_id %s: %v", userID, err)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "City and location updated successfully",
	})
}
