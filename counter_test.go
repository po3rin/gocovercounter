package gocovercounter_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/po3rin/gocovercounter"
)

func TestAnnotate(t *testing.T) {
	w := &bytes.Buffer{}
	gocovercounter.Annotate(w, "./testdata/testfile.go")

	f, err := os.Open("./testdata/wanttestfile.txt")
	if err != nil {
		t.Fatalf("got unexpected error: %v\n", err)
	}
	defer f.Close()
	want, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("got unexpected error: %v\n", err)
	}

	if string(want) != w.String() {
		t.Fatalf("unexpected result:\nwant:\n%v\ngot:\n%v\n", string(want), w.String())
	}
}
