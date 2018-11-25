package entities

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

var initialUserMentions = UserMentions{
	ScreenName: "Screen Name",
	Name:       "Just Name",
	ID:         "Id as string",
}

func TestItUnmarshalsUserMentionsJson(t *testing.T) {

	data, err := ioutil.ReadFile("testdata/usermentions.json")

	if err != nil {
		panic(err)
	}

	var userMentions UserMentions

	json.Unmarshal(data, &userMentions)

	assertUserMentions(t, initialUserMentions, userMentions)
}

func TestItMarshalsUserMentionsJson(t *testing.T) {

	body, err := json.Marshal(initialUserMentions)

	tmpfile, err := ioutil.TempFile("testdata/", "usermentions_*.json")
	tmpName := tmpfile.Name()
	defer os.Remove(tmpName)

	tmpfile.Write(body)

	tmpfile.Close()

	data, err := ioutil.ReadFile(tmpName)

	if err != nil {
		panic(err)
	}

	var userMentions UserMentions

	json.Unmarshal(data, &userMentions)

	assertUserMentions(t, initialUserMentions, userMentions)
}

func assertUserMentions(t *testing.T, initialUserMention UserMentions, usermentions UserMentions) {
	if usermentions.ScreenName != initialUserMention.ScreenName {
		t.Error("ScreenName expected", initialUserMentions.ScreenName, "got", usermentions.ScreenName)
	}

	if usermentions.Name != initialUserMention.Name {
		t.Error("Name expected", initialUserMentions.Name, "got", usermentions.Name)
	}

	if usermentions.ID != initialUserMention.ID {
		t.Error("ID expected", initialUserMentions.ID, "got", usermentions.ID)
	}
}
