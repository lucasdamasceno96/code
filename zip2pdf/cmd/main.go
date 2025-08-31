// zip2pdf/cmd/main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lucasdamasceno96/code/zip2pdf/internal" // Import our internal package
)

func main() {
	// Configure logger to have a prefix and no timestamp.
	log.SetFlags(0)
	log.SetPrefix("✅ ")

	// 1. Check for command-line arguments.
	if len(os.Args) < 2 {
		log.Println("Error: Missing file path argument.")
		fmt.Println("Usage: ./zip2pdf <path_to_zip_file>")
		os.Exit(1)
	}
	zipPath := os.Args[1]

	// 2. Check if the file exists.
	if _, err := os.Stat(zipPath); os.IsNotExist(err) {
		log.Printf("Error: File not found at '%s'", zipPath)
		os.Exit(1)
	}

	// 3. Initialize and run the conversion service.
	service := internal.NewConversionService()
	outputPath, err := service.ConvertZipToPDF(zipPath)
	if err != nil {
		log.Printf("Error: Conversion failed: %v", err)
		os.Exit(1)
	}

	// 4. Print success message.
	log.Printf("Success! PDF generated at: %s\n", outputPath)
}
