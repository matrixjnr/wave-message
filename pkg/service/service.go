package service

import (
	"encoding/json"
	"errors"
	"fmt"
	_ "log"

	"github.com/matrixjnr/wave-message/pkg/message"
)

// MessageService provides functionalities related to Message handling
type MessageService struct{}

// NewMessageService creates a new instance of MessageService
func NewMessageService() *MessageService {
	return &MessageService{}
}

// CreateMessage creates a new message with the given parameters
func (svc *MessageService) CreateMessage(channelID, senderID string, payload interface{}, isPersistent bool) (*message.Message, error) {
	if channelID == "" {
		return nil, errors.New("CreateMessage failed: channelID cannot be empty")
	}
	if senderID == "" {
		return nil, errors.New("CreateMessage failed: senderID cannot be empty")
	}
	if payload == nil {
		return nil, errors.New("CreateMessage failed: payload cannot be nil")
	}

	payloadBytes, err := svc.SerializePayload(payload)
	if err != nil {
		return nil, fmt.Errorf("CreateMessage failed to serialize payload: %w", err)
	}

	return &message.Message{
		ChannelId:    channelID,
		SenderId:     senderID,
		Payload:      payloadBytes,
		IsPersistent: isPersistent,
	}, nil
}

// ValidateMessage validates the given message
func (svc *MessageService) ValidateMessage(msg *message.Message) error {
	if msg.ChannelId == "" {
		return errors.New("ValidateMessage failed: ChannelId is missing")
	}
	if msg.SenderId == "" {
		return errors.New("ValidateMessage failed: SenderId is missing")
	}
	if len(msg.Payload) == 0 {
		return errors.New("ValidateMessage failed: Payload is empty")
	}
	return nil
}

// SerializePayload converts an interface{} to bytes
func (svc *MessageService) SerializePayload(payload interface{}) ([]byte, error) {
	switch v := payload.(type) {
	case string:
		return []byte(v), nil
	case []byte:
		return v, nil
	default:
		return json.Marshal(v)
	}
}

// DeserializePayload converts bytes back into a string or a map
func (svc *MessageService) DeserializePayload(data []byte) (interface{}, error) {
	var obj map[string]interface{}
	if err := json.Unmarshal(data, &obj); err == nil {
		return obj, nil
	}
	return string(data), nil
}

// SerializeData serializes a Message into bytes
func (svc *MessageService) SerializeData(msg *message.Message) ([]byte, error) {
	return json.Marshal(msg)
}

// DeserializeData deserializes bytes into a Message
func (svc *MessageService) DeserializeData(data []byte) (*message.Message, error) {
	var msg message.Message
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("DeserializeData failed: %w", err)
	}
	return &msg, nil
}
