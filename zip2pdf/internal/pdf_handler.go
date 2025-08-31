// zip2pdf/internal/pdf_handler.go
package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
)

const (
	fontFamily   = "Courier" // Monospaced font for code
	fontSize     = 10
	headerHeight = 15
	lineHeight   = 5
	pageMargin   = 10
)

// GeneratePDF creates a single PDF document from multiple file data.
func GeneratePDF(files []FileData, outputPath string) error {
	// Create a new PDF instance.
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Set a monospaced font that is good for code.
	pdf.AddFont(fontFamily, "", "courier.json") // Assumes you have a font definition file

	// Ensure the output directory exists.
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	for _, file := range files {
		// Add a new page for each file.
		pdf.AddPage()

		// --- Header ---
		pdf.SetFont(fontFamily, "B", fontSize+2) // Bold for the header
		pdf.Cell(0, headerHeight, file.Path)
		pdf.Ln(headerHeight)

		// --- Content ---
		pdf.SetFont(fontFamily, "", fontSize)

		// gofpdf uses ISO-8859-1 encoding, so we need to convert our UTF-8 string.
		// A simple way is to replace non-compatible runes, but a better way would be proper conversion.
		// For simplicity, we use a built-in translator.
		contentStr := string(file.Content)
		tr := pdf.UnicodeTranslatorFromDescriptor("")

		// Write content line by line using MultiCell for automatic line breaks.
		pdf.MultiCell(0, lineHeight, tr(contentStr), "", "", false)
	}

	// Save the PDF to the specified path.
	if err := pdf.OutputFileAndClose(outputPath); err != nil {
		return fmt.Errorf("failed to save PDF: %w", err)
	}

	return nil
}
