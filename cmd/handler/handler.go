package handler

import (
	"github.com/spf13/cobra"
)

type CmdHandler interface {
	Handle(cmd *cobra.Command, args []string) error
}
