package mocks

import (
	"github.com/josemiguelmelo/commitio/bash"
	"github.com/stretchr/testify/mock"
)

type BashCommandMock struct {
	mock.Mock
}

func (cmd *BashCommandMock) Exec(command string) (string, *bash.CommandError) {
	args := cmd.Called(command)
	if args.Get(1) == nil {
		return args.String(0), nil
	}
	return args.String(0), args.Get(1).(*bash.CommandError)
}
