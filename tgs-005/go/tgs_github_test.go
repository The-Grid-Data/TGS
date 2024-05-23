package tgs_005

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_TGS_005_fetch(t *testing.T) {
	indexFileURL := "https://github.com/The-Grid-Data/TGS/blob/main/tgs-005/doc/index.md"

	var err error

	//t.Logf("Fetching %v", indexFileURL)
	//err = test(indexFileURL)
	//
	//if err != nil {
	//	t.Fatalf("Fetch failed: %v", err)
	//}

	READ_TGS_PARAMS_LOCALLY = true
	defer func() {
		READ_TGS_PARAMS_LOCALLY = false
	}()

	t.Logf("Testing locally")
	err = test(indexFileURL)
	if err != nil {
		t.Fatalf("Parsing failed: %v", err)
	}
}

func test(indexFileURL string) error {
	allParameters, err := FetchAllParameters(indexFileURL)
	if err != nil {
		return err
	}
	jsonData, err := json.MarshalIndent(allParameters, "", "    ") // Indent for pretty print
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData))

	return nil
}
