package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "test.log")
	err := os.WriteFile(logFile, []byte("line1\n"+
		"line2\n"+
		"line3\n"+
		"line4\n"+
		"line5\n"+
		"line6\n"+
		"line7\n"+
		"line8\n"+
		"line9\n"+
		"line10\n"+
		"line11\n"+
		"line12\n"+
		"line13\n"+
		"line14\n"+
		"line15\n"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	logFileWith5Lines := filepath.Join(tmpDir, "test2.log")

	err2 := os.WriteFile(logFileWith5Lines, []byte("line1\n"+
		"line2\n"+
		"line3\n"+
		"line4\n"+
		"line5\n"), 0644)
	if err2 != nil {
		t.Fatal(err2)
	}

	tests := []struct {
		name       string
		args       []string
		wantError  bool
		wantOutput string
	}{
		{
			// Simulates: go run ./cmd/main.go
			name:       "Missing logfile",
			args:       []string{"gtail"},
			wantError:  true,
			wantOutput: "",
		},
		{
			// Simulates: go run ./cmd/main.go missingfile.log
			name:       "File does not exist",
			args:       []string{"gtail", "missingfile.log"},
			wantError:  true,
			wantOutput: "",
		},
		{
			name:       "Print last 10 lines",
			args:       []string{"gtail", logFile},
			wantError:  false,
			wantOutput: "line6\nline7\nline8\nline9\nline10\nline11\nline12\nline13\nline14\nline15\n",
		},
		{
			name:       "Print all lines if less than 10",
			args:       []string{"gtail", logFileWith5Lines},
			wantError:  false,
			wantOutput: "line1\nline2\nline3\nline4\nline5\n",
		},
		{
			name:       "Print last 3 lines using -n",
			args:       []string{"gtail", "-n", "3", logFile},
			wantError:  false,
			wantOutput: "line13\nline14\nline15\n",
		},
	}

	// Execute
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := run(tt.args, &buf)

			if (err != nil) != tt.wantError {
				t.Errorf("run(%v) error = %v, wantError = %v", tt.args, err, tt.wantError)
				return
			}
			gotOutput := buf.String()
			if gotOutput != tt.wantOutput {
				t.Errorf("run(%v) gotOutput = %v, wantOutput = %v", tt.args, gotOutput, tt.wantOutput)
			}
		})
	}
}
