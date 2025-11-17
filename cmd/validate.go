package cmd

import (
	"fmt"
	"os"

	"github.com/Shihasz/go-config-validator/internal/utils"
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

		fmt.Printf("Starting validation for file: %s\n", configPath)

		// Step 1: Check file existence and access using utility function
		if err := utils.CheckFileAccess(configPath); err != nil {
			fmt.Printf("Validation Failed: %v\n", err)
			os.Exit(1)
		}

		// Step 2: Read the file contents using the highly efficient os.ReadFile
		content, err := os.ReadFile(configPath)
		if err != nil {
			fmt.Printf("Validation Failed: Could not read file content: %v\n", err)
			os.Exit(1)
		}

		// Step 2: Infer the file type
		configType := utils.InferFileType(configPath)
		if configType == utils.Unknown {
			fmt.Printf("Validation Failed: Unsupported file extension for path: %s. Must be .yaml, .yml, or .json.\n", configPath)
			os.Exit(1)
		}
		fmt.Printf("Inferred configuration type: **%s**\n", configType.String())

		// Step 3: Parse the content
		parsedData, err := utils.ParseConfigContent(content, configType)
		if err != nil {
			fmt.Printf("Validation Failed: Parsing Error: %v\n", err)
			os.Exit(1)
		}

		// Step 4: Success and next steps
		fmt.Printf("Successfully parsed configuration file with **%d** top-level keys.\n", len(parsedData))
		fmt.Println("Next Step: Implementing custom validation rules.")
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
