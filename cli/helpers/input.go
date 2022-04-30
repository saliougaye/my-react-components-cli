package helpers

import (
	"os"

	"github.com/manifoldco/promptui"
)

type InputContent struct {
	Label    string
	Validate func(string) error
}

func InputString(ic InputContent) string {

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
	}

	prompt := promptui.Prompt{
		Label:     ic.Label,
		Templates: templates,
		Validate:  ic.Validate,
	}

	result, err := prompt.Run()

	if err != nil {
		PrintError(err)
		os.Exit(1)
	}
	return result
}
