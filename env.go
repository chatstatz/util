package util

import "os"

// GetEnv returns the value for a give environment variable,
// or the provided fallback if no value is found.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
