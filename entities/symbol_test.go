package entities

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

var initialSymbol = Symbol{
	Text: "Lorum ipsum",
}

func TestItUnmarshalsSymbolJson(t *testing.T) {

	data, err := ioutil.ReadFile("testdata/symbol.json")

	if err != nil {
		panic(err)
	}

	var symbol Symbol

	json.Unmarshal(data, &symbol)

	assertSymbol(t, initialSymbol, symbol)
}

func TestItMarshalsSymbolJson(t *testing.T) {

	body, err := json.Marshal(initialSymbol)

	tmpfile, err := ioutil.TempFile("testdata/", "symbol_*.json")
	tmpName := tmpfile.Name()
	defer os.Remove(tmpName)

	tmpfile.Write(body)

	tmpfile.Close()

	data, err := ioutil.ReadFile(tmpName)

	if err != nil {
		panic(err)
	}

	var symbol Symbol

	json.Unmarshal(data, &symbol)

	assertSymbol(t, initialSymbol, symbol)
}

func assertSymbol(t *testing.T, initialSymbol Symbol, symbol Symbol) {
	if symbol.Text != initialSymbol.Text {
		t.Error("Text expected", initialSymbol.Text, "got", symbol.Text)
	}
}
