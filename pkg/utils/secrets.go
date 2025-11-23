package utils

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

// GetSecret fetches a secret payload from Google Secret Manager.
// secretName should be the name of the secret (e.g., "BIBLE_API_KEY").
// The function constructs the full resource name: projects/{projectID}/secrets/{secretName}/versions/latest
func GetSecret(projectID, secretName string) (string, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create secret manager client: %v", err)
	}
	defer client.Close()

	// Build the request.
	// We use "latest" to fetch the most recent version of the secret.
	name := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, secretName)

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to access secret version: %v", err)
	}

	// Return the secret payload.
	return string(result.Payload.Data), nil
}
