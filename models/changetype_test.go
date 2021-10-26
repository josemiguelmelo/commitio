package models

import (
	"github.com/stretchr/testify/suite"

	"testing"
)

type ChangeTypeModelTestSuite struct {
	suite.Suite
}

func TestChangeTypeModelTestSuite(t *testing.T) {
	suite.Run(t, new(ChangeTypeModelTestSuite))
}

func (suite *ChangeTypeModelTestSuite) TestChangeTypeFromStringRepresentation() {
	expected := FeatureType
	result, err := ChangeTypeFromString("feature")
	suite.Nil(err)
	suite.Equal(expected, *result)
}

func (suite *ChangeTypeModelTestSuite) TestChangeTypeFromStringRepresentationWithNonValid() {
	result, err := ChangeTypeFromString("non-existing")
	suite.NotNil(err)
	suite.Nil(result)
}

func (suite *ChangeTypeModelTestSuite) TestChangeTypeEnumsListAsString() {
	expected := []string{
		"feature", "fix", "docs", "test",
	}
	result := ChangeTypeValuesAsString()
	suite.ContainsAll(expected, result)
}

func (suite *ChangeTypeModelTestSuite) ContainsAll(expected []string, values []string) {
	for _, v := range values {
		suite.Contains(expected, v)
	}
}
