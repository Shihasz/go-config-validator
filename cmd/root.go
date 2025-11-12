package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "config-validator",
	Short: "A powerful CLI tool to validate infrastructure configuration files (YAML/JSON).",
	Long: `config-validator is a professional CLI utility written in Go
to ensure configuration quality, compliance, and correctness 
before deployment to staging or production environments.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default action when no subcommand is given (e.g., just printing help)
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
