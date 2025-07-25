package internal

import (
	"errors"

	"github.com/manifoldco/promptui"
)

// This function presents a list of commit types for the user to choose from
// and returns the selected commit type.
func InputCommitType() string {
	typeTemplate := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "{{ .Label | cyan }}",
		Inactive: "{{ .Label }}",
	}

	typePrompt := promptui.Select{
		Label:     "Select the type of change you are committing",
		Items:     COMMIT_TYPES,
		Templates: typeTemplate,
	}

	index, _, err := typePrompt.Run()
	AbortOnError(err, "failed to select commit type. Please try again.")

	return COMMIT_TYPES[index].Type
}

// This function prompts the user to enter a scope for the commit
// and returns the entered scope value.
func InputScope() string {
	scopePrompt := promptui.Prompt{
		Label: "Scope (e.g., auth, db, api, ui, cli) - leave empty for none",
	}

	scope, err := scopePrompt.Run()
	AbortOnError(err, "failed to get a commit message. Please try again.")

	return scope
}

// This function prompts the user to enter a commit message
// and returns the entered commit message.
func InputMessage() string {
	messagePrompt := promptui.Prompt{
		Label: "Enter short commit description",
		Validate: func(s string) error {
			if len(s) < 10 {
				return errors.New("please enter at least 10 characters")
			}

			return nil
		},
	}

	message, err := messagePrompt.Run()
	AbortOnError(err, "failed to get the commit message. Please try again.")

	return message
}

// This function prompts the user to enter a commit body
// and returns the entered commit body.
func InputBody() string {
	bodyPrompt := promptui.Prompt{
		Label:     "Enter detailed commit message (optional)",
		IsVimMode: true,
	}

	body, err := bodyPrompt.Run()
	AbortOnError(err, "failed to get the commit body. Please try again.")

	return body
}

// This function prompts the user to confirm something
func InputConfirm(label string) bool {
	confirmPrompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	_, err := confirmPrompt.Run()
	return err == nil
}

func InputString(label string) string {
	prompt := promptui.Prompt{
		Label: label,
	}

	input, err := prompt.Run()
	AbortOnError(err, "failed to get the input. Please try again.")

	return input
}
