package tgs_004

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"testing"
)

// Simplify and sanitize string to be used in Go code
func sanitize(str string) string {
	return strings.ReplaceAll(str, `"`, `\"`)
}

// Escape backticks in strings to avoid syntax errors in Go
func escapeBackticks(str string) string {
	return strings.ReplaceAll(str, "`", "` + \"`\" + `")
}
func TestHelloName(t *testing.T) {
	csvFile, err := os.Open("tgs.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))
	// Assuming the first row is headers
	_, err = reader.Read()
	if err != nil {
		fmt.Println("Error reading CSV header:", err)
		return
	}

	// Open the Go file for writing
	goFile, err := os.Create("definitions.go")
	if err != nil {
		fmt.Println("Error creating definitions.go:", err)
		return
	}
	defer goFile.Close()

	// Write the package declaration and initial part of the map
	_, err = goFile.WriteString(
		`package tgs_004

var TGS_004 = map[string]ParamDefinition{
`)
	if err != nil {
		fmt.Println("Error writing to definitions.go:", err)
		return
	}

	// Process each row of the CSV
	for {
		record, err := reader.Read()
		if err != nil {
			if err == csv.ErrFieldCount || err == csv.ErrBareQuote || err == csv.ErrQuote {
				fmt.Println("Warning: Skipping malformed CSV row:", err)
				continue
			}
			break // End of file or an error occurred
		}

		id := escapeBackticks(record[2])
		description := escapeBackticks(record[4])
		validation := escapeBackticks(record[5])
		isLinkToAnotherTable := strings.ToLower(record[6]) == "true"

		// Write the map entry for each row using backticks for multiline support
		entry := fmt.Sprintf("\t`%s`: {\n\t\tID: `%s`,\n\t\tDescription: `%s`,\n\t\tValidation: `%s`,\n\t\tIsLinkToAnotherTable: %t,\n\t},\n",
			id, id, description, validation, isLinkToAnotherTable)
		_, err = goFile.WriteString(entry)
		if err != nil {
			fmt.Println("Error writing entry to definitions.go:", err)
			return
		}
	}

	// Close the map and file
	_, err = goFile.WriteString("}\n")
	if err != nil {
		fmt.Println("Error finalizing definitions.go:", err)
		return
	}

	fmt.Println("definitions.go generated successfully.")
}
