package service

import (
	"encoding/json"
	"errors"
	"log"
	"wave-message/pkg/message"
)

// MessageService provides functionalities related to Message handling
type MessageService struct{}

// NewMessageService creates a new instance of MessageService
func NewMessageService() *MessageService {
	return &MessageService{}
}

/*
CreateMessage creates a new message with the given parameters.
Returns a pointer to the created message and an error if the input parameters are invalid.

Parameters:
- channelID: the ID of the channel where the message will be sent
- senderID: the ID of the sender of the message
- payload: the content of the message
- isPersistent: a flag indicating if the message should be stored persistently
*/
func (svc *MessageService) CreateMessage(channelID, senderID string, payload interface{}, isPersistent bool) (*message.Message, error) {
	if channelID == "" {
		return nil, errors.New("channelID cannot be empty")
	}
	if senderID == "" {
		return nil, errors.New("senderID cannot be empty")
	}
	if payload == nil {
		return nil, errors.New("payload cannot be nil")
	}

	// Assuming SerializePayload is a method that converts payload to a byte slice
	payloadBytes, err := SerializePayload(payload)
	if err != nil {
		return nil, err
	}

	// Create the message
	msg := &message.Message{
		ChannelId:    channelID,
		SenderId:     senderID,
		Payload:      payloadBytes,
		IsPersistent: isPersistent,
	}

	return msg, nil
}

/*
ValidateMessage validates the given message.
Returns an error if the message is invalid.

Parameters:
- msg: the message to be validated
*/
func (s *MessageService) ValidateMessage(msg *message.Message) error {
	if msg.ChannelId == "" || msg.SenderId == "" || len(msg.Payload) == 0 {
		return errors.New("invalid message: missing required fields")
	}
	return nil
}

/*
SerializePayload serializes an interface{} to bytes.
If the input is a struct or map, it will serialize it to JSON bytes.
If the input is a string, it will convert it directly to bytes.
*/
func SerializePayload(payload interface{}) ([]byte, error) {
	switch v := payload.(type) {
	case string:
		return []byte(v), nil
	case []byte:
		return v, nil
	default:
		jsonData, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		return jsonData, nil
	}
}

/*
DeserializePayload deserializes bytes into either a struct (if the bytes represent JSON) or a string.

Parameters:
- data: the byte array to deserialize
*/
func DeserializePayload(data []byte) (interface{}, error) {
	var obj map[string]interface{}
	if err := json.Unmarshal(data, &obj); err == nil {
		// If data can be unmarshaled into a map, return the map
		return obj, nil
	}
	// Otherwise, return as a simple string
	return string(data), nil
}

/*
SerializeData serializes a Message struct or string into bytes.

Parameters:
- data: the Message struct or string to serialize
*/
func SerializeData(data interface{}) []byte {
	switch v := data.(type) {
	case message.Message:
		jsonData, err := json.Marshal(v)
		if err != nil {
			log.Fatalf("Failed to serialize Message to JSON: %v", err)
		}
		return jsonData
	case string:
		return []byte(v)
	default:
		log.Fatalf("Unsupported data type for serialization: %T", v)
		return nil
	}
}

/*
DeserializeData deserializes bytes into a Message struct or string.

Parameters:
- data: the byte array to deserialize
*/
func DeserializeData(data []byte) interface{} {
	var msg message.Message
	if err := json.Unmarshal(data, &msg); err == nil {
		return msg
	}
	return string(data)
}
