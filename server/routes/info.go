package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	middleware "match_me_module/middleware"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

// UserInfo handles the /api/user endpoint
func UserInfo(w http.ResponseWriter, r *http.Request) {
	// Extract and decode the JWT token from Authorization header
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
	userID := claims["user_id"].(string)

	// Fetch user info from the database
	userInfo, err := fetchUserInfo(userID)
	if err != nil {
		http.Error(w, "Error fetching user info: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Invalid token: %v", err)
		return
	}

	// Return the user data as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userInfo)
}

// fetchUserInfo queries the database for user data based on the user_id (user_uuid)
func fetchUserInfo(userID string) (map[string]interface{}, error) {
	var firstName, middleName, lastName, city string

	// Query the database for the user's information
	err := db.QueryRow(`
		SELECT i.first_name, i.middle_name, i.last_name, d.user_city 
		FROM user_info i 
		JOIN user_data d ON i.user_uuid = d.user_uuid  
		WHERE i.user_uuid=$1`, userID).Scan(&firstName, &middleName, &lastName, &city)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("user not found: %v", err)
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	// If no middle name, set it as an empty string
	if middleName == "" {
		middleName = "(N/A)"
	}

	// Return user data as a map

	userInfo := map[string]interface{}{
		"firstName":  firstName,
		"middleName": middleName,
		"lastName":   lastName,
		"city":       city,
		"user_id":    userID,
	}

	return userInfo, nil
}
