package directmessages

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestItUnmarchalsMessageCreateJson(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/messagecreate.json")

	if err != nil {
		t.Error("Failed to parse messagecreate.json", err)
	}

	var messageCreate MessageCreate

	json.Unmarshal(data, &messageCreate)

	if messageCreate.SenderID != "SenderId" {
		t.Error("SenderID expected", "SenderId", "got", messageCreate.SenderID)
	}

	if messageCreate.Target.RecipientID != "RecipientId" {
		t.Error("Target RecipientID expected", "RecipientId", "got", messageCreate.Target.RecipientID)
	}

	if messageCreate.MessageData.Text != "Some Text" {
		t.Error("MessageData Text expected", "Some Text", "got", messageCreate.MessageData.Text)
	}
}
