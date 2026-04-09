package main

import (
	"fmt"
	"log"
	"os"
)

// terraform-modules — Reusable Terraform modules for AWS, GCP, and Azure infrastructure
func main() {
	logger := log.New(os.Stdout, "[terraform-modules] ", log.LstdFlags)
	logger.Println("Starting application...")

	if err := run(); err != nil {
		logger.Fatalf("Fatal error: %v", err)
	}
}

func run() error {
	fmt.Println("Application initialized successfully")
	return nil
}
