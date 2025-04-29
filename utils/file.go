package utils

import (
	"encoding/base64"
	"io"
	"os"
)

// FileToBase64String reads a file at the given path and returns its base64 encoded string
func FileToBase64String(filePath string) (string, error) {
	// Read the file
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Encode to base64
	base64String := base64.StdEncoding.EncodeToString(fileBytes)
	return base64String, nil
}

// FileToBase64Writer reads a file at the given path and writes its base64 encoded content to the provided writer
func FileToBase64Writer(filePath string, output io.Writer) error {
	// Open input file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create base64 encoder that writes to output
	encoder := base64.NewEncoder(base64.StdEncoding, output)
	defer encoder.Close()

	// Copy data from file to encoder
	_, err = io.Copy(encoder, file)
	return err
}

// Base64EncodeToFile reads a file at inputPath and writes its base64 encoded content to outputPath
func Base64EncodeToFile(inputPath, outputPath string) error {
	// Open output file
	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Use our writer function
	return FileToBase64Writer(inputPath, out)
}
