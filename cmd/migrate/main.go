package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func main() {
	mode := flag.String("mode", "", "Mode of operation: 'export' or 'import'")
	project := flag.String("gcloud_project_id", "", "Google Cloud Project ID")
	file := flag.String("file", "users.json", "Path to the JSON file for export/import")

	flag.Parse()

	if *mode == "" || *project == "" {
		fmt.Println("Usage: go run cmd/migrate/main.go -mode [export|import] -gcloud_project_id [GCLOUD_PROJECT_ID] -file [FILENAME]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	ctx := context.Background()

	switch *mode {
	case "export":
		runExport(ctx, *project, *file)
	case "import":
		runImport(ctx, *project, *file)
	default:
		log.Fatalf("Unknown mode: %s. Use 'export' or 'import'", *mode)
	}
}

func runExport(ctx context.Context, project, filename string) {
	log.Printf("Starting export from project: %s", project)

	client := utils.OpenClient(&ctx, project)
	if client == nil {
		log.Fatalf("Failed to create datastore client")
	}
	defer client.Close()

	var users []def.UserData
	query := datastore.NewQuery("User")
	_, err := client.GetAll(ctx, query, &users)
	if err != nil {
		log.Fatalf("Failed to query users: %v", err)
	}

	if len(users) == 0 {
		log.Println("No users found.")
		return
	}

	log.Printf("Retrieved %d users. Saving to %s...", len(users), filename)

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal users to JSON: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatalf("Failed to write to file %s: %v", filename, err)
	}

	log.Printf("Export completed successfully. Data saved to %s", filename)
}

func runImport(ctx context.Context, project, filename string) {
	log.Printf("Starting import to project: %s from file: %s", project, filename)

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filename, err)
	}

	var users []def.UserData
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	log.Printf("Found %d users in file. Starting upload...", len(users))

	client := utils.OpenClient(&ctx, project)
	if client == nil {
		log.Fatalf("Failed to create datastore client")
	}
	defer client.Close()

	successCount := 0
	failCount := 0

	for _, user := range users {
		if user.Id == "" {
			log.Printf("Skipping user with empty ID: %v", user)
			failCount++
			continue
		}

		key := datastore.NameKey("User", user.Id, nil)
		_, err := client.Put(ctx, key, &user)
		if err != nil {
			log.Printf("Failed to put user (ID: %s): %v", user.Id, err)
			failCount++
		} else {
			successCount++
		}
	}

	log.Printf("Import completed. Successfully imported: %d, Failed: %d", successCount, failCount)
}
