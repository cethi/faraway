package tcp_server

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MemoryStoreTestSuite struct {
	suite.Suite
}

func (suite *MemoryStoreTestSuite) TestSetGet() {
	store := NewMemoryStore()
	clientId := "client_1"
	expectedTaskName := "task"
	store.Set(clientId, expectedTaskName)

	actualTaskName, success := store.GetAndDelete(clientId)
	assert.True(suite.T(), success)
	assert.Equal(suite.T(), expectedTaskName, actualTaskName)

	_, success = store.GetAndDelete(clientId)
	assert.False(suite.T(), success)
}

func TestMemoryStoreTestSuite(t *testing.T) {
	suite.Run(t, new(MemoryStoreTestSuite))
}
