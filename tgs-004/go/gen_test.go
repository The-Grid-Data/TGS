package tgs_004

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestUploadCountries(t *testing.T) {
	// Open the CSV file
	csvFile, err := os.Open("countries.csv")
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer csvFile.Close()

	// Read CSV data into a slice of records
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse CSV file: %v", err)
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbConfig.user, dbConfig.password, dbConfig.host, dbConfig.port, dbConfig.database)

	// Connect to the MySQL database
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Prepare SQL statement for inserting data
	stmt, err := db.Prepare("INSERT INTO countries (Name, Code) VALUES (?, ?)")
	if err != nil {
		log.Fatalf("Failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	// Skip the header row and iterate over the records
	for _, record := range records[1:] {
		_, err = stmt.Exec(record[0], record[1])
		if err != nil {
			log.Printf("Failed to execute statement: %v", err)
		}
	}

	fmt.Println("Data uploaded successfully")
}
