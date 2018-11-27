package directmessages

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestItUnmarchalsAEventsJson(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/events.json")

	if err != nil {
		t.Error("Failed to parse attachment.json", err)
	}

	var events Events

	json.Unmarshal(data, &events)

	if events.NextCursor != "SomeCursor" {
		t.Error("NextCursor expected", "SomeCursor", "got", events.NextCursor)
	}

	if len(events.Events) != 1 {
		t.Error("Events expected to be 1 got", len(events.Events))
	}
}
