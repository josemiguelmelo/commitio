package cmd

import (
	"github.com/josemiguelmelo/commitio/bash"
	"github.com/josemiguelmelo/commitio/cmd/handler"
	"github.com/josemiguelmelo/commitio/git"
	"github.com/josemiguelmelo/commitio/prompt"
	"github.com/spf13/cobra"
)

var (
	commitCommandRunner = git.NewGitCommitCommandRunner(bash.NewBashCommand())
	commitCmdHandler    = handler.NewCommitCommandHandler(commitCommandRunner, prompt.PromptImpl{})
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create commit",
	Long:  `Create a git commit with a standard message`,
	RunE:  commitCmdHandler.Handle,
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
