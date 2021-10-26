package bash

import (
	"bytes"
	"os/exec"
)

type BashCommand interface {
	Exec(string) (string, *CommandError)
}

type BashCommandImpl struct{}

func NewBashCommand() BashCommandImpl {
	return BashCommandImpl{}
}

func (cmd BashCommandImpl) Exec(command string) (string, *CommandError) {
	bashCmd := exec.Command("sh", "-c", command)

	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	bashCmd.Stdout = &outBuffer
	bashCmd.Stderr = &errBuffer

	err := bashCmd.Run()
	if err != nil {
		return "", &CommandError{
			Err: err,
			Msg: errBuffer.String(),
		}
	}

	return outBuffer.String(), nil
}
