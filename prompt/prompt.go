package prompt

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

type Prompt interface {
	GetInput(inputLabel string, validate promptui.ValidateFunc) string
	GetSelect(inputLabel string, options []string) string
}

type PromptImpl struct{}

func (p PromptImpl) GetInput(inputLabel string, validate promptui.ValidateFunc) string {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     inputLabel,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func (p PromptImpl) GetSelect(inputLabel string, options []string) string {
	var result string
	var err error

	prompt := promptui.Select{
		Label: inputLabel,
		Items: options,
	}

	_, result, err = prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func (p PromptImpl) GetSelectWithAdd(inputLabel string, options []string) string {
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    inputLabel,
			Items:    options,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			options = append(options, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
