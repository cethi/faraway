package proof_of_work

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HashCashTestSuite struct {
	suite.Suite
	ProofOfWorkFunction HashCashPOW
}

func (suite *HashCashTestSuite) SetupTest() {
	suite.ProofOfWorkFunction = HashCashPOW{}
}

func (suite *HashCashTestSuite) TestComputeHash() {
	cases := []struct {
		task string
	}{
		{"test"},
		{"anotherValue"},
	}
	for _, c := range cases {
		_, err := suite.ProofOfWorkFunction.Compute(c.task)
		assert.NoError(suite.T(), err)
	}
}

func (suite *HashCashTestSuite) TestDifferentTasksDifferentHashes() {
	result1, err := suite.ProofOfWorkFunction.Compute("one")
	assert.NoError(suite.T(), err)
	result2, err := suite.ProofOfWorkFunction.Compute("another")
	assert.NoError(suite.T(), err)

	assert.NotEqual(suite.T(), result1, result2)
}

func (suite *HashCashTestSuite) TestValidate() {
	cases := []struct {
		hash           string
		expectedResult bool
	}{
		{"1:20:220531211659:one::tJKE0ti3rvw=:NjAzNDQw", true},
		{"1:20:220531212013:another::DqO5mkzqAmI=:MjI0ODQ=", false},
	}
	for _, c := range cases {
		result, err := suite.ProofOfWorkFunction.Validate("one", c.hash)
		if c.expectedResult {
			assert.NoError(suite.T(), err)
		} else {
			assert.Error(suite.T(), err)
		}

		assert.Equal(suite.T(), result, c.expectedResult)
	}
}

func (suite *HashCashTestSuite) TestComputeAndValidate() {
	result, err := suite.ProofOfWorkFunction.Compute("task")
	assert.NoError(suite.T(), err)
	validationResult, err := suite.ProofOfWorkFunction.Validate("task", result)
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), validationResult)
}

func TestHashCashTestSuite(t *testing.T) {
	suite.Run(t, new(HashCashTestSuite))
}
