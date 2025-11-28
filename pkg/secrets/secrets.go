package secrets

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

// SecretsData holds all the secrets for the application.
type SecretsData struct {
	TELEGRAM_ID string
	PROJECT_ID  string
	// Add other secrets here as needed
}

func init() {
	LoadAndLog()
}

// LoadAndLog loads environment variables from a .env file (if present) and logs
// the status of the GCLOUD_PROJECT_ID.
func LoadAndLog() {
	err := godotenv.Overload()
	if err != nil {
		log.Println("No .env file found, using environment variables.")
	}
	if projectID, ok := os.LookupEnv("GCLOUD_PROJECT_ID"); ok {
		log.Printf("GCLOUD_PROJECT_ID is set: %s", projectID)
	} else {
		log.Println("GCLOUD_PROJECT_ID is not set. Assuming local development.")
	}
}

// LoadSecrets populates the SecretsData struct by fetching secrets.
func LoadSecrets() (SecretsData, error) {
	projectID := os.Getenv("GCLOUD_PROJECT_ID")

	var secrets SecretsData
	secrets.PROJECT_ID = projectID

	var wg sync.WaitGroup
	var errs = make(chan error, 1) // Buffer to hold the first error

	// List of secret names to fetch
	secretNames := []string{"TELEGRAM_ID"} // Add other secret names here

	for _, secretName := range secretNames {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			value, err := Get(name)
			if err != nil {
				select {
				case errs <- fmt.Errorf("failed to load secret '%s': %v", name, err):
				default:
				}
				return
			}
			switch name {
			case "TELEGRAM_ID":
				secrets.TELEGRAM_ID = value
			}
		}(secretName)
	}

	wg.Wait()
	close(errs)

	if err := <-errs; err != nil {
		return SecretsData{}, err
	}

	return secrets, nil
}

// Get retrieves a secret.
// It prioritizes environment variables. If not found, and GCLOUD_PROJECT_ID is set,
// it fetches from Google Secret Manager.
func Get(secretName string) (string, error) {
	// Check environment variables first.
	// This allows overriding secrets for local development or testing.
	if value, ok := os.LookupEnv(secretName); ok {
		log.Printf("Loaded '%s' from environment", secretName)
		return value, nil
	}

	projectID, isCloudRun := os.LookupEnv("GCLOUD_PROJECT_ID")
	if isCloudRun && projectID != "" {
		// Cloud environment: Use Secret Manager if not found in environment.
		secretValue, err := getFromSecretManager(projectID, secretName)
		if err != nil {
			return "", fmt.Errorf("failed to get secret '%s' from Secret Manager: %v", secretName, err)
		}
		log.Printf("Loaded '%s' from Secret Manager", secretName)
		return secretValue, nil
	}

	return "", fmt.Errorf("secret '%s' not found in environment variables", secretName)
}

func getFromSecretManager(projectID, secretName string) (string, error) {
	ctx := context.Background()

	var client *secretmanager.Client
	var err error

	if saKey, ok := os.LookupEnv("GCLOUD_SA_KEY"); ok && saKey != "" {
		// Authenticate with the service account key if provided
		client, err = secretmanager.NewClient(ctx, option.WithCredentialsJSON([]byte(saKey)))
	} else {
		// Fallback to Application Default Credentials (ADC)
		client, err = secretmanager.NewClient(ctx)
	}

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
