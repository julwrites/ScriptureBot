package secrets

import (
	"fmt"
	"os"
	"testing"
)

// setupTestEnvFile creates a temporary .env file for testing.
func setupTestEnvFile(t *testing.T, content string) func() {
	t.Helper()
	tmpfile, err := os.Create(".env")
	if err != nil {
		t.Fatalf("Failed to create temporary .env file: %v", err)
	}
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temporary .env file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temporary .env file: %v", err)
	}

	// The cleanup function to be returned and deferred by the caller.
	return func() {
		os.Remove(tmpfile.Name())
		// Reset environment variables changed by godotenv.Overload()
		os.Unsetenv("FROM_DOTENV")
		os.Unsetenv("OVERLOAD_VAR")
	}
}

func TestGetSecret(t *testing.T) {
	// Sub-test 1: Test loading from an environment variable.
	t.Run("from environment variable", func(t *testing.T) {
		const envVarName = "FROM_ENV"
		const expectedValue = "env_value"
		t.Setenv(envVarName, expectedValue)

		// Re-run the loading logic to pick up the new env var for the test.
		LoadAndLog()

		val, err := Get(envVarName)
		if err != nil {
			t.Errorf("Get() returned an error: %v", err)
		}
		if val != expectedValue {
			t.Errorf("Expected to get '%s', but got '%s'", expectedValue, val)
		}
	})

	// Sub-test 2: Test loading from a .env file.
	t.Run("from .env file", func(t *testing.T) {
		const secretName = "FROM_DOTENV"
		const expectedValue = "dotenv_value"
		cleanup := setupTestEnvFile(t, fmt.Sprintf("%s=%s", secretName, expectedValue))
		defer cleanup()

		// Re-run the loading logic to load the .env file.
		LoadAndLog()

		val, err := Get(secretName)
		if err != nil {
			t.Errorf("Get() returned an error: %v", err)
		}
		if val != expectedValue {
			t.Errorf("Expected to get '%s' from .env, but got '%s'", expectedValue, val)
		}
	})

	// Sub-test 3: Test that .env file takes precedence over environment variables.
	t.Run(".env overloads environment variable", func(t *testing.T) {
		const varName = "OVERLOAD_VAR"
		const envValue = "from_shell"
		const dotenvValue = "from_dotenv_file"

		// Set the environment variable first.
		t.Setenv(varName, envValue)

		// Create the .env file with the same variable but a different value.
		cleanup := setupTestEnvFile(t, fmt.Sprintf("%s=%s", varName, dotenvValue))
		defer cleanup()

		// Re-run the loading logic. godotenv.Overload should prioritize the .env file.
		LoadAndLog()

		val, err := Get(varName)
		if err != nil {
			t.Errorf("Get() returned an error: %v", err)
		}
		if val != dotenvValue {
			t.Errorf("Expected value from .env ('%s'), but got value from shell ('%s')", dotenvValue, val)
		}
	})

	// Sub-test 4: Test error when secret is not found.
	t.Run("secret not found", func(t *testing.T) {
		const nonExistentSecret = "THIS_SECRET_SHOULD_NOT_EXIST"
		// Ensure the variable is not set in the environment.
		os.Unsetenv(nonExistentSecret)

		// Re-run the loading logic.
		LoadAndLog()

		_, err := Get(nonExistentSecret)
		if err == nil {
			t.Error("Expected an error when getting a non-existent secret, but got nil")
		}
	})
}

func TestLoadSecrets(t *testing.T) {
	// Mock environment variables
	t.Setenv("TELEGRAM_ID", "test_telegram_id")
	t.Setenv("TELEGRAM_ADMIN_ID", "test_admin_id")
	t.Setenv("GCLOUD_PROJECT_ID", "test_project_id")

	secrets, err := LoadSecrets()
	if err != nil {
		t.Fatalf("LoadSecrets failed: %v", err)
	}

	if secrets.TELEGRAM_ID != "test_telegram_id" {
		t.Errorf("Expected TELEGRAM_ID 'test_telegram_id', got '%s'", secrets.TELEGRAM_ID)
	}
	if secrets.TELEGRAM_ADMIN_ID != "test_admin_id" {
		t.Errorf("Expected TELEGRAM_ADMIN_ID 'test_admin_id', got '%s'", secrets.TELEGRAM_ADMIN_ID)
	}
	if secrets.PROJECT_ID != "test_project_id" {
		t.Errorf("Expected PROJECT_ID 'test_project_id', got '%s'", secrets.PROJECT_ID)
	}
}
