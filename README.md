# PDF Converter CLI

**PDF Converter** — це проста CLI-утиліта для конвертації текстових файлів та зображень у PDF.

---

## Функціонал

- Конвертація **текстових файлів** (`.txt`, `.md`) у PDF
- Конвертація **зображень** (`.jpg`, `.jpeg`, `.png`, `.bmp`, `.gif`) у PDF
- **Інтерактивний режим** у консолі: обираєте файли для конвертації з запропонованого списку
- Автоматичне створення директорій для вихідного PDF
- Можливість перезапису існуючих файлів

---

## Використані бібліотеки

- [gofpdf](https://pkg.go.dev/github.com/jung-kurt/gofpdf) — для генерації PDF
- [cobra](https://github.com/spf13/cobra) — для CLI-команд та аргументів

---

## 💻 Як використовувати

1. Клонуйте репозиторій:
```bash
git clone https://github.com/MaksimBBB/maks-GoJunior-cli.git
cd maks-GoJunior-cli
