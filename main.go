package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func deleteSoloFile(dir string) error {
	fileMap := make(map[string]string)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			base := strings.TrimSuffix(info.Name(), ext)
			fullPath := filepath.Join(dir, info.Name())

			// Check if the file is a solo file
			if _, exists := fileMap[base]; exists {
				delete(fileMap, base)
			} else {
				fileMap[base] = fullPath
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to walk directory: %w", err)
	}

	// Delete solo files
	for _, path := range fileMap {
		if err := os.Remove(path); err != nil {
			fmt.Printf("Failed to delete %s: %v\n", path, err)
		} else {
			fmt.Printf("Deleted %s\n", path)
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: solo-cleaner <dir>")
	}
	dir := os.Args[1]

	if err := deleteSoloFile(dir); err != nil {
		log.Println(err)
	}
}
