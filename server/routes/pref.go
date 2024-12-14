package routes

import (
	"database/sql"
	"encoding/json"
	"log"
	databaseSetup "match_me_module/database"
	middleware "match_me_module/middleware"
	"match_me_module/structures"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func FoodPref(w http.ResponseWriter, r *http.Request) {
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

	// Parse the request body to get the "code" and "remove" flag
	var requestBody struct {
		Code   string `json:"code"`
		Remove bool   `json:"isUnchecked"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	code := requestBody.Code
	if code == "" {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}

	// Connect to the database
	db := databaseSetup.GetDB() // Assume GetDB() returns *sql.DB

	// Change the currentData type to sql.NullString
	var currentData sql.NullString

	query := "SELECT food_myvariabledata FROM profile_info WHERE user_uuid = $1"
	err = db.QueryRow(query, userID).Scan(&currentData)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database query error", http.StatusInternalServerError)
			log.Printf("Error querying database: %v", err)
		}
		return
	}

	// Handle NULL case in the database
	var codes []string
	if currentData.Valid && currentData.String != "" {
		codes = strings.Split(currentData.String, ",")
	}

	// Add or remove the code from the list
	var updatedCodes []string
	codeFound := false
	codeExists := false

	// Check for existing code presence and update the list
	for _, existingCode := range codes {
		if existingCode == code {
			if requestBody.Remove {
				codeFound = true
				continue // Skip the code to remove it
			}
			codeExists = true
		}
		updatedCodes = append(updatedCodes, existingCode)
	}

	if requestBody.Remove && !codeFound {
		http.Error(w, "Code not found, nothing to remove", http.StatusNotFound)
		return
	}

	if !requestBody.Remove && codeExists {
		http.Error(w, "Code already exists", http.StatusBadRequest)
		return
	}

	if !requestBody.Remove {
		// Add the code if it doesn't already exist
		updatedCodes = append(updatedCodes, code)
	}

	// Join the updated codes into a string
	newData := strings.Join(updatedCodes, ",")

	// Update the database
	updateQuery := "UPDATE profile_info SET food_myvariabledata = $1 WHERE user_uuid = $2"
	_, err = db.Exec(updateQuery, newData, userID)
	if err != nil {
		http.Error(w, "Failed to update database", http.StatusInternalServerError)
		log.Printf("Error updating database: %v", err)
		return
	}

	// Respond with success
	if requestBody.Remove {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Code removed successfully"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Code added successfully"))
	}
}

func HobbyPref(w http.ResponseWriter, r *http.Request) {
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

	// Parse the request body to get the "code" and "remove" flag
	var requestBody struct {
		Code   string `json:"code"`
		Remove bool   `json:"isUnchecked"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	code := requestBody.Code
	if code == "" {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}

	// Connect to the database
	db := databaseSetup.GetDB() // Assume GetDB() returns *sql.DB

	// Query the current interests_myvariabledata for the user
	var currentData sql.NullString
	query := "SELECT hobbies_myvariabledata FROM profile_info WHERE user_uuid = $1"
	err = db.QueryRow(query, userID).Scan(&currentData)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database query error", http.StatusInternalServerError)
			log.Printf("Error querying database: %v", err)
		}
		return
	}

	// Handle NULL case in the database
	var codes []string
	if currentData.Valid && currentData.String != "" {
		codes = strings.Split(currentData.String, ",")
	}

	// Add or remove the code from the list
	var updatedCodes []string
	codeFound := false
	codeExists := false

	// Check for existing code presence and update the list
	for _, existingCode := range codes {
		if existingCode == code {
			if requestBody.Remove {
				codeFound = true
				continue // Skip the code to remove it
			}
			codeExists = true
		}
		updatedCodes = append(updatedCodes, existingCode)
	}

	if requestBody.Remove && !codeFound {
		http.Error(w, "Code not found, nothing to remove", http.StatusNotFound)
		return
	}

	if !requestBody.Remove && codeExists {
		http.Error(w, "Code already exists", http.StatusBadRequest)
		return
	}

	if !requestBody.Remove {
		// Add the code if it doesn't already exist
		updatedCodes = append(updatedCodes, code)
	}

	// Join the updated codes into a string
	newData := strings.Join(updatedCodes, ",")
	updateQuery := "UPDATE profile_info SET hobbies_myvariabledata = $1 WHERE user_uuid = $2"
	_, err = db.Exec(updateQuery, newData, userID)
	if err != nil {
		http.Error(w, "Failed to update database", http.StatusInternalServerError)
		log.Printf("Error updating database: %v", err)
		return
	}

	if requestBody.Remove {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Code removed successfully"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Code added successfully"))
	}
}

func MusicPref(w http.ResponseWriter, r *http.Request) {
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

	// Parse the request body to get the "code" and "remove" flag
	var requestBody struct {
		Code   string `json:"code"`
		Remove bool   `json:"isUnchecked"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	code := requestBody.Code
	if code == "" {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}

	// Connect to the database
	db := databaseSetup.GetDB() // Assume GetDB() returns *sql.DB

	// Query the current interests_myvariabledata for the user
	var currentData sql.NullString
	query := "SELECT music_myvariabledata FROM profile_info WHERE user_uuid = $1"
	err = db.QueryRow(query, userID).Scan(&currentData)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database query error", http.StatusInternalServerError)
			log.Printf("Error querying database: %v", err)
		}
		return
	}

	// Handle NULL case in the database
	var codes []string
	if currentData.Valid && currentData.String != "" {
		codes = strings.Split(currentData.String, ",")
	}

	// Add or remove the code from the list
	var updatedCodes []string
	codeFound := false
	codeExists := false

	// Check for existing code presence and update the list
	for _, existingCode := range codes {
		if existingCode == code {
			if requestBody.Remove {
				codeFound = true
				continue // Skip the code to remove it
			}
			codeExists = true
		}
		updatedCodes = append(updatedCodes, existingCode)
	}

	if requestBody.Remove && !codeFound {
		http.Error(w, "Code not found, nothing to remove", http.StatusNotFound)
		return
	}

	if !requestBody.Remove && codeExists {
		http.Error(w, "Code already exists", http.StatusBadRequest)
		return
	}

	if !requestBody.Remove {
		// Add the code if it doesn't already exist
		updatedCodes = append(updatedCodes, code)
	}

	// Join the updated codes into a string
	newData := strings.Join(updatedCodes, ",")
	updateQuery := "UPDATE profile_info SET music_myvariabledata = $1 WHERE user_uuid = $2"
	_, err = db.Exec(updateQuery, newData, userID)
	if err != nil {
		http.Error(w, "Failed to update database", http.StatusInternalServerError)
		log.Printf("Error updating database: %v", err)
		return
	}

	if requestBody.Remove {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Code removed successfully"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Code added successfully"))
	}
}

func PrefGet(w http.ResponseWriter, r *http.Request) {
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

	// Query the current preferences for the user
	var hobbies, music, food *string // Use pointers to capture NULL values

	query := `
        SELECT hobbies_myvariabledata, music_myvariabledata, food_myvariabledata 
        FROM profile_info 
        WHERE user_uuid = $1
    `
	err = db.QueryRow(query, userID).Scan(&hobbies, &music, &food)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database query error", http.StatusInternalServerError)
		log.Printf("Error querying database: %v", err)
		return
	}

	// Handle NULL values by converting pointers to empty strings
	preferences := struct {
		Hobbies string `json:"hobbies_myvariabledata"`
		Music   string `json:"music_myvariabledata"`
		Food    string `json:"food_myvariabledata"`
	}{
		Hobbies: defaultString(hobbies),
		Music:   defaultString(music),
		Food:    defaultString(food),
	}

	// Send the preferences as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(preferences)
}

// defaultString converts a nil pointer to an empty string, or returns the value if not nil
func defaultString(input *string) string {
	if input == nil {
		return ""
	}
	return *input
}

func PrefMappingGet(w http.ResponseWriter, r *http.Request) {
	// Connect to the database
	db := databaseSetup.GetDB() // Assume GetDB() returns *sql.DB

	// Helper function to fetch mapping data
	fetchMapping := func(query string, target *[]structures.PreferenceMapping) error {
		rows, err := db.Query(query)
		if err != nil {
			return err
		}

		for rows.Next() {
			var item structures.PreferenceMapping
			if err := rows.Scan(&item.Code, &item.Description); err != nil {
				return err
			}
			*target = append(*target, item)
		}
		return rows.Err()
	}

	var mappings struct {
		Food  []structures.PreferenceMapping `json:"food"`
		Hobby []structures.PreferenceMapping `json:"hobby"`
		Music []structures.PreferenceMapping `json:"music"`
	}

	// Fetch food mappings
	if err := fetchMapping("SELECT food_code, food_description FROM pref_food", &mappings.Food); err != nil {
		http.Error(w, "Failed to fetch food mappings", http.StatusInternalServerError)
		log.Printf("Error fetching food mappings: %v", err)
		return
	}

	// Fetch hobby mappings
	if err := fetchMapping("SELECT hobby_code, hobby_description FROM pref_hobby", &mappings.Hobby); err != nil {
		http.Error(w, "Failed to fetch hobby mappings", http.StatusInternalServerError)
		log.Printf("Error fetching hobby mappings: %v", err)
		return
	}

	// Fetch music mappings
	if err := fetchMapping("SELECT music_code, music_description FROM pref_music", &mappings.Music); err != nil {
		http.Error(w, "Failed to fetch music mappings", http.StatusInternalServerError)
		log.Printf("Error fetching music mappings: %v", err)
		return
	}

	// Send the mappings as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mappings)
}
