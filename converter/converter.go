package converter

import (
	"os"
	"path/filepath"
)

type Converter struct {
	//можна додати конфігурацію
}

// створює та повертає новий екземпляр Converter
func NewConverter() *Converter {
	return &Converter{}
}

func (c *Converter) TextToPDF(inputPath, outputPath string, overwrite bool) error {
	return textToPDF(inputPath, outputPath, overwrite)
}

func (c *Converter) ImagesToPDF(inputPaths []string, outputPath string, overwrite bool) error {
	return imagesToPDF(inputPaths, outputPath, overwrite)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// створює директорію для outputPath, наприклад для path/to/out.pdf створить path/to
func ensureOutputDir(outputPath string) error {
	dir := filepath.Dir(outputPath)
	// Якщо вказано тільки ім'я файлу, то нічого не створюємо
	if dir == "." || dir == "" {
		return nil
	}
	return os.MkdirAll(dir, 0o755)
}
