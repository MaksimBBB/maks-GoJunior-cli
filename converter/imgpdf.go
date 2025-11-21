package converter

import (
	"fmt"
	"os"
	"path/filepath"

	gofpdf "github.com/jung-kurt/gofpdf"
)

// воно додає кожне зображення як окрему сторінку в PDF
func imagesToPDF(inputPaths []string, outputPath string, overwrite bool) error {
	//перевірка вхідних файлів
	if len(inputPaths) == 0 {
		return fmt.Errorf("список вхідних файлів порожній")
	}
	for _, p := range inputPaths {
		if !fileExists(p) {
			return fmt.Errorf("вхідний файл не знайдено: %s", p)
		}
	}

	//якщо outputPath не вказаний то помилка
	if outputPath == "" {
		return fmt.Errorf("не задано шлях до вихідного PDF")
	}

	//обробка перезапису
	if fileExists(outputPath) {
		if overwrite {
			if err := os.Remove(outputPath); err != nil {
				return fmt.Errorf("не вдалося перезаписати файл %s: %w", outputPath, err)
			}
		} else {
			return fmt.Errorf("файл %s вже існує (вимкнений overwrite)", outputPath)
		}
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetTitle(filepath.Base(outputPath), false)

	// параметри сторінки
	left, _, right, _ := pdf.GetMargins()
	pageW, _ := pdf.GetPageSize()
	maxWidth := pageW - left - right
	const padding = 10.0
	imgWidth := maxWidth - padding

	for _, ip := range inputPaths {
		pdf.AddPage()

		opt := gofpdf.ImageOptions{
			ImageType:             "",
			ReadDpi:               true,
			AllowNegativePosition: false,
		}

		// Вставка зображення
		pdf.ImageOptions(ip, padding/2, padding/2, imgWidth, 0, false, opt, 0, "")
	}

	//створює директорію для output(якщо потрібно)
	if err := ensureOutputDir(outputPath); err != nil {
		return err
	}

	//записує PDF
	if err := pdf.OutputFileAndClose(outputPath); err != nil {
		return fmt.Errorf("не вдалось записати PDF: %w", err)
	}
	return nil
}
