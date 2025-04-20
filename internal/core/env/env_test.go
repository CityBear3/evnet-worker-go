package env

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	// Define test cases
	tests := []struct {
		name         string
		key          string
		defaultValue string
		envValue     string
		setEnv       bool
		want         string
	}{
		{
			name:         "existing environment variable",
			key:          "TEST_ENV_VAR",
			defaultValue: "default_value",
			envValue:     "test_value",
			setEnv:       true,
			want:         "test_value",
		},
		{
			name:         "non-existing environment variable",
			key:          "NON_EXISTING_ENV_VAR",
			defaultValue: "default_value",
			setEnv:       false,
			want:         "default_value",
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setEnv {
				os.Setenv(tt.key, tt.envValue)
				defer os.Unsetenv(tt.key)
			}

			// Test
			if got := Get(tt.key, tt.defaultValue); got != tt.want {
				t.Errorf("Get(%q, %q) = %q, want %q", tt.key, tt.defaultValue, got, tt.want)
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	// Define test cases
	tests := []struct {
		name         string
		key          string
		defaultValue bool
		envValue     string
		setEnv       bool
		want         bool
	}{
		{
			name:         "true value - 'true'",
			key:          "TEST_BOOL_VAR",
			defaultValue: false,
			envValue:     "true",
			setEnv:       true,
			want:         true,
		},
		{
			name:         "true value - '1'",
			key:          "TEST_BOOL_VAR",
			defaultValue: false,
			envValue:     "1",
			setEnv:       true,
			want:         true,
		},
		{
			name:         "true value - 'yes'",
			key:          "TEST_BOOL_VAR",
			defaultValue: false,
			envValue:     "yes",
			setEnv:       true,
			want:         true,
		},
		{
			name:         "true value - 'y'",
			key:          "TEST_BOOL_VAR",
			defaultValue: false,
			envValue:     "y",
			setEnv:       true,
			want:         true,
		},
		{
			name:         "true value - 'TRUE'",
			key:          "TEST_BOOL_VAR",
			defaultValue: false,
			envValue:     "TRUE",
			setEnv:       true,
			want:         true,
		},
		{
			name:         "true value - 'YES'",
			key:          "TEST_BOOL_VAR",
			defaultValue: false,
			envValue:     "YES",
			setEnv:       true,
			want:         true,
		},
		{
			name:         "true value - 'Y'",
			key:          "TEST_BOOL_VAR",
			defaultValue: false,
			envValue:     "Y",
			setEnv:       true,
			want:         true,
		},
		{
			name:         "false value",
			key:          "TEST_BOOL_VAR",
			defaultValue: true,
			envValue:     "false",
			setEnv:       true,
			want:         false,
		},
		{
			name:         "non-existing environment variable",
			key:          "NON_EXISTING_BOOL_VAR",
			defaultValue: true,
			setEnv:       false,
			want:         true,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setEnv {
				os.Setenv(tt.key, tt.envValue)
				defer os.Unsetenv(tt.key)
			}

			// Test
			if got := GetBool(tt.key, tt.defaultValue); got != tt.want {
				t.Errorf("GetBool(%q, %v) = %v, want %v", tt.key, tt.defaultValue, got, tt.want)
			}
		})
	}
}

func TestGetRequired(t *testing.T) {
	// Define test cases
	tests := []struct {
		name      string
		key       string
		envValue  string
		setEnv    bool
		want      string
		wantPanic bool
	}{
		{
			name:      "existing environment variable",
			key:       "TEST_REQUIRED_ENV_VAR",
			envValue:  "required_value",
			setEnv:    true,
			want:      "required_value",
			wantPanic: false,
		},
		{
			name:      "non-existing environment variable",
			key:       "NON_EXISTING_REQUIRED_ENV_VAR",
			setEnv:    false,
			wantPanic: true,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setEnv {
				os.Setenv(tt.key, tt.envValue)
				defer os.Unsetenv(tt.key)
			}

			// Test
			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("GetRequired(%q) did not panic as expected", tt.key)
					}
				}()
				GetRequired(tt.key) // This should panic
			} else {
				if got := GetRequired(tt.key); got != tt.want {
					t.Errorf("GetRequired(%q) = %q, want %q", tt.key, got, tt.want)
				}
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	// Define test cases
	tests := []struct {
		name         string
		key          string
		defaultValue int
		envValue     string
		setEnv       bool
		want         int
	}{
		{
			name:         "valid integer value",
			key:          "TEST_INT_VAR",
			defaultValue: 0,
			envValue:     "42",
			setEnv:       true,
			want:         42,
		},
		{
			name:         "invalid (non-integer) value",
			key:          "TEST_INVALID_INT_VAR",
			defaultValue: 99,
			envValue:     "not-an-integer",
			setEnv:       true,
			want:         99,
		},
		{
			name:         "non-existing environment variable",
			key:          "NON_EXISTING_INT_VAR",
			defaultValue: 99,
			setEnv:       false,
			want:         99,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setEnv {
				os.Setenv(tt.key, tt.envValue)
				defer os.Unsetenv(tt.key)
			}

			// Test
			if got := GetInt(tt.key, tt.defaultValue); got != tt.want {
				t.Errorf("GetInt(%q, %d) = %d, want %d", tt.key, tt.defaultValue, got, tt.want)
			}
		})
	}
}
