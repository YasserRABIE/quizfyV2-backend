package utils

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

const (
	BasePath = "uploads/quiz/"
)

// UploadImage uploads an image to the server and returns the path to the image
func UploadImage(imageBase64 string, extension string, quizID, questionID uint) string {
	imageData, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		fmt.Println("Error decoding base64 string:", err)
		return ""
	}

	// Ensure the base path exists
	if err := os.MkdirAll(BasePath, os.ModePerm); err != nil {
		fmt.Println("Error creating base path:", err)
		return ""
	}

	// Ensure the quiz directory exists
	quizPath := filepath.Join(BasePath, fmt.Sprintf("%d", quizID))
	if err := os.MkdirAll(quizPath, os.ModePerm); err != nil {
		fmt.Println("Error creating quiz directory:", err)
		return ""
	}

	fileName := fmt.Sprintf("question_%d.%s", questionID, extension)
	destPath := filepath.Join(quizPath, fileName)

	// Write the image data to the file
	err = os.WriteFile(destPath, imageData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return ""
	}

	return destPath
}
