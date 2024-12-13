package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run ./Migration/migrate.go [up|down|steps|force] [optional step/count/version]")
	}

	action := os.Args[1]

	// Get the absolute path to the sql directory
	absPath, err := filepath.Abs("Migration/sql")
	if err != nil {
		log.Fatalf("Failed to resolve absolute path: %v", err)
	}

	// Convert Windows backslashes to forward slashes
	absPath = strings.ReplaceAll(absPath, "\\", "/")

	log.Printf("Using migration path: %s", absPath)

	m, err := migrate.New(
		"file://"+absPath, // Correctly formatted path with "file://" scheme
		"postgres://postgres:Cel-365.@localhost:5432/postgres?sslmode=disable",
	)

	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	switch action {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
		log.Println("Migrations applied successfully")
	case "down":
		if err := m.Steps(-1); err != nil {
			log.Fatalf("Failed to rollback migration: %v", err)
		}
		log.Println("Rolled back last migration")
	case "steps":
		if len(os.Args) < 3 {
			log.Fatalf("Please provide the number of steps to migrate (positive for up, negative for down)")
		}
		steps, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid steps value: %v", err)
		}
		if err := m.Steps(steps); err != nil {
			log.Fatalf("Failed to apply steps: %v", err)
		}
		log.Printf("Applied %d step(s)", steps)
	case "force":
		if len(os.Args) < 3 {
			log.Fatalf("Please provide the version to force")
		}
		version, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid version value: %v", err)
		}
		if err := m.Force(version); err != nil {
			log.Fatalf("Failed to force version: %v", err)
		}
		log.Printf("Forced migration to version %d", version)
	default:
		log.Fatalf("Invalid action. Usage: go run ./Migration/migrate.go [up|down|steps|force]")
	}
}
