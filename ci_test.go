package main

import (
	"os"
	"testing"
)

// TestCIEnvPropagation checks that environment propagation is working.
// The variable was originally set in the Travis UI for this specific repo.
// It's propagated by rembuild (per config in yourbase.json) and made available
// to this test.
func TestCIEnvPropagation(t *testing.T) {
	tests := []struct {
		env  string
		want string
	}{
		{"TEST_ENV_IS_SET", "1"},
	}
	for _, test := range tests {
		envVarContent := os.Getenv(test.env)
		if envVarContent != test.want {
			t.Fatalf("For env %q got content: %q, wanted %q", test.env, envVarContent, test.want)
		}
	}
}
