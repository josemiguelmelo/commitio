package models

import (
	"github.com/stretchr/testify/suite"

	"testing"
)

type CommitModelTestSuite struct {
	suite.Suite
}

func TestCommitModelTestSuite(t *testing.T) {
	suite.Run(t, new(CommitModelTestSuite))
}

func (suite *CommitModelTestSuite) TestCommitGetTypeDefaultValue() {
	commit := NewCommit("subject", nil, nil, nil)

	expected := DefaultType
	result := commit.GetTypeOrDefault()
	suite.Equal(expected, result)
}

func (suite *CommitModelTestSuite) TestCommitGetType() {
	expectedType := FeatureType

	commit := NewCommit("subject", nil, nil, &expectedType)

	result := commit.GetTypeOrDefault()
	suite.Equal(expectedType, result)
}

func (suite *CommitModelTestSuite) TestCommitGetSubjectWithoutIssue() {
	changeType := FeatureType
	commit := NewCommit("subject", nil, nil, &changeType)

	expected := `feature: subject`
	result := commit.GetSubject()
	suite.Equal(expected, result)
}

func (suite *CommitModelTestSuite) TestCommitGetSubjectWithIssue() {
	changeType := DocumentationType
	commit := NewCommit("subject", nil, &Issue{Reference: "ISS-123"}, &changeType)

	expected := `docs: subject ISS-123`
	result := commit.GetSubject()
	suite.Equal(expected, result)
}
