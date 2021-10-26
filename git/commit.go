package git

import (
	"fmt"
	"github.com/josemiguelmelo/commitio/bash"
	"github.com/josemiguelmelo/commitio/models"
)

type GitCommitCommandRunner struct {
	bashCommand bash.BashCommand
}

func NewGitCommitCommandRunner(bashCommand bash.BashCommand) GitCommitCommandRunner {
	return GitCommitCommandRunner{
		bashCommand: bashCommand,
	}
}

func (runner GitCommitCommandRunner) Run(commit models.Commit) (string, error) {
	cmd := runner.Command(commit)
	output, err := runner.bashCommand.Exec(cmd)
	if err != nil {
		return "", fmt.Errorf("Failed running command. %s: %s", err.Msg, err.Err.Error())
	}

	return output, nil
}

func (runner GitCommitCommandRunner) Command(commit models.Commit) string {
	baseCommand := fmt.Sprintf(`git commit -m "%s"`, commit.GetSubject())
	if commit.Body != nil {
		return fmt.Sprintf(`%s -m "%s"`, baseCommand, *commit.Body)
	}
	return baseCommand
}
