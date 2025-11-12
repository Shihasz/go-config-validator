package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configPath string // Variable to hold the flag value

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates a configuration file against defined policies.",
	Long: `The validate command takes a configuration file path (YAML or JSON)
and checks it for common misconfigurations, security violations, and 
compliance errors.`,
	Run: func(cmd *cobra.Command, args []string) {
		if configPath == "" {
			fmt.Println("Error: Configuration file path must be provided using the -f flag.")
			cmd.Help()
			os.Exit(1)
		}

		// Core logic TODO
		fmt.Printf("Validation request received for file: %s\n", configPath)
		fmt.Println("File read and parsing logic to be implemented.")
	},
}

func init() {
	// Add the validate command to the root command
	rootCmd.AddCommand(validateCmd)

	// Define the mandatory flag for the config file path
	validateCmd.Flags().StringVarP(
		&configPath, // The variable to bind the flag value to
		"file",      // Long flag name: --file
		"f",         // Short flag name: -f
		"",          // Default value
		"Path to the YAML or JSON configuration file to validate (required)",
	)

	// Mark the flag as required
	validateCmd.MarkFlagRequired("file")
}
