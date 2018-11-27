package entities

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestItUnmarshalsEntitiesJson(t *testing.T) {

	initialEntities := createInitialEntities()

	data, err := ioutil.ReadFile("testdata/entities.json")

	if err != nil {
		panic(err)
	}

	var entities Entities

	json.Unmarshal(data, &entities)

	assertEntities(t, initialEntities, entities)

	if len(entities.Symbols) != 2 {
		t.Error("Symbols expected 2 got", len(entities.Symbols))
	}
}

func TestItMarshalsEntitiesJson(t *testing.T) {

	initialEntities := createInitialEntities()

	body, err := json.Marshal(initialEntities)

	tmpfile, err := ioutil.TempFile("testdata/", "entities_*.json")
	tmpName := tmpfile.Name()
	defer os.Remove(tmpName)

	tmpfile.Write(body)

	tmpfile.Close()

	data, err := ioutil.ReadFile(tmpName)

	if err != nil {
		panic(err)
	}

	var entities Entities

	json.Unmarshal(data, &entities)

	assertEntities(t, initialEntities, entities)
}

func createInitialEntities() Entities {
	initialEntities := Entities{
		HashTags: []HashTag{
			{
				Text: "Hashtag 1 Lorum ipsum",
			},
			{
				Text: "Hashtag 2 Lorum ipsum",
			},
		},
		Symbols: []Symbol{
			{
				Text: "Symbol 1 Lorum ipsum",
			},
			{
				Text: "Symbol 2 Lorum ipsum",
			},
		},
		UserMentions: []UserMention{
			{
				ScreenName: "User Mention 1 Screen Name",
				Name:       "User Mention 1 Just Name",
				IDStr:      "14568",
				ID:         14568,
			},
			{
				ScreenName: "User Mention 2 Screen Name",
				Name:       "User Mention 2 Just Name",
				IDStr:      "14569",
				ID:         14569,
			},
		},
		Urls: []URL{
			{
				URL:         "http://www.example.com/1",
				ExpandedURL: "http://www.example.com/expanded_url/1",
				DisplayURL:  "http://www.example.com/display_url/1",
			},
			{
				URL:         "http://www.example.com/2",
				ExpandedURL: "http://www.example.com/expanded_url/2",
				DisplayURL:  "http://www.example.com/display_url/2",
			},
		},
	}

	return initialEntities
}

func assertEntities(t *testing.T, initialEntities Entities, entities Entities) {
	// HashTags
	if len(entities.HashTags) != 2 {
		t.Error("HashTags expected 2 got", len(entities.HashTags))
	}
	assertHashTag(t, initialEntities.HashTags[0], entities.HashTags[0])
	assertHashTag(t, initialEntities.HashTags[1], entities.HashTags[1])

	// Urls
	if len(entities.Urls) != 2 {
		t.Error("Urls expected 2 got", len(entities.Urls))
	}
	assertURL(t, initialEntities.Urls[0], entities.Urls[0])
	assertURL(t, initialEntities.Urls[1], entities.Urls[1])

	// UserMentions
	if len(entities.UserMentions) != 2 {
		t.Error("UserMentions expected 2 got", len(entities.UserMentions))
	}
	assertUserMention(t, initialEntities.UserMentions[0], entities.UserMentions[0])
	assertUserMention(t, initialEntities.UserMentions[1], entities.UserMentions[1])

	// Symbols
	if len(entities.Symbols) != 2 {
		t.Error("Symbols expected 2 got", len(entities.Symbols))
	}
	assertSymbol(t, initialEntities.Symbols[0], entities.Symbols[0])
	assertSymbol(t, initialEntities.Symbols[1], entities.Symbols[1])
}
