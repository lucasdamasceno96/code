// zip2pdf/internal/pdf_handler.go
package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
)

// Constants for PDF layout
const (
	fontFamily   = "Courier" // Monospaced font for code
	fontSize     = 10
	headerHeight = 15
	lineHeight   = 5
	pageMargin   = 10
)

// GeneratePDF creates a single PDF document from multiple file data.
func GeneratePDF(files []FileData, outputPath string) error {
	// Create a new PDF instance with A4 size and mm units
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Set document margins
	pdf.SetMargins(pageMargin, pageMargin, pageMargin)

	// Ensure the output directory exists
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	for _, file := range files {
		// Add a new page for each file
		pdf.AddPage()

		// --- Header ---
		pdf.SetFont(fontFamily, "B", fontSize+2) // Bold for header
		pdf.Cell(0, headerHeight, file.Path)
		pdf.Ln(headerHeight)

		// --- File Content ---
		pdf.SetFont(fontFamily, "", fontSize)

		// gofpdf uses ISO-8859-1 encoding, so translate UTF-8
		tr := pdf.UnicodeTranslatorFromDescriptor("")
		contentStr := string(file.Content)

		// Write content line by line with automatic wrapping
		pdf.MultiCell(0, lineHeight, tr(contentStr), "", "", false)
	}

	// Save the PDF to the specified path
	if err := pdf.OutputFileAndClose(outputPath); err != nil {
		return fmt.Errorf("failed to save PDF: %w", err)
	}

	return nil
}
