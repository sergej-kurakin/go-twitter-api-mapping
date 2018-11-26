package entities

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestItUnmarshalsEntitiesJson(t *testing.T) {

	data, err := ioutil.ReadFile("testdata/entities.json")

	if err != nil {
		panic(err)
	}

	var entities Entities

	json.Unmarshal(data, &entities)

	if len(entities.HashTags) != 2 {
		t.Error("HashTags expected 2 got", len(entities.HashTags))
	}

	if len(entities.Urls) != 2 {
		t.Error("Urls expected 2 got", len(entities.Urls))
	}

	if len(entities.UserMentions) != 2 {
		t.Error("UserMentions expected 2 got", len(entities.UserMentions))
	}

	if len(entities.Symbols) != 2 {
		t.Error("Symbols expected 2 got", len(entities.Symbols))
	}
}
