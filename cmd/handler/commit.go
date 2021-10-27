package handler

import (
	"fmt"

	"github.com/josemiguelmelo/commitio/git"
	"github.com/josemiguelmelo/commitio/models"
	"github.com/josemiguelmelo/commitio/prompt"
	"github.com/josemiguelmelo/commitio/prompt/validators"
	"github.com/spf13/cobra"
)

type CommitCommandHandler struct {
	inputPrompt         prompt.Prompt
	commitCommandRunner git.GitCommitCommandRunner
}

func NewCommitCommandHandler(
	commitCommandRunner git.GitCommitCommandRunner,
	inputPrompt prompt.Prompt,
) CommitCommandHandler {
	return CommitCommandHandler{
		inputPrompt:         inputPrompt,
		commitCommandRunner: commitCommandRunner,
	}
}

func (h CommitCommandHandler) Handle(cmd *cobra.Command, args []string) error {
	message := h.inputPrompt.GetInput("Commit message:", validators.NonEmpty)
	body := h.inputPrompt.GetInput("Commit body", validators.NoValidation)
	issueReference := h.inputPrompt.GetInput("Issue ref:", validators.NoValidation)
	changeType := h.inputPrompt.GetSelect("Change Type:", models.ChangeTypeValuesAsString())

	changeTypeEnum, err := models.ChangeTypeFromString(changeType)
	if err != nil {
		return err
	}

	commit := models.Commit{
		Subject: message,
		Body:    &body,
		Issue:   &models.Issue{Reference: issueReference},
		Type:    changeTypeEnum,
	}

	out, cmdErr := h.commitCommandRunner.Run(commit)
	fmt.Println(out)
	return cmdErr
}
