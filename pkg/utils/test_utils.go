package utils

import (
	"os"
)

// SetEnv sets an environment variable and returns a function to restore it.
func SetEnv(key, value string) func() {
	originalValue, exists := os.LookupEnv(key)
	os.Setenv(key, value)
	return func() {
		if exists {
			os.Setenv(key, originalValue)
		} else {
			os.Unsetenv(key)
		}
	}
}
