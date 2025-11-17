package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
