package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateFolder creates a folder in the given path.
// It also creates the files create.go, get.go, update.go, delete.go in the folder.
func CreateFolder(pkgName, name string) error {
	// Create the controller folder
	controllerPath := filepath.Join("../controllers", name)
	if err := os.Mkdir(controllerPath, 0755); err != nil {
		return fmt.Errorf("failed to create controller folder: %w", err)
	}

	// Create the migration folder
	migrationPath := filepath.Join("../migrations", name+"_migrations")
	if err := os.Mkdir(migrationPath, 0755); err != nil {
		return fmt.Errorf("failed to create migration folder: %w", err)
	}

	// Create files in both folders
	if err := createFiles(pkgName, controllerPath); err != nil {
		return fmt.Errorf("failed to create files in controller folder: %w", err)
	}
	if err := createFiles(pkgName+"_migrations", migrationPath); err != nil {
		return fmt.Errorf("failed to create files in migration folder: %w", err)
	}

	return nil
}

// createFiles creates a set of files in the specified path with the given package name.
func createFiles(pkgName, path string) error {
	files := []string{"create.go", "get.go", "update.go", "delete.go"}

	for _, file := range files {
		fullPath := filepath.Join(path, file)

		// Create or open the file
		f, err := os.Create(fullPath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", fullPath, err)
		}

		// Write package declaration
		if _, err := f.WriteString("package " + pkgName + "\n"); err != nil {
			f.Close() // Attempt to close the file before returning the error
			return fmt.Errorf("failed to write to file %s: %w", fullPath, err)
		}

		// Close the file
		if err := f.Close(); err != nil {
			return fmt.Errorf("failed to close file %s: %w", fullPath, err)
		}
	}

	return nil
}
