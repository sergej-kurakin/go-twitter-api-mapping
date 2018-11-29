package directmessages

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestItUnmarchalsMessageDataJson(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/messagedata.json")

	if err != nil {
		t.Error("Failed to parse messagedata.json", err)
	}

	var messageData MessageData

	json.Unmarshal(data, &messageData)

	if messageData.Text != "Message Text" {
		t.Error("Text expected", "Message Text", "got", messageData.Text)
	}

	if messageData.Attachment.Type != "AttachmentType" {
		t.Error("Attachment Type expected", "AttachmentType", "got", messageData.Attachment.Type)
	}
}
