package files

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

type JsonDB struct {
	filename string
}

func NewJsonDB(name string) *JsonDB {
	return &JsonDB{
		filename: name,
	}
}

func (db JsonDB) Read() ([]byte, error) {
	bytes, readErr := os.ReadFile(db.filename)

	if readErr != nil {
		return nil, readErr
	}

	return bytes, nil
}

func (db JsonDB) Write(content []byte) {
	file, createFileError := os.Create(db.filename)
	if createFileError != nil {
		fmt.Println("Ошибка: ", createFileError.Error())
		return
	}
	_, writeError := file.Write(content)
	defer file.Close()

	if writeError != nil {
		color.Red("Ошибка: ", writeError.Error())
		return
	}
	color.Green("Запись успешна")
}
