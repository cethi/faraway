package tcp_server

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
	"strings"

	proofOfWork "tcp-proof-of-work/internal/proof-of-work"
)

const TaskLength = 10

var ErrIncorrectMessageFormat = errors.New("incorrect message format")

type ProofOfWorkHandler struct {
	// store to keep information about not finished computation tasks
	store ServerStore

	// proof of work function used to validate client responses
	powFunction proofOfWork.POWFunction
}

func NewProofOfWorkHandler(powFunction proofOfWork.POWFunction, store ServerStore) *ProofOfWorkHandler {
	return &ProofOfWorkHandler{
		powFunction: powFunction,
		store:       store,
	}
}

func (h *ProofOfWorkHandler) ServeTCP(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		request, err := r.ReadString('\n')
		switch err {
		case nil:
			response, handleErr := h.handle(strings.TrimSpace(request))
			if handleErr != nil {
				log.Printf("failed to handle client message: %v\n", err)
			}
			if _, err := conn.Write([]byte(response + "\n")); err != nil {
				log.Printf("failed to respond to client: %v\n", err)
			}
		case io.EOF:
			log.Println("connection was closed")
			return
		default:
			log.Printf("error: %v\n", err)
		}
	}
}

// parse and process payload, possible values:
// - hello$<name> say hello to the server and get task for computation
// - result$<name>$<hash> send computation result to the server and get the quote
func (h *ProofOfWorkHandler) handle(payload string) (response string, err error) {
	parts := strings.Split(payload, "$")
	if len(parts) < 1 {
		err = ErrIncorrectMessageFormat
		return
	}

	switch parts[0] {
	case "hello":
		if len(parts) != 2 {
			err = ErrIncorrectMessageFormat
			return
		}
		return h.createTask(parts[1]), nil
	case "result":
		if len(parts) != 3 {
			err = ErrIncorrectMessageFormat
			return
		}

		if h.validate(parts[1], parts[2]) {
			response = "Here is your quote: Believe and act as if it were impossible to fail."
			return
		} else {
			response = "You have provided incorrect solution"
			return
		}
	default:
		err = ErrIncorrectMessageFormat
		return
	}
}

func (h *ProofOfWorkHandler) createTask(clientId string) string {
	task := proofOfWork.GenerateRandomTask(TaskLength)
	h.store.Set(clientId, task)
	log.Printf("[clientId=%s] Created task: %s\n", clientId, task)
	return task
}

func (h *ProofOfWorkHandler) validate(clientId string, solution string) bool {
	log.Printf("[clientId=%s] Validate solution: %s\n", clientId, solution)
	task, ok := h.store.GetAndDelete(clientId)
	if !ok {
		log.Printf("[clientId=%s] Unknown clientId\n", clientId)
		return false
	}

	result, err := h.powFunction.Validate(task, solution)
	if err != nil {
		log.Printf("[clientId=%s] Validation error: %v\n", clientId, err)
		return false
	}

	log.Printf("[clientId=%s] Validation result: %v\n", clientId, result)
	return result
}
