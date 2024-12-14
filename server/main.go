package main

import (
	"log"
	databaseSetup "match_me_module/database"
	"match_me_module/routes"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Open or create a log file
	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Set log output to the file
	log.SetOutput(logFile)

	// Log server start
	log.Println("Server is starting")

	// Initialize the database
	if err := databaseSetup.SetupDatabase(); err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}

	if err := databaseSetup.InitDB(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Create the router
	r := mux.NewRouter()

	// Define API routes
	r.HandleFunc("/api/user", routes.UserInfo).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/login", routes.Login).Methods("POST")
	r.HandleFunc("/api/register", routes.Register).Methods("POST")
	r.HandleFunc("/api/edit/user", routes.EditUsername).Methods("POST")
	r.HandleFunc("/api/edit/email", routes.EditEmail).Methods("POST")
	r.HandleFunc("/api/edit/first", routes.EditFirst).Methods("POST")
	r.HandleFunc("/api/edit/middle", routes.EditMiddle).Methods("POST")
	r.HandleFunc("/api/edit/last", routes.EditLast).Methods("POST")
	r.HandleFunc("/api/edit/pass", routes.EditPassword).Methods("POST")
	r.HandleFunc("/api/edit/city", routes.EditCity).Methods("POST")
	r.HandleFunc("/api/biog/about", routes.AboutYou).Methods("POST")
	r.HandleFunc("/api/biog/birthday", routes.Birthday).Methods("POST")
	r.HandleFunc("/api/biog/aboutget", routes.AboutYouGet).Methods("GET")
	r.HandleFunc("/api/biog/birthdayget", routes.BirthdayGet).Methods("GET")
	r.HandleFunc("/api/pref/food", routes.FoodPref).Methods("POST")
	r.HandleFunc("/api/pref/hobby", routes.HobbyPref).Methods("POST")
	r.HandleFunc("/api/pref/music", routes.MusicPref).Methods("POST")
	r.HandleFunc("/api/pref/get", routes.PrefGet).Methods("GET")
	r.HandleFunc("/api/pref/mapget", routes.PrefMappingGet).Methods("GET")
	r.HandleFunc("/api/wigh/dist", routes.WeightDistance).Methods("POST")
	r.HandleFunc("/api/wigh/age", routes.WeightAge).Methods("POST")
	r.HandleFunc("/api/wigh/food", routes.WeightFood).Methods("POST")
	r.HandleFunc("/api/wigh/hobby", routes.WeightHobbies).Methods("POST")
	r.HandleFunc("/api/wigh/music", routes.WeightMusic).Methods("POST")
	r.HandleFunc("/api/wigh/get", routes.WeightGet).Methods("GET")

	// Set up CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow React frontend
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},        // Include OPTIONS for preflight
		AllowedHeaders:   []string{"Authorization", "Content-Type"}, // Headers expected by the client
	}).Handler(r)

	// Define server port
	port := "3001"
	log.Println("Server is running on http://localhost:" + port)

	// Start the HTTP server
	if err := http.ListenAndServe(":"+port, corsHandler); err != nil {
		log.Fatal(err)
	}
}
