package helpers

import (
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
		Label:       ic.Label,
		Templates:   templates,
		Validate:    ic.Validate,
		HideEntered: false,
	}

	result, err := prompt.Run()

	CheckError(err)

	return result
}
