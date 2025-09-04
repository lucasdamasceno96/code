// zip2pdf/tests/zip_handler_test.go
package tests

import (
	"archive/zip"
	"bytes"
	"os"
	"testing"

	"github.com/lucasdamasceno96/code/zip2pdf/internal"
)

// Helper function to create a dummy zip file for testing.
func createDummyZip(t *testing.T, filePath string) {
	t.Helper()

	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("failed to create dummy zip: %v", err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	// Add a dummy file to the zip.
	w, err := zipWriter.Create("dummy.txt")
	if err != nil {
		t.Fatalf("failed to create entry in zip: %v", err)
	}
	_, err = w.Write([]byte("hello world"))
	if err != nil {
		t.Fatalf("failed to write to entry in zip: %v", err)
	}
}

func TestExtractZip(t *testing.T) {
	// Setup: create a temporary dummy zip file.
	tmpFile, err := os.CreateTemp("", "test*.zip")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up.
	createDummyZip(t, tmpFile.Name())

	// Test the extraction.
	filesData, err := internal.ExtractZip(tmpFile.Name())
	if err != nil {
		t.Fatalf("ExtractZip failed: %v", err)
	}

	if len(filesData) != 1 {
		t.Errorf("expected 1 file, got %d", len(filesData))
	}

	if filesData[0].Path != "dummy.txt" {
		t.Errorf("expected file path 'dummy.txt', got '%s'", filesData[0].Path)
	}

	expectedContent := []byte("hello world")
	if !bytes.Equal(filesData[0].Content, expectedContent) {
		t.Errorf("expected content '%s', got '%s'", expectedContent, filesData[0].Content)
	}
}

func TestExtractZip_FileNotFound(t *testing.T) {
	_, err := internal.ExtractZip("non_existent_file.zip")
	if err == nil {
		t.Fatal("expected an error for a non-existent file, but got nil")
	}
}
