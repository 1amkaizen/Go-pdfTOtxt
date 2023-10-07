// pdf_extractor.go

package main

import (
	"os/exec"
)



func getPDFMetadata(inputPDFPath string) (string, error) {
	cmd := exec.Command("pdfinfo", inputPDFPath)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

