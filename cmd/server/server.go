package main

import (
	"log"
	proofOfWork "tcp-proof-of-work/internal/proof-of-work"
	tcpserver "tcp-proof-of-work/internal/tcp-server"
)

func main() {
	// Choose one of proof of work functions
	powFunction := &proofOfWork.HashCashPOW{}

	// Chose one of stores
	memoryStore := tcpserver.NewMemoryStore()

	// Initialize server handler using selected function
	tcpHandler := tcpserver.NewProofOfWorkHandler(powFunction, memoryStore)

	// Listen and serve clients
	err := tcpserver.ListenAndServe(":3456", tcpHandler)
	if err != nil {
		log.Printf("error starting server: %v\n", err)
	}
}
