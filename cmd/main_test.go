package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "test.log")
	err := os.WriteFile(logFile, []byte("line1\nline2\n"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name      string
		args      []string
		wantError bool
	}{
		{
			// Simulates: go run ./cmd/main.go
			name:      "Missing logfile",
			args:      []string{"gtail"},
			wantError: true,
		},
	}

	// Execute
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := run(tt.args)

			if (err != nil) != tt.wantError {
				t.Errorf("run(%v) error = %v, wantError = %v", tt.args, err, tt.wantError)
			}
		})
	}
}
