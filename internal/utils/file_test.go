package utils

import (
	"testing"
)

// Helper function to create a test case struct
type inferenceTest struct {
	path     string
	expected ConfigType
}

func TestInferFileType(t *testing.T) {
	tests := []inferenceTest{
		{"config.yaml", YAML},
		{"config.yml", YAML},
		{"config.JSON", JSON}, // Test case insensitivity
		{"config.tf", Unknown},
		{"noextension", Unknown},
		{"/path/to/my-service.json", JSON},
	}

	for _, test := range tests {
		result := InferFileType(test.path)
		if result != test.expected {
			t.Errorf("InferFileType(%q) FAILED. Expected %v, Got %v", test.path, test.expected.String(), result.String())
		}
	}
}

func TestParseConfigContent(t *testing.T) {
	// 1. YAML Test Case
	yamlContent := []byte(`
service: api-gateway
port: 8080
enabled: true
`)
	// 2. JSON Test Case
	jsonContent := []byte(`{"zone": "us-west-2a", "instance_count": 5}`)

	// 3. Invalid YAML Test Case (Parsing Failure)
	invalidYaml := []byte(`bad-yaml: [`)

	// Run tests
	t.Run("Valid YAML", func(t *testing.T) {
		data, err := ParseConfigContent(yamlContent, YAML)
		if err != nil {
			t.Fatalf("YAML parsing failed unexpectedly: %v", err)
		}
		if len(data) != 3 {
			t.Errorf("Expected 3 keys, got %d", len(data))
		}
	})

	t.Run("Valid JSON", func(t *testing.T) {
		data, err := ParseConfigContent(jsonContent, JSON)
		if err != nil {
			t.Fatalf("JSON parsing failed unexpectedly: %v", err)
		}
		if len(data) != 2 {
			t.Errorf("Expected 2 keys, got %d", len(data))
		}
	})

	t.Run("Parsing Failure", func(t *testing.T) {
		_, err := ParseConfigContent(invalidYaml, YAML)
		if err == nil {
			t.Error("Expected parsing to fail, but it succeeded.")
		}
	})
}
