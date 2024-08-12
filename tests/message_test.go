package tests

import (
	"encoding/json"
	"reflect"
	"testing"
	"wave-message/internal/service"
	"wave-message/pkg/message"
)

// TestCreateMessage_Success tests the successful creation of a message
func TestCreateMessage_Success(t *testing.T) {
	svc := service.NewMessageService()

	channelID := "test-channel"
	senderID := "test-sender"
	payload := "Hello, World!"
	isPersistent := true

	msg, err := svc.CreateMessage(channelID, senderID, payload, isPersistent)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if msg.ChannelId != channelID {
		t.Errorf("expected channelID %v, got %v", channelID, msg.ChannelId)
	}

	if msg.SenderId != senderID {
		t.Errorf("expected senderID %v, got %v", senderID, msg.SenderId)
	}

	if string(msg.Payload) != payload {
		t.Errorf("expected payload %v, got %v", payload, string(msg.Payload))
	}

	if !msg.IsPersistent {
		t.Errorf("expected isPersistent to be true, got false")
	}
}

// TestCreateMessage_InvalidParams tests the creation of a message with invalid parameters
func TestCreateMessage_InvalidParams(t *testing.T) {
	svc := service.NewMessageService()

	_, err := svc.CreateMessage("", "test-sender", "Hello, World!", true)
	if err == nil {
		t.Fatal("expected error for missing channelID, got none")
	}

	_, err = svc.CreateMessage("test-channel", "", "Hello, World!", true)
	if err == nil {
		t.Fatal("expected error for missing senderID, got none")
	}

	_, err = svc.CreateMessage("test-channel", "test-sender", nil, true)
	if err == nil {
		t.Fatal("expected error for nil payload, got none")
	}
}

// TestValidateMessage tests the validation of a message
func TestValidateMessage(t *testing.T) {
	svc := service.NewMessageService()

	msg := &message.Message{
		ChannelId:    "test-channel",
		SenderId:     "test-sender",
		Payload:      []byte("Hello, World!"),
		Timestamp:    1234567890,
		MessageId:    "test-message-id",
		IsPersistent: true,
	}

	err := svc.ValidateMessage(msg)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test for missing ChannelId
	msg.ChannelId = ""
	err = svc.ValidateMessage(msg)
	if err == nil {
		t.Fatal("expected error for missing channelID, got none")
	}

	// Reset ChannelId and test for missing SenderId
	msg.ChannelId = "test-channel"
	msg.SenderId = ""
	err = svc.ValidateMessage(msg)
	if err == nil {
		t.Fatal("expected error for missing senderID, got none")
	}

	// Reset SenderId and test for empty payload
	msg.SenderId = "test-sender"
	msg.Payload = []byte("")
	err = svc.ValidateMessage(msg)
	if err == nil {
		t.Fatal("expected error for empty payload, got none")
	}
}

// TestSerializePayload tests the serialization of different payloads
func TestSerializePayload(t *testing.T) {
	// Test with a string
	strPayload := "Hello, World!"
	bytes, err := service.SerializePayload(strPayload)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !reflect.DeepEqual(bytes, []byte(strPayload)) {
		t.Errorf("expected %v, got %v", strPayload, string(bytes))
	}

	// Test with a map
	mapPayload := map[string]interface{}{
		"key": "value",
	}
	bytes, err = service.SerializePayload(mapPayload)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expectedBytes, _ := json.Marshal(mapPayload)
	if !reflect.DeepEqual(bytes, expectedBytes) {
		t.Errorf("expected %v, got %v", string(expectedBytes), string(bytes))
	}
}

// TestDeserializePayload tests the deserialization of different payloads
func TestDeserializePayload(t *testing.T) {
	// Test with a string
	strPayload := "Hello, World!"
	data := []byte(strPayload)
	result, err := service.DeserializePayload(data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != strPayload {
		t.Errorf("expected %v, got %v", strPayload, result)
	}

	// Test with a JSON map
	mapPayload := map[string]interface{}{
		"key": "value",
	}
	data, _ = json.Marshal(mapPayload)
	result, err = service.DeserializePayload(data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mapPayload) {
		t.Errorf("expected %v, got %v", mapPayload, result)
	}
}

// TestSerializeData tests the serialization of a Message struct or string
func TestSerializeData(t *testing.T) {
	// Test with a Message struct
	msg := message.Message{
		ChannelId:    "test-channel",
		SenderId:     "test-sender",
		Payload:      []byte("Hello, World!"),
		Timestamp:    1234567890,
		MessageId:    "test-message-id",
		IsPersistent: true,
	}

	bytes := service.SerializeData(msg)
	expectedBytes, _ := json.Marshal(msg)
	if !reflect.DeepEqual(bytes, expectedBytes) {
		t.Errorf("expected %v, got %v", string(expectedBytes), string(bytes))
	}

	// Test with a string
	str := "Hello, World!"
	bytes = service.SerializeData(str)
	if !reflect.DeepEqual(bytes, []byte(str)) {
		t.Errorf("expected %v, got %v", str, string(bytes))
	}
}

// TestDeserializeData tests the deserialization of a Message struct or string
func TestDeserializeData(t *testing.T) {
	// Test with a Message struct
	msg := message.Message{
		ChannelId:    "test-channel",
		SenderId:     "test-sender",
		Payload:      []byte("Hello, World!"),
		Timestamp:    1234567890,
		MessageId:    "test-message-id",
		IsPersistent: true,
	}
	data, _ := json.Marshal(msg)
	result := service.DeserializeData(data)

	deserializedMsg, ok := result.(message.Message)
	if !ok {
		t.Fatalf("expected message.Message type, got %T", result)
	}
	if !reflect.DeepEqual(deserializedMsg, msg) {
		t.Errorf("expected %v, got %v", msg, deserializedMsg)
	}

	// Test with a string
	str := "Hello, World!"
	data = []byte(str)
	result = service.DeserializeData(data)
	if result != str {
		t.Errorf("expected %v, got %v", str, result)
	}
}
