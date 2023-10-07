// main.go

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./program <input_pdf_file> <output_text_file>")
		return
	}

	pdfFilePath := os.Args[1]
	outputFilePath := os.Args[2]

	// Ekstraksi teks dari PDF
	err := extractTextFromPDF(pdfFilePath, outputFilePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Mendapatkan metadata dari PDF
	metadata, err := getPDFMetadata(pdfFilePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("PDF Metadata:")
	fmt.Println(metadata)
}

