package converter

import (
	"bufio"
	"fmt"
	"os"

	gofpdf "github.com/jung-kurt/gofpdf"
)

// читає текстовий файл і генерує PDF з одним стовпцем
func textToPDF(inputPath, outputPath string, overwrite bool) error {
	if !fileExists(inputPath) {
		return fmt.Errorf("вхідний файл не знайдено: %s", inputPath)
	}
	if outputPath == "" {
		return fmt.Errorf("не задано шлях до вихідного PDF")
	}

	// Обробка перезапису
	if fileExists(outputPath) {
		if overwrite {
			if err := os.Remove(outputPath); err != nil {
				return fmt.Errorf("не вдалося перезаписати файл %s: %w", outputPath, err)
			}
		} else {
			return fmt.Errorf("файл %s вже існує (вимкнений overwrite)", outputPath)
		}
	}

	//відкриваємо текстовий файл для читання
	f, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("не вдалось відкрити вхідний файл: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetTitle("Converted from text", false)
	pdf.AddPage()
	// Використовуємо базовий шрифт
	pdf.SetFont("Arial", "", 12)

	// Параметри для переносу рядків
	leftMargin, _, _, _ := pdf.GetMargins()
	pageW, _ := pdf.GetPageSize()
	usableWidth := pageW - leftMargin - leftMargin //відступи
	lineHt := 6.0

	// Встановлюємо початковий X
	pdf.SetLeftMargin(leftMargin)

	for scanner.Scan() {
		text := scanner.Text()
		pdf.SetX(leftMargin)
		pdf.MultiCell(usableWidth, lineHt, text, "", "L", false) //MultiCell для автоматичного переносу по ширині
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("помилка при читанні тексту: %w", err)
	}

	//переконуємось, що
	if err := ensureOutputDir(outputPath); err != nil {
		return err
	}

	//Записати PDF
	if err := pdf.OutputFileAndClose(outputPath); err != nil {
		return fmt.Errorf("не вдалось записати PDF: %w", err)
	}
	return nil
}
