// backend/main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"  // Needed for more robust routing
	//"github.com/rs/cors"     // Needed for CORS handling
)

// Message struct to represent the data we'll send
type Message struct {
	Text string `json:"text"`
}

func main() {
	router := mux.NewRouter()

	// Define a route that returns a simple message
	router.HandleFunc("/api/message", getMessage).Methods("GET")

	// Serve static files from the frontend directory (after building the Svelte app)
	staticDir := http.Dir("../frontend/dist") // Point to the Svelte build output
	router.PathPrefix("/").Handler(http.FileServer(staticDir))

	// Handle CORS (Cross-Origin Resource Sharing)
	//c := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:8000"}, // Allow requests from the Svelte dev server (adjust port if needed)
	//	AllowCredentials: true,                                                     // Optional: If you need to allow cookies or authorization headers
	//	AllowedMethods:   []string{"GET", "POST", "OPTIONS"},                            // Allowed HTTP methods
	//	Debug:            true,                                                     // Enable debug mode (for development only!)
	//})
	//handler := c.Handler(router)

	port := getPort()
	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// getMessage handles the /api/message endpoint
func getMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the Content-Type header

	message := Message{Text: "Hello from the Go backend!"}

	json.NewEncoder(w).Encode(message)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port if not specified
		log.Printf("Defaulting to port %s", port)
	}
	return port
}