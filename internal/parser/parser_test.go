package parser

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// TestMissingFileError confirms that an error is raised when a file is missing
func TestMissingFileError(t *testing.T) {
	inferfilePath := filepath.Join("tests", "Inferfile.missing_file_path")

	// Parse the Inferfile and expect an error
	_, err := ParseInferfile(inferfilePath)
	if err == nil {
		t.Fatalf("Expected an error for missing file, but got none")
	}
	expectedErrMsg := "file not found: ./this/file/does/not/exist (Inferfile line: 1)"
	if err.Error() != expectedErrMsg {
		t.Fatalf("Expected error message to be: %s, but got: %s", expectedErrMsg, err.Error())
	}
}

// TestParseInferfileWithStubFile tests parsing an Inferfile with a stub file
func TestParseInferfileWithStubFile(t *testing.T) {
	// Create a temporary stub Go file.
	stubGoCode := `package main

// Infer: OpenAI client
func openaiClient() {
	// This is a stub function
}
// EndInfer: OpenAI client
`
	tmpGoFilePath := "/tmp/mycode.go"
	err := ioutil.WriteFile(tmpGoFilePath, []byte(stubGoCode), 0644)
	if err != nil {
		t.Fatalf("Failed to create temporary Go file: %s", err)
	}
	defer os.Remove(tmpGoFilePath) // Clean up after the test

	// Assume that Inferfile.valid is correctly pointing to /tmp/mycode.go
	inferfilePath := filepath.Join("tests", "Inferfile.valid")

	var config *InferConfiguration
	// Parse the Inferfile and expect no error
	config, err = ParseInferfile(inferfilePath)
	if err != nil {
		t.Fatalf("Expected no error, but got: %s", err)
	} else {
		// Expect the data to have a tag "OpenAI client"
		if config.Files[0].Tags[0].Code == "" {
			t.Fatalf("Expected the data to have a tag 'OpenAI client', but it did not")
		}
	}
}

// TestParseInferfileNoTags confirms that an Inferfile with a file resource but no tags can be loaded without errors.
func TestParseInferfileNoTagsWithStubFile(t *testing.T) {
	// Create a temporary stub Go file.
	stubGoCode := `package main

// Infer: example_tag
func example() {
	// Some example code
}
// EndInfer: example_tag
`
	tmpGoFilePath := "/tmp/mycode.go"
	err := ioutil.WriteFile(tmpGoFilePath, []byte(stubGoCode), 0644)
	if err != nil {
		t.Fatalf("Failed to create temporary Go file: %s", err)
	}
	defer os.Remove(tmpGoFilePath) // Clean up after the test

	inferfilePath := filepath.Join("tests", "Inferfile.no_tags")

	// Parse the Inferfile and expect no error
	config, err := ParseInferfile(inferfilePath)
	if err != nil {
		t.Fatalf("Expected no error, but got: %s", err)
	}

	// Additionally, you can check if the file paths are parsed correctly.
	if len(config.Files) != 1 {
		t.Errorf("Expected 1 file resource, got %d", len(config.Files))
	}

	// You can also check if the file path is as expected.
	expectedPath := "/tmp/mycode.go"
	if config.Files[0].Path != expectedPath {
		t.Errorf("Expected file path to be '%s', got '%s'", expectedPath, config.Files[0].Path)
	}

	// Check that there are no tags in the file resource.
	if len(config.Files[0].Tags) != 0 {
		t.Errorf("Expected no tags, got %d", len(config.Files[0].Tags))
	}
}
