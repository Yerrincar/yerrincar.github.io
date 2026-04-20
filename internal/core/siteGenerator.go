package site

import (
	"fmt"
	htmltemplate "html/template"
	"os"
	"path/filepath"
)

type Generator struct {
	templatesDir string
	outputDir    string
}

func NewGenerator() Generator {
	return Generator{
		templatesDir: filepath.Join("ui", "web", "templates"),
		outputDir:    "dist",
	}
}

func (g Generator) Generate() error {
	if err := os.MkdirAll(g.outputDir, 0o755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	if err := g.generateHome(); err != nil {
		return fmt.Errorf("generate home page: %w", err)
	}

	return nil
}

func (g Generator) generateHome() error {
	templatePath := filepath.Join(g.templatesDir, "home.html.tmpl")
	outputPath := filepath.Join(g.outputDir, "index.html")

	tmpl, err := htmltemplate.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("parse %q: %w", templatePath, err)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create %q: %w", outputPath, err)
	}
	defer outputFile.Close()

	if err := tmpl.Execute(outputFile, nil); err != nil {
		return fmt.Errorf("render %q: %w", templatePath, err)
	}

	return nil
}
