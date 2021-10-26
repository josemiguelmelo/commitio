package mocks

import (
	"github.com/josemiguelmelo/commitio/bash"
)

type BashCommandMock struct{}

func (cmd BashCommandMock) Exec(command string) (string, *bash.CommandError) {
	return "OK", nil
}
