package tcp_server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	proofOfWork "tcp-proof-of-work/internal/proof-of-work"
)

type ServerTestSuite struct {
	suite.Suite
}

func (suite *ServerTestSuite) SetupTest() {
	powFunction := &proofOfWork.HashCashPOW{}
	memoryStore := NewMemoryStore()
	tcpHandler := NewProofOfWorkHandler(powFunction, memoryStore)

	go func() {
		err := ListenAndServe("localhost:3456", tcpHandler)
		if err != nil {
			log.Printf("error starting server: %v\n", err)
		}
	}()
}

// Check that the server is up and can accept connections
func (suite *ServerTestSuite) TestServerRun() {
	conn, err := net.Dial("tcp", "localhost:3456")
	if err != nil {
		suite.T().Error("could not connect to server: ", err)
	}
	defer conn.Close()
}

func (suite *ServerTestSuite) TestServerHandleHello() {
	conn, err := net.Dial("tcp", "localhost:3456")
	if err != nil {
		suite.T().Error("could not connect to server: ", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("hello$user\n"))
	assert.NoError(suite.T(), err)

	r := bufio.NewReader(conn)
	response, err := r.ReadString('\n')
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), response)
}

func (suite *ServerTestSuite) TestServerHandleValidation() {
	conn, err := net.Dial("tcp", "localhost:3456")
	if err != nil {
		suite.T().Error("could not connect to server: ", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("hello$user\n"))
	assert.NoError(suite.T(), err)

	r := bufio.NewReader(conn)
	response, err := r.ReadString('\n')
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), response)

	powFunction := &proofOfWork.HashCashPOW{}
	hash, err := powFunction.Compute(response)
	assert.NoError(suite.T(), err)

	_, err = conn.Write([]byte(fmt.Sprintf("result$user$%s\n", hash)))
	assert.NoError(suite.T(), err)

	response, err = r.ReadString('\n')
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), response)
}

// Concurrency tests, stress tests - out of the scope the task

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
