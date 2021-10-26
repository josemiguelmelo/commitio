package mocks

import (
	"github.com/manifoldco/promptui"
	"github.com/stretchr/testify/mock"
)

type InputPromptMock struct {
	mock.Mock
}

func (p *InputPromptMock) GetInput(inputLabel string, validate promptui.ValidateFunc) string {
	args := p.Called(inputLabel)
	return args.String(0)
}

func (p *InputPromptMock) GetSelect(inputLabel string, options []string) string {
	args := p.Called(inputLabel, options)
	return args.String(0)
}

func (p *InputPromptMock) GetSelectWithAdd(inputLabel string, options []string) string {
	args := p.Called(inputLabel, options)
	return args.String(0)
}
