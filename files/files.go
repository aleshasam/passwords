package files

import (
	"fmt"
	"os"
)

func Write(content []byte, path string) {
	file, createFileError := os.Create(path)
	if createFileError != nil {
		fmt.Println("Ошибка: ", createFileError.Error())
		return
	}
	_, writeError := file.Write(content)
	defer file.Close()

	if writeError != nil {
		fmt.Println("Ошибка: ", writeError.Error())
		return
	}
	fmt.Println("Запись успешна")
}

func Read(path string) ([]byte, error) {
	bytes, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}
