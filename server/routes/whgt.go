package routes

import (
	"database/sql"
	"encoding/json"
	"log"
	databaseSetup "match_me_module/database"
	middleware "match_me_module/middleware"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func WeightDistance(w http.ResponseWriter, r *http.Request) {
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

	// Parse the request body to get the number
	var requestBody struct {
		Number float64 `json:"number"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	// Validate the number
	if requestBody.Number <= 0 {
		http.Error(w, "Invalid number: must be greater than 0", http.StatusBadRequest)
		log.Println("Invalid number provided")
		return
	}

	// Connect to the database
	db := databaseSetup.GetDB() // Assume GetDB() returns *sql.DB

	// Update the weigh_distance column for the given user
	updateQuery := "UPDATE weights SET weigh_distance = $1 WHERE user_uuid = $2"
	_, err = db.Exec(updateQuery, requestBody.Number, userID)
	if err != nil {
		http.Error(w, "Failed to update weigh_distance", http.StatusInternalServerError)
		log.Printf("Error updating weigh_distance: %v", err)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("weigh_distance updated successfully"))
}

func WeightAge(w http.ResponseWriter, r *http.Request) {
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

	// Parse the request body to get the number
	var requestBody struct {
		Number float64 `json:"number"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	// Validate the number
	if requestBody.Number <= 0 {
		http.Error(w, "Invalid number: must be greater than 0", http.StatusBadRequest)
		log.Println("Invalid number provided")
		return
	}

	// Connect to the database
	db := databaseSetup.GetDB() // Assume GetDB() returns *sql.DB

	// Update the weigh_age column for the given user
	updateQuery := "UPDATE weights SET weigh_age = $1 WHERE user_uuid = $2"
	_, err = db.Exec(updateQuery, requestBody.Number, userID)
	if err != nil {
		http.Error(w, "Failed to update weigh_age", http.StatusInternalServerError)
		log.Printf("Error updating weigh_age: %v", err)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("weigh_ange updated successfully"))
}

func WeightFood(w http.ResponseWriter, r *http.Request) {
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

	// Parse the request body to get the number
	var requestBody struct {
		Number float64 `json:"number"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	// Validate the number
	if requestBody.Number <= 0 {
		http.Error(w, "Invalid number: must be greater than 0", http.StatusBadRequest)
		log.Println("Invalid number provided")
		return
	}

	// Connect to the database
	db := databaseSetup.GetDB() // Assume GetDB() returns *sql.DB

	// Update the weigh_food column for the given user
	updateQuery := "UPDATE weights SET weigh_food = $1 WHERE user_uuid = $2"
	_, err = db.Exec(updateQuery, requestBody.Number, userID)
	if err != nil {
		http.Error(w, "Failed to update weigh_food", http.StatusInternalServerError)
		log.Printf("Error updating weigh_food: %v", err)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("weigh_food updated successfully"))
}

func WeightHobbies(w http.ResponseWriter, r *http.Request) {
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

	// Parse the request body to get the number
	var requestBody struct {
		Number float64 `json:"number"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	// Validate the number
	if requestBody.Number <= 0 {
		http.Error(w, "Invalid number: must be greater than 0", http.StatusBadRequest)
		log.Println("Invalid number provided")
		return
	}

	// Connect to the database
	db := databaseSetup.GetDB() // Assume GetDB() returns *sql.DB

	// Update the weigh_hobbies column for the given user
	updateQuery := "UPDATE weights SET weigh_hobbies = $1 WHERE user_uuid = $2"
	_, err = db.Exec(updateQuery, requestBody.Number, userID)
	if err != nil {
		http.Error(w, "Failed to update weigh_hobbies", http.StatusInternalServerError)
		log.Printf("Error updating weigh_hobbies: %v", err)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("weigh_hobbies updated successfully"))
}

func WeightMusic(w http.ResponseWriter, r *http.Request) {
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

	// Parse the request body to get the number
	var requestBody struct {
		Number float64 `json:"number"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	// Validate the number
	if requestBody.Number <= 0 {
		http.Error(w, "Invalid number: must be greater than 0", http.StatusBadRequest)
		log.Println("Invalid number provided")
		return
	}

	// Connect to the database
	db := databaseSetup.GetDB() // Assume GetDB() returns *sql.DB

	// Update the weigh_music column for the given user
	updateQuery := "UPDATE weights SET weigh_music = $1 WHERE user_uuid = $2"
	_, err = db.Exec(updateQuery, requestBody.Number, userID)
	if err != nil {
		http.Error(w, "Failed to update weigh_music", http.StatusInternalServerError)
		log.Printf("Error updating weigh_music: %v", err)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("weigh_music updated successfully"))
}

func WeightGet(w http.ResponseWriter, r *http.Request) {
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

	// Query the weights for the user
	var weights struct {
		WeighDistance float64 `json:"weigh_distance"`
		WeighAge      float64 `json:"weigh_age"`
		WeighFood     float64 `json:"weigh_food"`
		WeighHobbies  float64 `json:"weigh_hobbies"`
		WeighMusic    float64 `json:"weigh_music"`
	}

	query := `
        SELECT weigh_distance, weigh_age, weigh_food, weigh_hobbies, weigh_music 
        FROM weights 
        WHERE user_uuid = $1
    `
	err = db.QueryRow(query, userID).Scan(
		&weights.WeighDistance,
		&weights.WeighAge,
		&weights.WeighFood,
		&weights.WeighHobbies,
		&weights.WeighMusic,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No weights found for user", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to query weights", http.StatusInternalServerError)
		}
		log.Printf("Error querying weights: %v", err)
		return
	}

	// Send the weights as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weights)
}
