package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2/v6"
)

func renderPage(templatePath, outputFilePath string, data interface{}) {
	tpl, err := pongo2.FromFile(templatePath)
	if err != nil {
		log.Fatalf("Error loading template from file %s: %v", templatePath, err)
	}
	outputDir := filepath.Dir(outputFilePath)
	ensureDir(outputDir)

	file, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("Error creating file %s: %v", outputFilePath, err)
	}
	defer file.Close()

	if err := tpl.ExecuteWriter(pongo2.Context{"data": data, "sitename": "Kaustubha"}, file); err != nil {
		log.Fatalf("Error executing template and writing to file %s: %v", outputFilePath, err)
	}
	log.Printf("Successfully rendered template to file: %s", outputFilePath)
}

func ensureDir(dir string) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatalf("Failed to create directory %s: %v", dir, err)
	}
}

func copyStatic(srcDir, destDir string) {
	files, err := os.ReadDir(srcDir)
	if err != nil {
		log.Fatalf("Failed to read directory %s: %v", srcDir, err)
	}

	if err := os.MkdirAll(destDir, 0755); err != nil {
		log.Fatalf("Failed to create destination directory %s: %v", destDir, err)
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "!") {
			continue
		}
		srcPath := filepath.Join(srcDir, file.Name())
		destPath := filepath.Join(destDir, file.Name())

		if file.IsDir() {
			continue
		}

		srcFile, err := os.Open(srcPath)
		if err != nil {
			log.Fatalf("Failed to open source file %s: %v", srcPath, err)
		}
		defer srcFile.Close()

		destFile, err := os.Create(destPath)
		if err != nil {
			log.Fatalf("Failed to create destination file %s: %v", destPath, err)
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			log.Fatalf("Failed to copy file content from %s to %s: %v", srcPath, destPath, err)
		}
	}
}
