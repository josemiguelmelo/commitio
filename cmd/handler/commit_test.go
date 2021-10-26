package handler

import (
	"github.com/josemiguelmelo/commitio/cmd/handler/mocks"
	"github.com/josemiguelmelo/commitio/git"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"testing"
)

type CommitCmdHandlerTestSuite struct {
	suite.Suite
	inputPrompt      *mocks.InputPromptMock
	commitCmdHandler CommitCommandHandler
}

func (suite *CommitCmdHandlerTestSuite) SetupTest() {
	bashCommandRunner := mocks.BashCommandMock{}
	gitCommitCommandRunner := git.NewGitCommitCommandRunner(bashCommandRunner)
	suite.inputPrompt = &mocks.InputPromptMock{}
	suite.commitCmdHandler = NewCommitCommandHandler(gitCommitCommandRunner, suite.inputPrompt)
}

func TestCommitCmdHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CommitCmdHandlerTestSuite))
}

func (suite *CommitCmdHandlerTestSuite) TestSuccessfullParametersHandle() {
	suite.inputPrompt.On("GetInput", mock.Anything).Return("test")
	suite.inputPrompt.On("GetSelect", mock.Anything, mock.Anything).Return("feature")

	result := suite.commitCmdHandler.Handle(nil, []string{})
	suite.Nil(result)
}

func (suite *CommitCmdHandlerTestSuite) TestErrorOnInvalidChangeType() {
	suite.inputPrompt.On("GetInput", mock.Anything).Return("test")
	suite.inputPrompt.On("GetSelect", mock.Anything, mock.Anything).Return("invalid")

	result := suite.commitCmdHandler.Handle(nil, []string{})
	suite.NotNil(result)
}
