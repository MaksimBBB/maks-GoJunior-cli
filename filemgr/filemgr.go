package filemgr

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// дає нам відсортований список імен файлів
func ListFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	list := make([]string, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		//iгноруємо файли що починаються з точки
		if strings.HasPrefix(name, ".") {
			continue
		}
		list = append(list, name)
	}
	return list, nil
}

// створює директорії якщо їх нема
func EnsureDirs(dirs ...string) error {
	for _, d := range dirs {
		if d == "" {
			continue
		}
		if err := os.MkdirAll(d, 0o755); err != nil {
			return fmt.Errorf("не вдалося створити теку %s: %w", d, err)
		}
	}
	return nil
}

// повертає шлях у outputDir для файлу inputPath з розширенням .pdf
func DefaultOutputPath(outputDir, inputPath string) string {
	base := filepath.Base(inputPath)
	name := strings.TrimSuffix(base, filepath.Ext(base))
	return filepath.Join(outputDir, name+".pdf")
}

// формує унікальне ім'я у outputDir на випадок конфлікту
func UniqueOutputForImages(outputDir string, inputs []string) string {
	if len(inputs) == 0 {
		return filepath.Join(outputDir, "images_combined.pdf")
	}
	first := filepath.Base(inputs[0])
	name := strings.TrimSuffix(first, filepath.Ext(first))
	candidate := filepath.Join(outputDir, name+"_combined.pdf")
	if !FileExists(candidate) {
		return candidate
	}
	for i := 1; ; i++ {
		c := filepath.Join(outputDir, fmt.Sprintf("%s_combined_%d.pdf", name, i))
		if !FileExists(c) {
			return c
		}
	}
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
