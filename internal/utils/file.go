package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// ConfigType represents the supported configuration file types.
type ConfigType int

const (
	Unknown ConfigType = iota // 0
	YAML                      // 1
	JSON                      // 2
)

// InferFileType checks the file extension and returns the corresponding ConfigType.
func InferFileType(path string) ConfigType {
	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".yaml", ".yml":
		return YAML
	case ".json":
		return JSON
	default:
		return Unknown
	}
}

// String returns the string representation of ConfigType.
func (ct ConfigType) String() string {
	switch ct {
	case YAML:
		return "YAML"
	case JSON:
		return "JSON"
	default:
		return "UNKNOWN"
	}
}

// CheckFileAccess verifies if a file exists and is accessible.
// It returns an error if the file does not exist or permissions are wrong.
func CheckFileAccess(path string) error {
	// 1. Check if the file exists and is accessible (read permission check is implicit here)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("file not found at path: %s", path)
	}
	if err != nil {
		// Handles other errors like permission denied
		return fmt.Errorf("error accessing file %s: %w", path, err)
	}

	return nil
}

// ParseConfigContent takes raw byte content and the file type,
// then unmarshals it into a generic map.
func ParseConfigContent(content []byte, configType ConfigType) (map[string]interface{}, error) {
	var data map[string]interface{}

	switch configType {
	case YAML:
		// Use yaml.v3 to unmarshal YAML content
		if err := yaml.Unmarshal(content, &data); err != nil {
			return nil, fmt.Errorf("failed to unmarshal YAML content: %w", err)
		}
	case JSON:
		// Use standard library for JSON content
		if err := json.Unmarshal(content, &data); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON content: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported configuration type: %s", configType.String())
	}

	// Return the parsed data map
	return data, nil
}
