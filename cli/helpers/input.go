package helpers

import (
	"github.com/manifoldco/promptui"
)

type InputContentString struct {
	Label    string
	Validate func(string) error
}

type InputContentSelect struct {
	Label string
	Items []InputContentSelectItem
}

type InputContentSelectItem struct {
	Name   string
	Detail string
}

func InputString(ic InputContentString) string {

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

func InputSelect(ic InputContentSelect) int {

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }} ",
		Selected: "\U0001F336 {{ .Name | red | cyan }}",
		Details:  `{{ .Detail }}`,
	}

	prompt := promptui.Select{
		Label:     ic.Label,
		Items:     ic.Items,
		Templates: templates,
	}

	key, _, err := prompt.Run()

	CheckError(err)

	return key
}
