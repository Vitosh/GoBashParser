package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var singleLine = "------------------------------------------------------------------------"

func TestFetchURL(t *testing.T) {
	// Tests 5 URLs, for responses, errors included
	t.Logf(singleLine)

	testCases := []struct {
		url            string
		expectedOutput string
		expectedErr    error
		statusCode     int
	}{
		{"https://www.google.com", "", nil, http.StatusOK},
		{"https://www.github.com", "", nil, http.StatusOK},
		{"https://www.bbc.com", "", nil, http.StatusOK},
		{"https://github.com/404", "", fmt.Errorf("unexpected status code 404"), http.StatusNotFound},
		{"https://httpbin.org/status/500", "", fmt.Errorf("unexpected status code 500"), http.StatusInternalServerError},
	}
	t.Logf(singleLine)

	for _, tc := range testCases {
		t.Run(tc.url, func(t *testing.T) {
			var buf bytes.Buffer
			err := fetchURL(tc.url, &buf)
			if tc.expectedErr != nil && err == nil {
				t.Errorf("expected error %v but got nil", tc.expectedErr)
			} else if tc.expectedErr == nil && err != nil {
				t.Errorf("unexpected error %v", err)
			} else if tc.expectedErr != nil && err != nil && tc.expectedErr.Error() != err.Error() {
				t.Errorf("expected error %v but got %v", tc.expectedErr, err)
			}
			if !strings.Contains(buf.String(), tc.expectedOutput) {
				t.Errorf("unexpected output for url %q", tc.url)
			}
			if tc.statusCode != http.StatusOK {
				if err == nil || !strings.Contains(err.Error(), fmt.Sprintf("unexpected status code %d", tc.statusCode)) {
					t.Errorf("unexpected status code for url %q: got %v, expected %d", tc.url, err, tc.statusCode)
				}
			}
		})
	}
	t.Logf("TestFetchURL() test completed successfully")
	t.Logf(singleLine)
}

func TestInformation(t *testing.T) {
	// Tests the information()

	t.Logf(singleLine)

	commands := []string{"ls", "cd", "mkdir", "pwd", "rm", "touch", "curl", "log", "info", "exit"}
	output := information()

	for _, cmd := range commands {
		t.Run(cmd, func(t *testing.T) {
			if !strings.Contains(output, cmd) {
				t.Errorf("information() output does not contain command %s.", cmd)
			}
			t.Logf("information() output contains command %q", cmd)
		})
	}

	if !strings.Contains(output, "Thank you for using our program") {
		t.Errorf("information() output does not contain farewell message.")
	}

	if !strings.Contains(output, "Your choice is awaited!") {
		t.Errorf("information() output does not contain welcome message.")
	}

	t.Logf("information() test completed successfully")
	t.Logf(singleLine)

}

func TestCreateFile(t *testing.T) {
	//Tests for createFile()
	t.Logf(singleLine)

	tmpDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	path := filepath.Join(tmpDir, "testfile.txt")

	f, err := createFile(path)
	if err != nil {
		t.Fatal(err)
	}

	f.Close()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("file was not created at %s", path)
	}

	if err := os.Remove(path); err != nil {
		t.Errorf("error removing test file: %v", err)
	}

	t.Logf("createFile() test completed successfully")
	t.Logf(singleLine)

}

func TestRemoveFile(t *testing.T) {
	// Test removeFile()
	// Create a temporary file for testing

	t.Logf(singleLine)

	f, err := os.CreateTemp("", "testfile.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(f.Name())

	err = f.Close()
	if err != nil {
		t.Errorf("Failed to close file: %v", err)
	}

	// Call the removeFile function and check if the file was removed successfully
	removeFile(f.Name())

	if _, err := os.Stat(f.Name()); !os.IsNotExist(err) {
		t.Errorf("Failed to remove file: %v", err)
	}
	t.Logf("removeFile() test completed successfully")
	t.Logf(singleLine)

}
