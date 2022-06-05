package main

import (
	"fmt"
	"log"

	proofOfWork "tcp-proof-of-work/internal/proof-of-work"
	tcpclient "tcp-proof-of-work/internal/tcp-client"
)

func StartConversation(client *tcpclient.Client, clientId string, powFunction proofOfWork.POWFunction) {
	// Say hello to the server and get task to calculate
	err := client.SendMessage("hello$" + clientId)
	if err != nil {
		log.Printf("failed to send the client request: %v\n", err)
		return
	}
	task, err := client.GetMessage()
	if err != nil {
		log.Printf("server error: %v\n", err)
		return
	}
	log.Printf("Task was received from server: %s", task)

	// Compute hash
	hash, err := powFunction.Compute(task)
	if err != nil {
		log.Printf("computation hash error: %v\n", err)
		return
	}
	log.Printf("Hash was calculated: %s", hash)

	// Send calculated hash to the server and get quote
	err = client.SendMessage(fmt.Sprintf("result$%s$%s", clientId, hash))
	if err != nil {
		log.Printf("failed to send the client request: %v\n", err)
		return
	}
	response, err := client.GetMessage()
	if err != nil {
		log.Printf("server error: %v\n", err)
		return
	}
	log.Printf("You have got the quote: %s\n", response)
}

func main() {
	// Choose one of proof of work functions
	powFunction := &proofOfWork.HashCashPOW{}

	// Initialize client
	client, err := tcpclient.NewClient("client:3456")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Start conversation with server as Alice using proof of work function
	StartConversation(client, "Alice", powFunction)
}
