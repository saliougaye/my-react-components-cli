package helpers

import (
	"github.com/saliougaye/my-react-components/cli_types"
)

func GetFileInput(filename, content string) cli_types.FileInput {
	return cli_types.FileInput{
		Filepath: filename,
		Content:  content,
	}
}

func GetFileInputFromMap(files map[string]string) []cli_types.FileInput {

	list := []cli_types.FileInput{}

	for k, v := range files {

		list = append(list, cli_types.FileInput{
			Filepath: k,
			Content:  v,
		})
	}

	return list
}
