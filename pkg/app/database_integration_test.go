package app

import (
	"testing"

	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/secrets"
	"github.com/julwrites/ScriptureBot/pkg/utils"
)

func TestUserDatabaseIntegration(t *testing.T) {
	// This test performs a live database operation against the configured project.
	// It relies on GCLOUD_PROJECT_ID being set.

	secretsData, err := secrets.LoadSecrets()
	if err != nil {
		t.Logf("Warning: Could not load secrets: %v", err)
	}

	projectID := secretsData.PROJECT_ID
	if projectID == "" {
		t.Skip("Skipping database test: GCLOUD_PROJECT_ID not set")
	}

	// Use a unique ID to avoid conflict with real users
	dummyID := "test-integration-user-DO-NOT-DELETE"

	var user def.UserData
	user.Id = dummyID
	user.Firstname = "Integration"
	user.Lastname = "Test"
	user.Username = "TestUser"
	user.Type = "Private"

	// Create/Update user
	// This exercises the connection to Datastore/Firestore
	updatedUser := utils.RegisterUser(user, projectID)

	if updatedUser.Id != dummyID {
		t.Errorf("Expected user ID %s, got %s", dummyID, updatedUser.Id)
	}

	// Verify update capability
	updatedUser.Action = "testing"
	finalUser := utils.RegisterUser(updatedUser, projectID)

	if finalUser.Action != "testing" {
		t.Errorf("Expected user Action 'testing', got '%s'", finalUser.Action)
	}
}
