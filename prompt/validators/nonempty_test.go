package validators

import (
	"github.com/stretchr/testify/suite"

	"testing"
)

type NonEmptyValidatorTestSuite struct {
	suite.Suite
}

func TestNonEmptyValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(NonEmptyValidatorTestSuite))
}

func (suite *NonEmptyValidatorTestSuite) TestEmptyReturnsError() {
	err := NonEmpty("")
	suite.NotNil(err)
}

func (suite *NonEmptyValidatorTestSuite) TestNonEmptyReturnsEmptyError() {
	err := NonEmpty("feature")
	suite.Nil(err)
}
