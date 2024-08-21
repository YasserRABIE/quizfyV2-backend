package utils

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"github.com/YasserRABIE/QUIZFYv2/config"
)

var (
	BasePath = getBasePath()
)

func getBasePath() string {
	return config.GetEnv("QUIZ_UPLOAD_PATH", "uploads/quiz/")
}

// UploadImage uploads an image to the server and returns the path to the image
func UploadImage(imageBase64 string, extension string, quizID, questionID uint) (string, error) {
	// Decode the base64 image data
	imageData, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		return "", err
	}

	// Check if the image size is too large (e.g., more than 5MB)
	if len(imageData) > 5*1024*1024 {
		return "", fmt.Errorf("file size exceeds limit")
	}

	// Ensure the base path exists
	if err := os.MkdirAll(BasePath, 0755); err != nil {
		return "", err
	}

	// Ensure the quiz directory exists
	quizPath := filepath.Join(BasePath, fmt.Sprintf("%d", quizID))
	if err := os.MkdirAll(quizPath, 0755); err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("question_%d.%s", questionID, extension)
	destPath := filepath.Join(quizPath, fileName)

	// Write the image data to the file
	err = os.WriteFile(destPath, imageData, 0644)
	if err != nil {
		return "", err
	}

	return destPath, nil
}
