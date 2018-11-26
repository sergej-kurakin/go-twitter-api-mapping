package entities

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

var initialHashTag = HashTags{
	Text: "Lorum ipsum",
}

func TestItUnmarshalsHashTagJson(t *testing.T) {

	data, err := ioutil.ReadFile("testdata/hashtag.json")

	if err != nil {
		panic(err)
	}

	var hashTags HashTags

	json.Unmarshal(data, &hashTags)

	assertHashTags(t, initialHashTag, hashTags)
}

func TestItMarshalsHashTagJson(t *testing.T) {

	body, err := json.Marshal(initialHashTag)

	tmpfile, err := ioutil.TempFile("testdata/", "hashtag_*.json")
	tmpName := tmpfile.Name()
	defer os.Remove(tmpName)

	tmpfile.Write(body)

	tmpfile.Close()

	data, err := ioutil.ReadFile(tmpName)

	if err != nil {
		panic(err)
	}

	var hashTags HashTags

	json.Unmarshal(data, &hashTags)

	assertHashTags(t, initialHashTag, hashTags)
}

func assertHashTags(t *testing.T, initialHashTag HashTags, hashtag HashTags) {
	if hashtag.Text != initialHashTag.Text {
		t.Error("Text expected", initialHashTag.Text, "got", hashtag.Text)
	}
}
