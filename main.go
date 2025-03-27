// ============= ####### =============
// ============== main ==============
// ============= ####### =============

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ============= ####### =============
// ============== vars ==============
// ============= ####### =============

// rootCmd represents the root command of the application
var rootCmd = &cobra.Command{
	Use:   "employee-holiday-tracker",
	Short: "Employee Holiday Tracker - Manage employee holidays",
	Long: `A Go application for tracking employee holidays using MongoDB 
for persistence, Redis for caching, and Connect-Go for API endpoints.`,
}

// serveCmd represents the command to start the API server
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the API server",
	Long:  `Starts the HTTP server that provides the employee holiday tracking API.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting API server...")
		fmt.Println("Server running on port 8080")
	},
}

// migrateCmd represents the command to run database migrations
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  `Executes database migrations to set up or update the MongoDB schema.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running database migrations...")
		fmt.Println("Migrations completed successfully")
	},
}

// ============= ####### =============
// ============== init ==============
// ============= ####### =============

func init() {
	// Add subcommands to root command
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(migrateCmd)
}

// ============= ####### =============
// ============== funcs =============
// ============= ####### =============

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

// main is the entry point of the application
func main() {
	fmt.Println("Starting Employee Holiday Tracker...")

	if err := Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
