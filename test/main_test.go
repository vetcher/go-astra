package test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/vetcher/go-astra"
)

const (
	source    = "source.go"
	result    = "result.json"
	assetsDir = "assets"
)

type AstraTest struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func TestParsingFile(t *testing.T) {
	d, err := ioutil.ReadFile("./list_of_tests.json")
	if err != nil {
		t.Fatal(err)
	}
	var tests []AstraTest
	err = json.Unmarshal(d, &tests)
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.Name, func(t *testing.T) {
			expected, err := ioutil.ReadFile(filepath.Join(assetsDir, tt.Path, result))
			if err != nil {
				t.Fatal(err)
			}
			file, err := astra.ParseFile(filepath.Join(assetsDir, tt.Path, source))
			if err != nil {
				t.Fatal(err)
			}
			actual, err := json.Marshal(file)
			if err != nil {
				t.Fatal(err)
			}
			if !testEq(expected, actual) {
				t.Fatalf("expected != actual:\n%s\n\n%s", string(expected), string(actual))
			}
		})
	}
}

func testEq(a, b []byte) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
