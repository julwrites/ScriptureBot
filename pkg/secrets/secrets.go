package secrets

import (
	"context"
	"fmt"
	"log"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/joho/godotenv"
)

func init() {
	LoadAndLog()
}

// LoadAndLog loads environment variables from a .env file (if present) and logs
// the status of the GCLOUD_PROJECT_ID. This function is called automatically on package initialization.
// It is also exported to allow for re-loading in test environments.
func LoadAndLog() {
	// godotenv.Overload will read your .env file and set the environment variables.
	// It will OVERWRITE any existing environment variables.
	err := godotenv.Overload()
	if err != nil {
		log.Println("No .env file found, continuing with environment variables")
	}

	// Log the status of the GCLOUD_PROJECT_ID for debugging purposes.
	if projectID, ok := os.LookupEnv("GCLOUD_PROJECT_ID"); ok {
		log.Printf("GCLOUD_PROJECT_ID is set: %s", projectID)
	} else {
		log.Println("GCLOUD_PROJECT_ID is not set. Google Secret Manager will not be used.")
	}
}

// Get retrieves a secret. It follows a specific order of precedence:
// 1. Google Secret Manager (if GCLOUD_PROJECT_ID is set)
// 2. Environment variables (which includes those loaded from a .env file)
//
// If the secret is not found in any of these locations, it returns an error.
func Get(secretName string) (string, error) {
	// Attempt to get the secret from Google Secret Manager first.
	projectID, isCloudRun := os.LookupEnv("GCLOUD_PROJECT_ID")
	if isCloudRun && projectID != "" {
		secretValue, err := getFromSecretManager(projectID, secretName)
		if err == nil {
			log.Printf("Loaded '%s' from Secret Manager", secretName)
			return secretValue, nil
		}
		log.Printf("Could not fetch '%s' from Secret Manager, falling back to environment variables: %v", secretName, err)
	}

	// Fallback to environment variables.
	if value, ok := os.LookupEnv(secretName); ok {
		log.Printf("Loaded '%s' from .env file or environment", secretName)
		return value, nil
	}

	return "", fmt.Errorf("secret '%s' not found in Secret Manager, .env file, or environment variables", secretName)
}

// getFromSecretManager fetches a secret from Google Secret Manager.
func getFromSecretManager(projectID, secretName string) (string, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create secret manager client: %v", err)
	}
	defer client.Close()

	name := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, secretName)
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to access secret version: %v", err)
	}

	return string(result.Payload.Data), nil
}
