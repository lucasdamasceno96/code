// zip2pdf/internal/service.go
package internal

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"
)

// ConversionService orchestrates the zip to PDF conversion process.
type ConversionService struct{}

// NewConversionService creates a new service instance.
func NewConversionService() *ConversionService {
	return &ConversionService{}
}

// ConvertZipToPDF handles the full conversion logic.
func (s *ConversionService) ConvertZipToPDF(zipPath string) (string, error) {
	log.Printf("Starting conversion for: %s", zipPath)
	startTime := time.Now()

	// Step 1: Extract the zip file.
	log.Println("Extracting zip file...")
	filesData, err := ExtractZip(zipPath)
	if err != nil {
		return "", fmt.Errorf("extraction failed: %w", err)
	}
	log.Printf("Extracted %d files.", len(filesData))

	// Step 2: Generate the PDF.
	log.Println("Generating PDF...")
	// Define the output path for the PDF.
	baseName := strings.TrimSuffix(filepath.Base(zipPath), filepath.Ext(zipPath))
	outputPath := filepath.Join("output", fmt.Sprintf("%s_content.pdf", baseName))

	if err := GeneratePDF(filesData, outputPath); err != nil {
		return "", fmt.Errorf("pdf generation failed: %w", err)
	}

	duration := time.Since(startTime)
	log.Printf("PDF generation completed in %s.", duration)

	return outputPath, nil
}
