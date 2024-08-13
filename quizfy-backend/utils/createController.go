package utils

import (
	"os"
	"path/filepath"
)

// CreateFolder creates a folder in the given path.
// It also creates the files create.go, get.go, update.go, delete.go in the folder.
func CreateFolder(pkgName, path string) error {
	if err := os.Mkdir(path, 0755); err != nil {
		return err
	}
	createFiles(pkgName, path)
	return nil
}

func createFiles(pkgName, path string) {
	files := []string{"create.go", "get.go", "update.go", "delete.go"}
	for _, file := range files {
		fullPath := filepath.Join(path, file)

		// Create or open the file
		f, err := os.Create(fullPath)
		if err != nil {
			panic(err)
		}

		// Write package declaration
		if _, err := f.WriteString("package " + pkgName); err != nil {
			panic(err)
		}

		// Close the file
		if err := f.Close(); err != nil {
			panic(err)
		}
	}
}
