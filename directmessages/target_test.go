package directmessages

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

var initialTarget = Target{
	RecipientID: "RecipientId",
}

func TestItUnmarshalsTargetJson(t *testing.T) {

	data, err := ioutil.ReadFile("testdata/target.json")

	if err != nil {
		panic(err)
	}

	var target Target

	json.Unmarshal(data, &target)

	assertTargetSame(t, initialTarget, target)
}

func TestItMarshalsSymbolJson(t *testing.T) {

	body, err := json.Marshal(initialTarget)

	tmpfile, err := ioutil.TempFile("testdata/", "target_*.json")
	tmpName := tmpfile.Name()
	defer os.Remove(tmpName)

	tmpfile.Write(body)

	tmpfile.Close()

	data, err := ioutil.ReadFile(tmpName)

	if err != nil {
		panic(err)
	}

	var target Target

	json.Unmarshal(data, &target)

	assertTargetSame(t, initialTarget, target)
}

func assertTargetSame(t *testing.T, initialTarget Target, target Target) {
	if initialTarget.RecipientID != target.RecipientID {
		t.Error("Target RecipientID expected", initialTarget.RecipientID, "got", target.RecipientID)
	}
}
