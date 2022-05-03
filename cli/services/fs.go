package services

import (
	"os"

	"github.com/saliougaye/my-react-components/cli_types"
	"github.com/saliougaye/my-react-components/helpers"
)

type fsService struct{}

func NewFsService() fsService {
	return fsService{}
}

func (fs fsService) CreateFolder(dirPath string) error {
	return os.Mkdir(dirPath, os.ModePerm)
}

func (fs fsService) CreateFolders(dirsPath []string) {

	for _, v := range dirsPath {
		err := fs.CreateFolder(v)

		helpers.CheckError(err)
	}
}

func (fs fsService) CreateEmptyFile(filePath string) (*os.File, error) {
	f, err := os.Create(filePath)
	return f, err
}

func (fs fsService) CreateFileWithContent(file cli_types.FileInput) error {
	err := os.WriteFile(file.Filepath, []byte(file.Content), os.ModePerm)

	return err
}

func (fs fsService) CreateFilesWithContent(files []cli_types.FileInput) {

	for _, v := range files {
		err := fs.CreateFileWithContent(v)

		helpers.CheckError(err)
	}

}

func (fs fsService) WriteToFile(filePath, content string) error {
	err := os.WriteFile(filePath, []byte(content), os.ModePerm)

	return err

}
