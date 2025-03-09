package utils

import "os"

// GetEnvWithDefault returns the value of the environment variable with the given key.
// If the environment variable is not set, it returns the default value.
func GetEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
