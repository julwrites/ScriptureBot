package app

import (
	"os"

	"github.com/julwrites/BotPlatform/pkg/def"
)

type MockBot struct{}

func (b *MockBot) Translate(body []byte) (def.SessionData, error) {
	return def.SessionData{}, nil
}

func (b *MockBot) Post(env def.SessionData) bool {
	return true
}

// SetEnv is a helper function to temporarily set an environment variable and return a function to restore it.
func SetEnv(key, value string) func() {
	originalValue, isSet := os.LookupEnv(key)
	os.Setenv(key, value)

	// Unset GCLOUD_PROJECT_ID to prevent Secret Manager usage during tests,
	// unless we are explicitly setting GCLOUD_PROJECT_ID itself.
	var projectID string
	var projectIDSet bool
	if key != "GCLOUD_PROJECT_ID" {
		projectID, projectIDSet = os.LookupEnv("GCLOUD_PROJECT_ID")
		if projectIDSet {
			os.Unsetenv("GCLOUD_PROJECT_ID")
		}
	}

	return func() {
		if isSet {
			os.Setenv(key, originalValue)
		} else {
			os.Unsetenv(key)
		}

		// Restore GCLOUD_PROJECT_ID if we unset it as a side effect
		if key != "GCLOUD_PROJECT_ID" && projectIDSet {
			os.Setenv("GCLOUD_PROJECT_ID", projectID)
		}
	}
}

// UnsetEnv is a helper function to temporarily unset an environment variable and return a function to restore it.
func UnsetEnv(key string) func() {
	originalValue, isSet := os.LookupEnv(key)
	os.Unsetenv(key)
	return func() {
		if isSet {
			os.Setenv(key, originalValue)
		}
	}
}
