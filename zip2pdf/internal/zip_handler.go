package internal

import (
	"archive/zip"
	"fmt"
	"io"
	// "os"
	// "path/filepath"
)

// FileData holds the path and content of a file.
type FileData struct {
	Path    string
	Content []byte
}

// ExtractZip reads a zip file and returns the content of its files.
// It creates a temporary directory to extract files, reads them, and then cleans up.
func ExtractZip(zipPath string) ([]FileData, error) {
	// Open the zip file for reading.
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open zip file: %w", err)
	}
	defer r.Close()

	var filesData []FileData

	// Iterate through each file in the zip archive.
	for _, f := range r.File {
		// We only process files, not directories.
		if f.FileInfo().IsDir() {
			continue
		}

		// Open the file inside the zip.
		rc, err := f.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file in zip (%s): %w", f.Name, err)
		}

		// Read the content of the file.
		content, err := io.ReadAll(rc)
		rc.Close() // Close the reader for the current file.
		if err != nil {
			return nil, fmt.Errorf("failed to read file in zip (%s): %w", f.Name, err)
		}

		// Store the file path and its content.
		filesData = append(filesData, FileData{
			Path:    f.Name,
			Content: content,
		})
	}

	return filesData, nil
}
