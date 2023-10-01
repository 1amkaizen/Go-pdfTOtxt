package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func displayLoading() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\rConverting PDF [%c]", r)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./program <input_pdf_file> <output_text_file>")
		return
	}

	// Input PDF file path
	pdfFilePath := os.Args[1]

	// Output text file path
	outputFilePath := os.Args[2]

	// Start goroutine to display loading animation
	go displayLoading()

	// Run pdftotext command to extract text from the PDF
	cmd := exec.Command("pdftotext", pdfFilePath, outputFilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("\nError: %v\n", err)
		return
	}

	// Stop the loading animation
	fmt.Print("\rConverting PDF Done!\n")
}
