package main

import (
	"fmt"
	"os"

	"github.com/vanclief/go-facturama/utils"
)

func main() {
	// Example 1: Get base64 as a string
	base64String, err := utils.FileToBase64String("example.jpg")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Base64 string (first 30 chars): %s...\n", base64String[:30])

	// Example 2: Write base64 directly to stdout
	fmt.Println("\nWriting file directly to base64:")
	err = utils.FileToBase64Writer("example.jpg", os.Stdout)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Example 3: Convert file to another file
	err = utils.Base64EncodeToFile("example.jpg", "example.jpg.b64")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("\nFile successfully encoded to example.jpg.b64")
}
