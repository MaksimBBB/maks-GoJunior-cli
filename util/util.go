package util

import (
	"bufio"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

// обробляє ввід і повертає індекси які треба обробити
func ParseSelection(s string, max int) []int {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	// заміняємо різні роздільники на пробіли
	sep := strings.NewReplacer(",", " ", ";", " ")
	s = sep.Replace(s)
	parts := strings.Fields(s)
	result := []int{}
	seen := map[int]bool{}
	for _, p := range parts {
		if strings.Contains(p, "-") {
			bounds := strings.SplitN(p, "-", 2)
			if len(bounds) != 2 {
				continue
			}
			a, err1 := strconv.Atoi(strings.TrimSpace(bounds[0]))
			b, err2 := strconv.Atoi(strings.TrimSpace(bounds[1]))
			if err1 != nil || err2 != nil {
				continue
			}
			if a < 1 {
				a = 1
			}
			if b > max {
				b = max
			}
			for i := a; i <= b; i++ {
				idx := i - 1
				if idx >= 0 && idx < max && !seen[idx] {
					result = append(result, idx)
					seen[idx] = true
				}
			}
		} else {
			n, err := strconv.Atoi(p)
			if err != nil {
				continue
			}
			if n < 1 || n > max {
				continue
			}
			idx := n - 1
			if !seen[idx] {
				result = append(result, idx)
				seen[idx] = true
			}
		}
	}
	return result
}

func AskYesNoDefault(r *bufio.Reader, prompt string, defaultYes bool) bool {
	fmt.Print(prompt)
	ans, _ := ReadLine(r)
	ans = strings.TrimSpace(strings.ToLower(ans))
	if ans == "" {
		return defaultYes
	}
	if ans == "y" || ans == "yes" || ans == "так" {
		return true
	}
	return false
}

// читає рядок з рідера, повертає без кінців рядка
func ReadLine(r *bufio.Reader) (string, error) {
	str, err := r.ReadString('\n')
	if err != nil {
		return strings.TrimRight(str, "\r\n"), nil
	}
	return strings.TrimRight(str, "\r\n"), nil
}

// розбиває рядок, роздільники це пробіл або кома
func SplitPaths(line string) []string {
	normalized := strings.ReplaceAll(line, ",", " ")
	parts := strings.Fields(normalized)
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		out = append(out, p)
	}
	return out
}

func PausePrompt(r *bufio.Reader) {
	fmt.Print("\nНатисніть Enter, щоб повернутись до меню...")
	_, _ = ReadLine(r)
}

// якщо розширення належить це зображення, то повертає true
func IsImageExt(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".bmp", ".gif":
		return true
	default:
		return false
	}
}

// повертає true якщр це текст
func IsTextExt(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	case ".txt", ".md":
		return true
	default:
		return false
	}
}
