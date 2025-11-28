package app

import (
	"os"
)

// SetEnv is a helper function to temporarily set an environment variable and return a function to restore it.
func SetEnv(key, value string) func() {
	originalValue, isSet := os.LookupEnv(key)
	os.Setenv(key, value)
	return func() {
		if isSet {
			os.Setenv(key, originalValue)
		} else {
			os.Unsetenv(key)
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
