package main_test

import (
	grep "github.com/Adedunmol/go-grep"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestMatch(t *testing.T) {
	fileSystem := fstest.MapFS{
		"test.txt": {Data: []byte("some random text\nsome other text")},
	}
	files := []string{"test.txt"}
	candidate := "random"

	got := grep.MatchFiles(fileSystem, files, candidate)
	wanted := []grep.Match{}
	if !reflect.DeepEqual(got, wanted) {
		t.Errorf("wanted %v, got %v", wanted, got)
	}
}
