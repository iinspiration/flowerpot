package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // Import your database driver package here
	"github.com/joho/godotenv"
)

// Member represents a member in the database
type Member struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Add more fields as needed
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", connString)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the GET /members endpoint handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Set the response headers and write the JSON data
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "API is running"}`))
	})

	// Create the GET /members endpoint handler
	http.HandleFunc("/members", func(w http.ResponseWriter, r *http.Request) {
		// Query the database for members
		rows, err := db.Query("SELECT id, name FROM members")
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Iterate over the rows and build the response
		var members []Member
		for rows.Next() {
			var member Member
			err := rows.Scan(&member.ID, &member.Name)
			if err != nil {
				log.Println(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			members = append(members, member)
		}

		// Convert the members slice to JSON
		jsonData, err := json.Marshal(members)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set the response headers and write the JSON data
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	})

	// Start the HTTP server
	log.Println("Server listening on :8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
