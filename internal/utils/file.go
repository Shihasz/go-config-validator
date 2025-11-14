package utils

import (
	"fmt"
	"os"
)

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
