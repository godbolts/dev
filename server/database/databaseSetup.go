package databaseSetup

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	superUser   = "postgres"
	host        = "localhost"
	database    = "match_me_db"
	newUser     = "kood_user"
	newUserPass = "kood_johvi"
)

var db *sql.DB

// Sets up a new database.
func SetupDatabase() error {

	// Load environment variables from .env file (optional)
	err := godotenv.Load("../server/config.env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Get SUPER_USER_PASS from environment variables
	superUserPass := os.Getenv("SUPER_USER_PASS")
	if superUserPass == "" {
		log.Fatal("SUPER_USER_PASS is not set in environment variables")
	}

	// Get the PORT from the environment variables
	portraw := os.Getenv("PORT")
	if portraw == "" {
		log.Fatal("PORT is not set in environment variables")
	}

	// Convert the PORT from string to int
	port, err := strconv.Atoi(portraw)
	if err != nil {
		log.Fatalf("Invalid PORT value: %s, must be a number", portraw)
	}

	superUserDB, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%d sslmode=disable",
		superUser, superUserPass, host, port))
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}
	defer superUserDB.Close()

	if err := superUserDB.Ping(); err != nil {
		return fmt.Errorf("error pinging the database: %v", err)
	}
	fmt.Println("Connected to PostgreSQL as superuser.")

	var exists bool
	userCheckQuery := `SELECT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'kood_user');`
	err = superUserDB.QueryRow(userCheckQuery).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking if user exists: %v", err)
	}

	if !exists {
		_, err = superUserDB.Exec(`CREATE USER kood_user WITH PASSWORD 'kood_johvi'`)
		if err != nil {
			return fmt.Errorf("error creating user: %v", err)
		}
		fmt.Printf("User '%s' created.\n", newUser)
		_, err = superUserDB.Exec(`ALTER USER kood_user WITH SUPERUSER`)
		if err != nil {
			return fmt.Errorf("error granting superuser privileges: %v", err)
		}
		fmt.Printf("Superuser privileges granted to '%s'.\n", newUser)
	} else {
		fmt.Printf("User '%s' already exists.\n", newUser)
	}

	dbCheckQuery := `SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'match_me_db');`
	err = superUserDB.QueryRow(dbCheckQuery).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking if database exists: %v", err)
	}

	if !exists {
		_, err = superUserDB.Exec(`CREATE DATABASE match_me_db OWNER kood_user`)
		if err != nil {
			return fmt.Errorf("error creating database: %v", err)
		}
		fmt.Printf("Database '%s' created.\n", database)
	} else {
		fmt.Printf("Database '%s' already exists.\n", database)
	}

	_, err = superUserDB.Exec(`GRANT ALL PRIVILEGES ON DATABASE match_me_db TO kood_user`)
	if err != nil {
		return fmt.Errorf("error granting privileges: %v", err)
	}
	fmt.Printf("Privileges granted to '%s' on '%s'.\n", newUser, database)

	newDB, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		newUser, newUserPass, host, port, database))
	if err != nil {
		return fmt.Errorf("error connecting to the new database: %v", err)
	}
	defer newDB.Close()

	if err := newDB.Ping(); err != nil {
		return fmt.Errorf("error pinging the new database: %v", err)
	}
	fmt.Printf("Connected to the new database '%s'.\n", database)

	_, err = newDB.Exec(`CREATE EXTENSION IF NOT EXISTS postgis;`)
	if err != nil {
		return fmt.Errorf("error adding PostGIS extension: %v", err)
	}
	fmt.Println("PostGIS extension added to the database.")

	err = createTables(newDB)
	if err != nil {
		return fmt.Errorf("error creating tables: %v", err)
	}

	err = mapMappingTables(newDB)
	if err != nil {
		return fmt.Errorf("error mapping tables: %v", err)
	}

	fmt.Println("Setup completed.")
	return nil
}

// Sets up SQL tables in a PostgreSQL database.
func createTables(db *sql.DB) error {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS user_table (
			id SERIAL PRIMARY KEY,
			user_uuid UUID UNIQUE,
			password_hash VARCHAR(255) UNIQUE,
			datetime_created TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS user_info (
			id SERIAL PRIMARY KEY,
			user_uuid UUID UNIQUE,
			username VARCHAR(50) UNIQUE,
			email VARCHAR(50),
			first_name VARCHAR(50),
			middle_name VARCHAR(50),
			last_name VARCHAR(50),
			birthdate DATE
		);`,
		`CREATE TABLE IF NOT EXISTS user_data (
			id SERIAL PRIMARY KEY,
			user_uuid UUID UNIQUE,
			user_city VARCHAR(50),
			register_location GEOGRAPHY(POINT, 4326), 
			browser_location GEOGRAPHY(POINT, 4326)
		);`,
		`CREATE TABLE IF NOT EXISTS sessions (
			id SERIAL PRIMARY KEY,
			session_guid VARCHAR(8) UNIQUE,
			user_uuid UUID UNIQUE,
			email VARCHAR(100)
		);`,
		`CREATE TABLE IF NOT EXISTS profile_info (
			id SERIAL PRIMARY KEY,
			user_uuid UUID UNIQUE,
			about_me VARCHAR(1000),
			food_myvariabledata VARCHAR(1000),
			hobbies_myvariabledata VARCHAR(1000),
			music_myvariabledata VARCHAR(1000)
		);`,
		`CREATE TABLE IF NOT EXISTS pref_food (
			id SERIAL PRIMARY KEY,
			food_code VARCHAR(2) UNIQUE,
			food_description VARCHAR(50)
		);`,
		`CREATE TABLE IF NOT EXISTS pref_hobby (
			id SERIAL PRIMARY KEY,
			hobby_code VARCHAR(2) UNIQUE,
			hobby_description VARCHAR(50)
		);`,
		`CREATE TABLE IF NOT EXISTS pref_music (
			id SERIAL PRIMARY KEY,
			music_code VARCHAR(2) UNIQUE,
			music_description VARCHAR(50)
		);`,
		`CREATE TABLE IF NOT EXISTS weights (
			id SERIAL PRIMARY KEY,
			user_uuid UUID UNIQUE,
			weigh_distance NUMERIC,
			weigh_age NUMERIC,
			weigh_food NUMERIC,
			weigh_hobbies NUMERIC,
			weigh_music NUMERIC
		);`,
		`CREATE TABLE IF NOT EXISTS pending_connections (
			user_uuid_of UUID,
			user_uuid_with UUID
		);`,
		`CREATE TABLE IF NOT EXISTS real_connections (
			user_uuid_of UUID,
			user_uuid_with UUID
		);`,
		`CREATE TABLE IF NOT EXISTS reccomendations (
			user_uuid_of UUID,
			user_uuid_with UUID,
			compability NUMERIC,
			distance NUMERIC
		);`,
	}

	for _, query := range tables {
		if _, err := db.Exec(query); err != nil {
			return fmt.Errorf("error creating table: %v", err)
		}
	}

	fmt.Println("All tables created.")
	return nil
}

func mapMappingTablesFunctionInternal(db *sql.DB) error {
	queries := []string{
		`INSERT INTO pref_food (food_code, food_description) VALUES 
			('A1', 'Carnivore Diet'),
			('A2', 'Vegetarian'),
			('A3', 'Vegan'),
			('B1', 'Keto Diet'),
			('B2', 'Paleo Diet'),
			('C1', 'Mediterranean Diet'),
			('C2', 'Low Carb'),
			('D1', 'High Protein'),
			('E1', 'Balanced Diet'),
			('F1', 'Gluten-Free'),
			('G1', 'Lactose-Free');`,
		`INSERT INTO pref_hobby (hobby_code, hobby_description) VALUES 
			('A1', 'Reading'),
			('A2', 'Writing'),
			('B1', 'Painting'),
			('B2', 'Photography'),
			('C1', 'Gaming'),
			('C2', 'Gardening'),
			('D1', 'Cooking'),
			('D2', 'Baking'),
			('E1', 'Traveling'),
			('F1', 'Fishing'),
			('G1', 'Hiking');`,
		`INSERT INTO pref_music (music_code, music_description) VALUES 
			('A1', 'Rock'),
			('A2', 'Pop'),
			('B1', 'Classical'),
			('B2', 'Jazz'),
			('C1', 'Hip-Hop'),
			('C2', 'Country'),
			('D1', 'Electronic'),
			('D2', 'Reggae'),
			('E1', 'Blues'),
			('F1', 'Folk'),
			('G1', 'Metal');`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Printf("Error executing query: %v", err)
		}
	}
	fmt.Println("All mapps mapped created.")
	return nil
}

func checkIfTableIsEmpty(db *sql.DB, tableName string) (bool, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking table %s: %v", tableName, err)
	}
	return count == 0, nil // Return true if the table is empty
}

func mapMappingTables(db *sql.DB) error {
	// Check if the tables are empty
	isFoodEmpty, err := checkIfTableIsEmpty(db, "pref_food")
	if err != nil {
		return fmt.Errorf("error checking pref_food table: %v", err)
	}
	isHobbyEmpty, err := checkIfTableIsEmpty(db, "pref_hobby")
	if err != nil {
		return fmt.Errorf("error checking pref_hobby table: %v", err)
	}
	isMusicEmpty, err := checkIfTableIsEmpty(db, "pref_music")
	if err != nil {
		return fmt.Errorf("error checking pref_music table: %v", err)
	}

	// Only call mapMappingTables if any of the tables are empty
	if isFoodEmpty || isHobbyEmpty || isMusicEmpty {
		// Proceed with mapping the tables
		err = mapMappingTablesFunctionInternal(db)
		if err != nil {
			return fmt.Errorf("error mapping tables: %v", err)
		}
	} else {
		// Optionally log that the tables are not empty and mapping is skipped
		log.Println("Tables are not empty, skipping mapping.")
	}

	return nil
}

func InitDB() error {

	var err error
	db, err = sql.Open("postgres", "user=kood_user password=kood_johvi dbname=match_me_db host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Printf("Error Opening database: %v", err)
		return err
	}

	// Ensure the database connection is available
	if err := db.Ping(); err != nil {
		log.Printf("Error Pinging database: %v", err)
		return err
	}

	return nil
}

func GetDB() *sql.DB {
	if db == nil {
		InitDB()
	}
	return db
}
