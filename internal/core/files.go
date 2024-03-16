package core

import (
	"fmt"
	"io"
	"os"
)

func OpenFiles(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла: ", err)
	}
	defer file.Close()

	// Читаем содержимое файла
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return nil
	}
	byteArray := []byte(data)

	return byteArray
}
