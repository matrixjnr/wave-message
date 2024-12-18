package main

import (
	"fmt"
	"github.com/matrixjnr/wave-message/pkg/service"
)

func main() {
	msgService := service.NewMessageService()

	// Create a message with a string payload
	msg, err := msgService.CreateMessage("channel1", "sender1", "This is a simple string payload", true)
	if err != nil {
		fmt.Printf("Error creating message: %v\n", err)
		return
	}

	// Serialize the message
	serializedMsg := service.SerializeData(*msg)
	fmt.Printf("Serialized Message: %s\n", serializedMsg)

	// Deserialize the message
	deserializedMsg := service.DeserializeData(serializedMsg)
	fmt.Printf("Deserialized Message: %+v\n", deserializedMsg)

	// Extract and deserialize payload
	payload, err := service.DeserializePayload(msg.Payload)
	if err != nil {
		fmt.Printf("Error deserializing payload: %v\n", err)
		return
	}
	fmt.Printf("Deserialized Payload: %+v\n", payload)
}
