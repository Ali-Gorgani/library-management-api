package main

import (
	"os"
	"os/exec"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	// Set up the logger to write to the console
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	// Define the root command
	rootCmd := &cobra.Command{
		Use:   "service",
		Short: "Service CLI for running different services",
	}

	// Add subcommands for each service
	rootCmd.AddCommand(apiGatewayCmd)
	rootCmd.AddCommand(authServiceCmd)
	rootCmd.AddCommand(usersServiceCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Failed to execute command")
	}
}

// Define the API Gateway command
var apiGatewayCmd = &cobra.Command{
	Use:   "api-gateway",
	Short: "Run the API Gateway",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Starting API Gateway...")
		runService("./api-gateway/cmd/main.go")
	},
}

// Define the Auth Service command
var authServiceCmd = &cobra.Command{
	Use:   "auth-service",
	Short: "Run the Auth Service",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Starting Auth Service...")
		runService("./auth-service/cmd/main.go")
	},
}

// Define the Users Service command
var usersServiceCmd = &cobra.Command{
	Use:   "users-service",
	Short: "Run the Users Service",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Starting Users Service...")
		runService("./users-service/cmd/main.go")
	},
}

// Helper function to run a Go service
func runService(path string) {
	cmd := exec.Command("go", "run", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command and check for errors
	if err := cmd.Run(); err != nil {
		log.Error().Err(err).Msg("Error running service")
	}
}
