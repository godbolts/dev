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
		Remove bool   `json:"remove"`
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
	defer db.Close()

	// Query the current interests_myvariabledata for the user
	var currentData string
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

	// Split the current data into a slice of codes
	codes := strings.Split(currentData, ",")
	updatedCodes := []string{}

	if requestBody.Remove {
		// Remove the code if it exists
		codeFound := false
		for _, existingCode := range codes {
			if existingCode == code {
				codeFound = true
				continue // Skip the code to remove it
			}
			updatedCodes = append(updatedCodes, existingCode)
		}

		if !codeFound {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Code not found, nothing to remove"))
			return
		}
	} else {
		// Add the code if it doesn't already exist
		codeExists := false
		for _, existingCode := range codes {
			if existingCode == code {
				codeExists = true
				break
			}
		}

		if codeExists {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Code already exists"))
			return
		}

		updatedCodes = append(codes, code)
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
		Remove bool   `json:"remove"`
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
	defer db.Close()

	// Query the current interests_myvariabledata for the user
	var currentData string
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

	// Split the current data into a slice of codes
	codes := strings.Split(currentData, ",")
	updatedCodes := []string{}

	if requestBody.Remove {
		// Remove the code if it exists
		codeFound := false
		for _, existingCode := range codes {
			if existingCode == code {
				codeFound = true
				continue // Skip the code to remove it
			}
			updatedCodes = append(updatedCodes, existingCode)
		}

		if !codeFound {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Code not found, nothing to remove"))
			return
		}
	} else {
		// Add the code if it doesn't already exist
		codeExists := false
		for _, existingCode := range codes {
			if existingCode == code {
				codeExists = true
				break
			}
		}

		if codeExists {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Code already exists"))
			return
		}

		updatedCodes = append(codes, code)
	}

	// Join the updated codes into a string
	newData := strings.Join(updatedCodes, ",")

	// Update the database
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
		Remove bool   `json:"remove"`
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
	defer db.Close()

	// Query the current interests_myvariabledata for the user
	var currentData string
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

	// Split the current data into a slice of codes
	codes := strings.Split(currentData, ",")
	updatedCodes := []string{}

	if requestBody.Remove {
		// Remove the code if it exists
		codeFound := false
		for _, existingCode := range codes {
			if existingCode == code {
				codeFound = true
				continue // Skip the code to remove it
			}
			updatedCodes = append(updatedCodes, existingCode)
		}

		if !codeFound {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Code not found, nothing to remove"))
			return
		}
	} else {
		// Add the code if it doesn't already exist
		codeExists := false
		for _, existingCode := range codes {
			if existingCode == code {
				codeExists = true
				break
			}
		}

		if codeExists {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Code already exists"))
			return
		}

		updatedCodes = append(codes, code)
	}

	// Join the updated codes into a string
	newData := strings.Join(updatedCodes, ",")

	// Update the database
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
	defer db.Close()

	// Query the current preferences for the user
	var preferences struct {
		InterestsMyVariableData string `json:"interests_myvariabledata"`
		MusicMyVariableData     string `json:"music_myvariabledata"`
		FoodMyVariableData      string `json:"food_myvariabledata"`
	}

	query := `
        SELECT interests_myvariabledata, music_myvariabledata, food_myvariabledata 
        FROM profile_info 
        WHERE user_uuid = $1
    `
	err = db.QueryRow(query, userID).Scan(
		&preferences.InterestsMyVariableData,
		&preferences.MusicMyVariableData,
		&preferences.FoodMyVariableData,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database query error", http.StatusInternalServerError)
			log.Printf("Error querying database: %v", err)
		}
		return
	}

	// Send the preferences as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(preferences)
}

func PrefMappingGet(w http.ResponseWriter, r *http.Request) {
	// Connect to the database
	db := databaseSetup.GetDB() // Assume GetDB() returns *sql.DB
	defer db.Close()

	// Helper function to fetch mapping data
	fetchMapping := func(query string, target *[]structures.PreferenceMapping) error {
		rows, err := db.Query(query)
		if err != nil {
			return err
		}
		defer rows.Close()
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
