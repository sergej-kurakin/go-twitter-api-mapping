package entities

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

var initialURL = URL{
	URL:         "http://www.example.com/",
	DisplayURL:  "http://www.example.com/",
	ExpandedURL: "http://www.example.com/",
}

func TestItUnmarshalsUrlJson(t *testing.T) {

	data, err := ioutil.ReadFile("testdata/url.json")

	if err != nil {
		panic(err)
	}

	var url URL

	json.Unmarshal(data, &url)

	assertURL(t, initialURL, url)
}

func TestItMarshalsUrlJson(t *testing.T) {

	body, err := json.Marshal(initialURL)

	tmpfile, err := ioutil.TempFile("testdata/", "url_*.json")
	tmpName := tmpfile.Name()
	defer os.Remove(tmpName)

	tmpfile.Write(body)

	tmpfile.Close()

	data, err := ioutil.ReadFile(tmpName)

	if err != nil {
		panic(err)
	}

	var url URL

	json.Unmarshal(data, &url)

	assertURL(t, initialURL, url)
}

func assertURL(t *testing.T, initialURL URL, url URL) {
	if url.URL != initialURL.URL {
		t.Error("URL expected", initialURL.URL, "got", url.URL)
	}

	if url.URL != initialURL.DisplayURL {
		t.Error("DisplayURL expected", initialURL.DisplayURL, "got", url.DisplayURL)
	}

	if url.URL != initialURL.ExpandedURL {
		t.Error("ExpandedURL expected", initialURL.ExpandedURL, "got", url.ExpandedURL)
	}
}
