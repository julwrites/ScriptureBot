package app

import "os"

// setEnv sets an environment variable and returns a function to restore it
func setEnv(key, value string) func() {
	// Reset config cache before setting env var
	resetAPIConfigCache()
	originalValue, exists := os.LookupEnv(key)
	os.Setenv(key, value)
	return func() {
		if exists {
			os.Setenv(key, originalValue)
		} else {
			os.Unsetenv(key)
		}
		// Reset config cache after restoring
		resetAPIConfigCache()
	}
}
