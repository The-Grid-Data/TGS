package tgs_004

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTGS_004(t *testing.T) {
	indexFileURL := "https://github.com/The-Grid-Data/TGS/blob/main/tgs-004/doc/index.md"
	allParameters, err := FetchAllParameters(indexFileURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonData, err := json.MarshalIndent(allParameters, "", "    ") // Indent for pretty print
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}
