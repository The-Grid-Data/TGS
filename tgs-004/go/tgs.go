package tgs_004

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type ParameterData struct {
	ParameterID     string `json:"Parameter ID"`
	Table           string `json:"Table"`
	Name            string `json:"Name"`
	Type            string `json:"Type"`
	Description     string `json:"Description"`
	ValidationSteps string `json:"Validation steps"`
	InDBD           string `json:"In DBD"`
	Notes           string `json:"Notes"`
}

// FetchParameterData retrieves parameter data by ID, providing detailed errors.
func FetchParameterData(id, indexFileURL string) (*ParameterData, error) {
	// Generate the base URL for the repository to fetch raw Markdown files
	repoBaseURL := strings.Join(strings.Split(indexFileURL, "/")[:len(strings.Split(indexFileURL, "/"))-1], "/")
	repoBaseURL = strings.ReplaceAll(strings.ReplaceAll(repoBaseURL, "blob/", ""), "github.com", "raw.githubusercontent.com")

	// Fetch the Markdown content by ID
	def, err := FetchMarkdownFileByID(repoBaseURL, id)
	if err != nil {
		// Wrap the error with detailed context
		return nil, fmt.Errorf("failed to fetch markdown file by ID '%s' from URL '%s': %w", id, repoBaseURL, err)
	}

	// Convert the fetched markdown map to a ParameterData struct
	paramData, err := MapToParameterData(def)
	if err != nil {
		// Wrap the error with detailed context
		return nil, fmt.Errorf("failed to map fetched data to ParameterData struct for ID '%s': %w", id, err)
	}

	return paramData, nil
}

// MapToParameterData converts a map of strings to a ParameterData struct and checks for required fields.
func MapToParameterData(data map[string]string) (*ParameterData, error) {
	// Check for mandatory fields
	requiredFields := []string{
		"Parameter ID",
		"Table",
		"Name",
		"Type",
		"Description",
		//"Validation steps",
		//"In DBD",
		//"Notes",
	}
	for _, field := range requiredFields {
		if value, ok := data[field]; !ok || value == "" {
			return nil, fmt.Errorf("missing or empty required field: %s", field)
		}
	}

	// Create a ParameterData instance with values from the map
	return &ParameterData{
		ParameterID:     data["Parameter ID"],
		Table:           data["Table"],
		Name:            data["Name"],
		Type:            data["Type"],
		Description:     data["Description"],
		ValidationSteps: data["Validation steps"],
		InDBD:           data["In DBD"],
		Notes:           data["Notes"],
	}, nil
}

// FetchFileContent makes an HTTP GET request to the specified URL and returns the response body as a string.
// It provides detailed errors on failure, helping with debugging and error tracking.
func FetchFileContent(url string) (string, error) {

	if READ_TGS_PARAMS_LOCALLY {

		discard := "https://raw.githubusercontent.com/The-Grid-Data/TGS/main/" // tgs-004/doc/index.md
		postfix := strings.TrimPrefix(url, discard)                            // tgs-004/doc/index.md

		path := "../../" + postfix
		return FetchFileContentLocally(path)
	}

	resp, err := http.Get(url)
	if err != nil {
		// Include the URL in the error message for clarity on what was being requested
		return "", fmt.Errorf("failed to make HTTP GET request to '%s': %v", url, err)
	}
	defer resp.Body.Close()

	// Check the HTTP status code for success (HTTP 200 OK)
	if resp.StatusCode != 200 {
		// Include the URL and the status code to help diagnose issues with the request or endpoint
		return "", fmt.Errorf("HTTP request to '%s' returned non-200 status code: %d", url, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Provide details about the part of the process where the error occurred
		return "", fmt.Errorf("failed to read response body from '%s': %v", url, err)
	}

	return string(body), nil
}

func FetchFileContentLocally(filePath string) (string, error) {
	// Read the contents of the file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		// Provide details about the error that occurred while reading the file
		return "", fmt.Errorf("failed to read file '%s': %v", filePath, err)
	}

	// Convert byte slice to string
	content := string(data)

	return content, nil
}

func ExtractParameterIDs(indexFileURL string) ([]string, error) {
	rawIndexURL := strings.ReplaceAll(strings.ReplaceAll(indexFileURL, "github.com", "raw.githubusercontent.com"), "blob/", "")
	indexContent, err := FetchFileContent(rawIndexURL)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`\((.*?).md\)`)
	matches := re.FindAllStringSubmatch(indexContent, -1)
	var parameterIDs []string
	for _, match := range matches {
		parameterIDs = append(parameterIDs, match[1])
	}
	return parameterIDs, nil
}
func FetchMarkdownFileByID(repoBaseURL, parameterID string) (map[string]string, error) {
	mdFileURL := fmt.Sprintf("%s/%s.md", repoBaseURL, parameterID)
	mdContent, err := FetchFileContent(mdFileURL)
	if err != nil {
		return nil, err
	}

	dataDict := make(map[string]string)
	//dataDict["Parameter ID"] = parameterID

	// Adjust the regex pattern to match lines with possible multiline values enclosed in triple backticks
	re := regexp.MustCompile("^([^:]+?): ```\\s*([\\s\\S]*?)\\s*```$")
	lines := strings.Split(mdContent, "\n")
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			key := strings.TrimSpace(matches[1])
			value := strings.TrimSpace(matches[2])
			dataDict[key] = value
		}
	}
	return dataDict, nil
}

var READ_TGS_PARAMS_LOCALLY = false

func FetchAllParameters(indexFileURL string) (map[string]*ParameterData, error) {
	repoBaseURL := strings.Join(strings.Split(indexFileURL, "/")[:len(strings.Split(indexFileURL, "/"))-1], "/")
	repoBaseURL = strings.ReplaceAll(strings.ReplaceAll(repoBaseURL, "blob/", ""), "github.com", "raw.githubusercontent.com")
	parameterIDs, err := ExtractParameterIDs(indexFileURL)
	if err != nil {
		return nil, err
	}

	allParameters := make(map[string]*ParameterData)
	for _, parameterID := range parameterIDs {
		paramData, err := FetchParameterData(parameterID, indexFileURL)
		if err != nil {
			return nil, err
		}
		allParameters[parameterID] = paramData
	}
	return allParameters, nil
}
