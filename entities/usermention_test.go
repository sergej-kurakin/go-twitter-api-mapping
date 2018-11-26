package entities

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

var initialUserMention = UserMention{
	ScreenName: "Screen Name",
	Name:       "Just Name",
	ID:         "Id as string",
}

func TestItUnmarshalsUserMentionJson(t *testing.T) {

	data, err := ioutil.ReadFile("testdata/usermention.json")

	if err != nil {
		panic(err)
	}

	var userMention UserMention

	json.Unmarshal(data, &userMention)

	assertUserMention(t, initialUserMention, userMention)
}

func TestItMarshalsUserMentionJson(t *testing.T) {

	body, err := json.Marshal(initialUserMention)

	tmpfile, err := ioutil.TempFile("testdata/", "usermention_*.json")
	tmpName := tmpfile.Name()
	defer os.Remove(tmpName)

	tmpfile.Write(body)

	tmpfile.Close()

	data, err := ioutil.ReadFile(tmpName)

	if err != nil {
		panic(err)
	}

	var userMention UserMention

	json.Unmarshal(data, &userMention)

	assertUserMention(t, initialUserMention, userMention)
}

func assertUserMention(t *testing.T, initialUserMention UserMention, userMention UserMention) {
	if userMention.ScreenName != initialUserMention.ScreenName {
		t.Error("ScreenName expected", initialUserMention.ScreenName, "got", userMention.ScreenName)
	}

	if userMention.Name != initialUserMention.Name {
		t.Error("Name expected", initialUserMention.Name, "got", userMention.Name)
	}

	if userMention.ID != initialUserMention.ID {
		t.Error("ID expected", initialUserMention.ID, "got", userMention.ID)
	}
}
