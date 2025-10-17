package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var content strings.Builder

	// Рекурсивный обход директории
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Проверка расширения .go
		if !d.IsDir() && strings.HasSuffix(path, ".go") {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Добавление разделителя с именем файла
			content.WriteString(fmt.Sprintf("\n%s\n%s\n", strings.Repeat("=", 40), path))
			content.Write(data)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Ошибка обхода директорий: %v\n", err)
		return
	}

	// Запись в файл
	if err := os.WriteFile("prompt.txt", []byte(content.String()), 0644); err != nil {
		fmt.Printf("Ошибка записи: %v\n", err)
		return
	}

	fmt.Println("Успешно создан prompt.txt")
}