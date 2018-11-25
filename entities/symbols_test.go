package entities

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

var initialSymbols = Symbols{
	Text: "Lorum ipsum",
}

func TestItUnmarshalsSymbolsJson(t *testing.T) {

	data, err := ioutil.ReadFile("testdata/symbols.json")

	if err != nil {
		panic(err)
	}

	var symbols Symbols

	json.Unmarshal(data, &symbols)

	assertSymbols(t, initialSymbols, symbols)
}

func TestItMarshalsSymbolsJson(t *testing.T) {

	body, err := json.Marshal(initialSymbols)

	tmpfile, err := ioutil.TempFile("testdata/", "symbols_*.json")
	tmpName := tmpfile.Name()
	defer os.Remove(tmpName)

	tmpfile.Write(body)

	tmpfile.Close()

	data, err := ioutil.ReadFile(tmpName)

	if err != nil {
		panic(err)
	}

	var symbols Symbols

	json.Unmarshal(data, &symbols)

	assertSymbols(t, initialSymbols, symbols)
}

func assertSymbols(t *testing.T, initialSymbols Symbols, symbols Symbols) {
	if symbols.Text != initialSymbols.Text {
		t.Error("Text expected", initialSymbols.Text, "got", symbols.Text)
	}
}
