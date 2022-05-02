package services

import (
	"os"
)

type fsService struct{}

func NewFsService() fsService {
	return fsService{}
}

func (fs fsService) CreateFolder(dirPath string) error {
	return os.Mkdir(dirPath, os.ModePerm)
}

func (fs fsService) CreateEmptyFile(filePath string) (*os.File, error) {
	f, err := os.Create(filePath)
	return f, err
}

func (fs fsService) CreateFileWithContent(filePath, content string) error {
	err := os.WriteFile(filePath, []byte(content), os.ModePerm)

	return err
}

func (fs fsService) WriteToFile(filePath, content string) error {
	err := os.WriteFile(filePath, []byte(content), os.ModePerm)

	return err

}
