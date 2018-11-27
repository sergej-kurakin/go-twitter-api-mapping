package directmessages

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestItUnmarchalsAttachmentJson(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/attachment.json")

	if err != nil {
		t.Error("Failed to parse attachment.json", err)
	}

	var attachment Attachment

	json.Unmarshal(data, &attachment)

	if attachment.Type != "AttachmentType" {
		t.Error("Type expected", "AttachmentType", "got", attachment.Type)
	}
}
