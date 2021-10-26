package models

import (
	"github.com/stretchr/testify/suite"

	"testing"
)

type IssueModelTestSuite struct {
	suite.Suite
}

func TestIssueModelTestSuite(t *testing.T) {
	suite.Run(t, new(IssueModelTestSuite))
}

func (suite *IssueModelTestSuite) TestIssueBuilder() {
	reference := "ref"
	issue := NewIssue(reference)

	suite.Equal(reference, issue.Reference)
}
