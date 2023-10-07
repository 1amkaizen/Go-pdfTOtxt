// pdfutil.go

package main

import (
	"os"
	"os/exec"
)

func extractTextFromPDF(inputPDFPath string, outputTextPath string) error {
	cmd := exec.Command("pdftotext", inputPDFPath, outputTextPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

