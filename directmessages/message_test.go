package directmessages

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestItUnmarchalsMessageJson(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/message.json")

	if err != nil {
		t.Error("Failed to parse message.json", err)
	}

	var message Message

	json.Unmarshal(data, &message)

	if message.ID != "3478" {
		t.Error("ID expected", "3478", "got", message.Type)
	}

	if message.Type != "message_type" {
		t.Error("Type expected", "message_type", "got", message.Type)
	}

	if message.CreatedTimestamp != "34775" {
		t.Error("CreatedTimestamp expected", "34775", "got", message.CreatedTimestamp)
	}

	if message.MessageCreate.SenderID != "44577" {
		t.Error("MessageCreate.SenderID expected", "44577", "got", message.MessageCreate.SenderID)
	}
}
