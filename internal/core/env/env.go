package env

import (
	"os"
	"strconv"
)

// Get retrieves the value of the environment variable named by the key.
// If the variable is not present, it returns the defaultValue.
func Get(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// GetRequired retrieves the value of the environment variable named by the key.
// If the variable is not present, it panics.
func GetRequired(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic("Required environment variable not set: " + key)
	}
	return value
}

// GetBool retrieves the value of the environment variable named by the key as a boolean.
// It returns true if the value is "true", "1", "yes", or "y" (case-insensitive).
// If the variable is not present, it returns the defaultValue.
func GetBool(key string, defaultValue bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	switch value {
	case "true", "1", "yes", "y", "TRUE", "YES", "Y":
		return true
	default:
		return false
	}
}

// GetInt retrieves the value of the environment variable named by the key as an integer.
// If the variable is not present or cannot be converted to an integer, it returns the defaultValue.
func GetInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return intValue
}
