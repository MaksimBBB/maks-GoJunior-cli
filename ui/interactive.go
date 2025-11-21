package ui

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/MaksimBBB/maks-GoJunior-cli/converter"
)

const (
	ConvertFilesDir = "convertFiles"
	OutputDir       = "pdfOutput"
)

// запускає інтерактивний режим
func RunInteractive(conv *converter.Converter) error {
	reader := bufio.NewReader(os.Stdin)

	//перевірка чи існує папка для файлів
	if _, err := os.Stat(ConvertFilesDir); os.IsNotExist(err) {
		fmt.Printf("Папка %s не знайдена. Створюємо...\n", ConvertFilesDir)
		if err := os.MkdirAll(ConvertFilesDir, 0o755); err != nil {
			return fmt.Errorf("не вдалося створити папку %s: %w", ConvertFilesDir, err)
		}
		fmt.Println("Додайте файли для конвертації у цю папку та перезапустіть програму.")
		return nil
	}

	// отримуємо список файлів
	files, err := os.ReadDir(ConvertFilesDir)
	if err != nil {
		return fmt.Errorf("Не вдалося прочитати папку %s: %w", ConvertFilesDir, err)
	}

	if len(files) == 0 {
		fmt.Println("Файли для конвертації відсутні. Додайте їх у", ConvertFilesDir)
		return nil
	}

	// виводимо список файлів
	fmt.Println("У теці файлів для конвертації знаходяться такі файли:")
	for i, f := range files {
		fmt.Printf("%d. %s\n", i+1, f.Name())
	}
	fmt.Print("Оберіть номер файлу для конвертації: ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	var choice int
	_, err = fmt.Sscanf(choiceStr, "%d", &choice)
	if err != nil || choice < 1 || choice > len(files) {
		return fmt.Errorf("невірний вибір")
	}

	selectedFile := files[choice-1].Name()
	inputPath := filepath.Join(ConvertFilesDir, selectedFile)
	outputPath := filepath.Join(OutputDir, strings.TrimSuffix(selectedFile, filepath.Ext(selectedFile))+".pdf")

	// cтворюємо папку для output
	if err := os.MkdirAll(OutputDir, 0o755); err != nil {
		return fmt.Errorf("не вдалося створити папку %s: %w", OutputDir, err)
	}

	// Визначаємо тип конвертації
	ext := strings.ToLower(filepath.Ext(selectedFile))
	switch ext {
	case ".txt":
		fmt.Println("Конвертація текстового файлу у PDF...")
		if err := conv.TextToPDF(inputPath, outputPath, true); err != nil {
			return fmt.Errorf("помилка конвертації: %w", err)
		}
		fmt.Println("Готово! PDF збережено у", outputPath)
	case ".png", ".jpg", ".jpeg":
		fmt.Println("Конвертація зображення у PDF...")
		if err := conv.ImagesToPDF([]string{inputPath}, outputPath, true); err != nil {
			return fmt.Errorf("помилка конвертації: %w", err)
		}
		fmt.Println("Готово! PDF збережено у", outputPath)
	default:
		return fmt.Errorf("невідомий тип файлу: %s", ext)
	}

	return nil
}
