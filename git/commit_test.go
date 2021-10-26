package git

import (
	"github.com/josemiguelmelo/commitio/bash"
	"github.com/josemiguelmelo/commitio/git/mocks"
	"github.com/josemiguelmelo/commitio/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"errors"
	"testing"
)

type CommitRunnerTestSuite struct {
	suite.Suite
	BashCommandRunner *mocks.BashCommandMock
	CommitRunner      GitCommitCommandRunner
}

var commitType = models.FeatureType

func (suite *CommitRunnerTestSuite) SetupTest() {
	suite.BashCommandRunner = &mocks.BashCommandMock{}
	suite.CommitRunner = NewGitCommitCommandRunner(suite.BashCommandRunner)
}

func (suite *CommitRunnerTestSuite) TestGenerateCommandWithoutBody() {
	commit := models.NewCommit("subject", nil, nil, &commitType)
	command := suite.CommitRunner.Command(commit)
	expected := `git commit -m "feature: subject"`

	suite.Equal(expected, command)
}

func (suite *CommitRunnerTestSuite) TestGenerateCommandWithBody() {
	commitBody := "body test"
	commit := models.NewCommit("subject", &commitBody, nil, &commitType)

	command := suite.CommitRunner.Command(commit)
	expected := `git commit -m "feature: subject" -m "body test"`

	suite.Equal(expected, command)
}

func (suite *CommitRunnerTestSuite) TestGenerateCommandWithIssue() {
	commitBody := "body test"
	issueRef := "ISS-123"
	commitIssue := models.NewIssue(issueRef)
	commit := models.NewCommit("subject", &commitBody, &commitIssue, &commitType)

	command := suite.CommitRunner.Command(commit)

	expected := `git commit -m "feature: subject ISS-123" -m "body test"`

	suite.Equal(expected, command)
}

func (suite *CommitRunnerTestSuite) TestRunWithSuccess() {
	output := "commit done"
	suite.BashCommandRunner.On("Exec", mock.Anything).Return(output, nil)

	commit := models.NewCommit("subject", nil, nil, nil)

	result, err := suite.CommitRunner.Run(commit)

	suite.Equal(output, result)
	suite.Nil(err)
}

func (suite *CommitRunnerTestSuite) TestRunWithError() {
	suite.BashCommandRunner.On("Exec", mock.Anything).Return("", &bash.CommandError{
		Err: errors.New("Error executing command"),
		Msg: "Failed to run command. Exit(1)",
	})

	commit := models.NewCommit("subject", nil, nil, nil)

	result, err := suite.CommitRunner.Run(commit)

	suite.Equal("", result)
	suite.NotNil(err)
}

func TestCommitRunnerTestSuite(t *testing.T) {
	suite.Run(t, new(CommitRunnerTestSuite))
}
