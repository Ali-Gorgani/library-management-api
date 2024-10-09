package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "service",
		Short: "Service CLI for running different services",
	}

	// Add subcommands for each service
	rootCmd.AddCommand(apiGatewayCmd)
	rootCmd.AddCommand(authServiceCmd)
	rootCmd.AddCommand(usersServiceCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var apiGatewayCmd = &cobra.Command{
	Use:   "api-gateway",
	Short: "Run the API Gateway",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting API Gateway...")
		runService("./api-gateway/cmd/main.go")
	},
}

var authServiceCmd = &cobra.Command{
	Use:   "auth-service",
	Short: "Run the Auth Service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Auth Service...")
		runService("./auth-service/cmd/main.go")
	},
}

var usersServiceCmd = &cobra.Command{
	Use:   "users-service",
	Short: "Run the Users Service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Users Service...")
		runService("./users-service/cmd/main.go")
	},
}

// Helper function to run a Go service
func runService(path string) {
	cmd := exec.Command("go", "run", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running service:", err)
	}
}
